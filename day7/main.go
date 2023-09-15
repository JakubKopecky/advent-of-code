package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type item struct {
	parent  *item
	name    string
	content []*item
	size    int
}

func (i *item) addItem(newItem *item) {
	i.content = append(i.content, newItem)
}

func (i *item) moveToItem(itemName string) *item {
	for _, j := range i.content {
		if j.name == itemName {
			return j
		}
	}
	return nil
}

func newItem(parent *item, name string, size int) *item {
	return &item{
		parent: parent,
		name:   name,
		size:   size,
	}
}

func isCd(line string) bool {
	return strings.Contains(line, "$ cd")
}

func isLs(line string) bool {
	return strings.Contains(line, "$ ls")
}

func isLsLine(line string) bool {
	return !strings.Contains(line, "$")
}

func getCdArg(line string) string {
	substring := strings.Split(line, " ")
	return substring[2]
}

func getLsLineInfo(line string) (int, string) {
	sub := strings.Split(line, " ")
	size, err := strconv.Atoi(sub[0])
	if err != nil {
		return 0, sub[1]
	}
	return size, sub[1]
}

var sizes []int

func getDirSize(dir *item) int {
	total := 0
	for _, i := range dir.content {
		total += i.size
		if i.size == 0 {
			total += getDirSize(i)
		}
	}
	sizes = append(sizes, total)
	return total
}

func main() {
	// file, _ := os.Open("example.txt")
	file, _ := os.Open("adventofcode.com_2022_day_7_input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	root := newItem(nil, "/", 0)
	pointer := root

	for scanner.Scan() {
		line := scanner.Text()

		if isLs(line) {
			continue
		}

		if isCd(line) {
			cdArg := getCdArg(line)
			if cdArg == "/" {
				pointer = root
				continue
			}
			if cdArg == ".." {
				pointer = pointer.parent
				continue
			}
			pointer = pointer.moveToItem(cdArg)
		}

		if isLsLine(line) {
			size, name := getLsLineInfo(line)
			newItem := newItem(pointer, name, size)
			pointer.addItem(newItem)
		}
	}

	getDirSize(root)

	sum := 0
	for _, size := range sizes {
		if size < 100000 {
			sum += size
		}
	}
	fmt.Println(sum)

	diskSize := 70000000
	neededFreeSpace := 30000000
	usedSpace := sizes[len(sizes)-1]
	unusedSpace := diskSize - usedSpace

	validSizes := []int{}
	for _, size := range sizes {
		if unusedSpace+size > neededFreeSpace {
			validSizes = append(validSizes, size)
		}
	}
	slices.Sort(validSizes)
	fmt.Println(validSizes[0])
}
