package main

import (
	"bufio"
	"fmt"
	"os"
)

type rucksack struct {
	halfA      string
	halfB      string
	uniqueRune rune
}

func newRucksack(line string) rucksack {
	halfLen := len(line) / 2
	x := rucksack{
		halfA: line[:halfLen],
		halfB: line[halfLen:],
	}
	findUniqueRune(&x)
	return x
}

func findUniqueRune(rucksack *rucksack) {
	for _, c := range rucksack.halfA {
		for _, cb := range rucksack.halfB {
			if c == cb {
				rucksack.uniqueRune = c
				return
			}
		}
	}
}

func getPoints(x rune) int {
	if 'a' <= x && x <= 'z' {
		return int(x - 96)
	}
	if 'A' <= x && x <= 'Z' {
		return int(x - 64 + 26)
	}
	return 0
}

// PART 2

type elfGroup struct {
	items      [3]string
	uniqueRune rune
}

func newElfGroup(newItems [3]string) *elfGroup {
	return &elfGroup{
		items: newItems,
	}
}

func (e *elfGroup) removeDuplicates() *elfGroup {
	for index := range e.items {
		e.items[index] = removeDuplicateFromString(e.items[index])
	}
	return e
}

func removeDuplicateFromString(line string) string {
	chars := make(map[rune]bool)
	toReturn := []rune{}
	for _, char := range line {
		if _, value := chars[char]; !value {
			chars[char] = true
			toReturn = append(toReturn, char)
		}
	}
	return string(toReturn)
}

func (e *elfGroup) findUniqueRune() *elfGroup {
	table := make(map[rune]int)

	for _, line := range e.items {
		for _, char := range line {
			table[char] += 1
		}
	}

	for key, value := range table {
		if value == 3 {
			e.uniqueRune = key
			return e
		}
	}
	return e
}

func main() {
	file, _ := os.Open("adventofcode.com_2022_day_3_input.txt")
	defer file.Close()
	score := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		x := newRucksack(scanner.Text())
		score += getPoints(x.uniqueRune)
	}
	fmt.Println("score is", score)

	file2, _ := os.Open("adventofcode.com_2022_day_3_input.txt")
	defer file2.Close()
	score2 := 0

	scanner = bufio.NewScanner(file2)
	for scanner.Scan() {
		var lines [3]string
		lines[0] = scanner.Text()
		scanner.Scan()
		lines[1] = scanner.Text()
		scanner.Scan()
		lines[2] = scanner.Text()

		group := newElfGroup(lines).removeDuplicates().findUniqueRune()
		score2 += getPoints(group.uniqueRune)
	}
	fmt.Println("score for part2 is", score2)
}
