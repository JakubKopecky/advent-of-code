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

func main() {
	file, _ := os.Open("adventofcode.com_2022_day_3_input.txt")
	score := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		x := newRucksack(scanner.Text())
		score += getPoints(x.uniqueRune)
	}
	fmt.Println("score is", score)
}
