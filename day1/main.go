package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
)

func main() {
	var results []int
	x := 0

	file, _ := os.Open("adventofcode.com_2022_day_1_input.txt")

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if scanner.Text() == "" {
			results = append(results, x)
			x = 0
		}
		add, _ := strconv.Atoi(scanner.Text())
		x += add
	}

	slices.Sort(results)
	slices.Reverse(results)

	fmt.Println("top elf has", results[0], "calories")

	top3 := results[0] + results[1] + results[2]
	fmt.Println("top 3 are carring", top3, "calories")
}
