package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func isMessage(line string, i int) bool {
	substring := line[i-13 : i+1]
	return isStringUniqueChars(substring)
}

func findMessage(line string) int {
	for i := 13; i < len(line); i++ {
		if isMessage(line, i) {
			return i
		}
	}
	return 0
}

func isMarker(line string, i int) bool {
	substring := line[i-3 : i+1]
	return isStringUniqueChars(substring)
}

func findMarker(line string) int {
	for i := 3; i < len(line); i++ {
		if isMarker(line, i) {
			return i
		}
	}
	return 0
}

func isStringUniqueChars(line string) bool {
	for _, char := range line {
		if strings.Count(line, string(char)) > 1 {
			return false
		}
	}
	return true
}

func main() {
	file, _ := os.Open("adventofcode.com_2022_day_6_input.txt")
	scanner := bufio.NewScanner(file)

	scanner.Scan()
	line := scanner.Text()
	fmt.Println(findMarker(line) + 1)
	fmt.Println(findMessage(line) + 1)
}
