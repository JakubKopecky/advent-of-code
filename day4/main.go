package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type workRangePair struct {
	pair [2]workRange
}

type workRange struct {
	start int
	end   int
}

func parseWorkRange(line string) workRangePair {
	split := strings.Split(line, ",")
	return workRangePair{
		pair: [2]workRange{
			newWorkRange(split[0]),
			newWorkRange(split[1]),
		},
	}
}

func newWorkRange(r string) workRange {
	leftSide := r[:strings.IndexRune(r, '-')]
	rightSide := r[strings.IndexRune(r, '-')+1:]

	leftSideInt, _ := strconv.Atoi(leftSide)
	rightSideInt, _ := strconv.Atoi(rightSide)

	return workRange{
		start: leftSideInt,
		end:   rightSideInt,
	}
}

func isFullContain(work workRangePair) bool {
	rangeA := work.pair[0]
	rangeB := work.pair[1]
	isFullContain := true

	for i := rangeA.start; i <= rangeA.end; i++ {
		if !isNumberInsideRange(i, rangeB) {
			isFullContain = false
			break
		}
	}
	if isFullContain {
		return isFullContain
	}

	isFullContain = true
	for i := rangeB.start; i <= rangeB.end; i++ {
		if !isNumberInsideRange(i, rangeA) {
			isFullContain = false
			break
		}
	}

	return isFullContain
}

func isPartialContain(work workRangePair) bool {
	rangeA := work.pair[0]
	rangeB := work.pair[1]

	for i := rangeA.start; i <= rangeA.end; i++ {
		if isNumberInsideRange(i, rangeB) {
			return true
		}
	}

	for i := rangeB.start; i <= rangeB.end; i++ {
		if isNumberInsideRange(i, rangeA) {
			return true
		}
	}

	return false
}

func isNumberInsideRange(i int, r workRange) bool {
	return i >= r.start && i <= r.end
}

func main() {
	file, _ := os.Open("adventofcode.com_2022_day_4_input.txt")
	defer file.Close()

	fullContainCount := 0
	partialContainCount := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if isFullContain(parseWorkRange(scanner.Text())) {
			fullContainCount++
		}

		if isPartialContain(parseWorkRange(scanner.Text())) {
			partialContainCount++
		}
	}

	fmt.Println("full containt count is", fullContainCount)
	fmt.Println("partial containt count is", partialContainCount)
}
