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

	fmt.Println(results[len(results)-1])
}
