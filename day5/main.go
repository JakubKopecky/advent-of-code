package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/golang-collections/collections/stack"
)

func isStackRow(line string) bool {
	return strings.Contains(line, "[")
}

func isNumberRow(line string) bool {
	return strings.Contains(line, "1") && !strings.Contains(line, "move")
}

func isInstructionRow(line string) bool {
	return strings.Contains(line, "move")
}

func getNumberOfStacks(lines []string) int {
	count := 0
	for _, line := range lines {
		if isNumberRow(line) {
			for _, char := range line {
				if char != ' ' {
					count++
				}
			}
		}
	}
	return count
}

func getIndexOfBottomOfStacks(lines []string) int {
	for index, line := range lines {
		if isNumberRow(line) {
			return index - 1
		}
	}
	return 0
}

func getIndexOfStack(stackNumber int) int {
	return 1 + ((stackNumber - 1) * 4)
}

func createStacks(lines []string) []*stack.Stack {
	start := getIndexOfBottomOfStacks(lines)
	numOfStacks := getNumberOfStacks(lines)
	stacks := []*stack.Stack{}
	for x := 0; x < numOfStacks; x++ {
		stacks = append(stacks, stack.New())
	}

	for i := start; i >= 0; i-- {
		for j := 0; j < numOfStacks; j++ {
			if lines[i][getIndexOfStack(j+1)] != ' ' {
				stacks[j].Push(lines[i][getIndexOfStack(j+1)])
			}
		}
	}
	return stacks
}

func readFile(file *os.File) []string {
	toReturn := []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		toReturn = append(toReturn, scanner.Text())
	}
	return toReturn
}

func showStacks(stacks []*stack.Stack) {
	for _, stack := range stacks {
		for x := stack.Pop(); x != nil; x = stack.Pop() {
			fmt.Printf("%c", x)
		}
		fmt.Println()
	}
}

func doInstruction(line string, stacks []*stack.Stack) {
	split := strings.Split(line, " ")
	from, _ := strconv.Atoi(split[3])
	to, _ := strconv.Atoi(split[5])

	x := stacks[from-1].Pop()
	stacks[to-1].Push(x)
}

func loopInstructions(lines []string, stacks []*stack.Stack) {
	for _, line := range lines {
		if isInstructionRow(line) {
			numberStr := strings.Split(line, " ")[1]
			number, _ := strconv.Atoi(numberStr)
			for number > 0 {
				doInstruction(line, stacks)
				number--
			}
		}
	}
}

func peekAllStacks(stacks []*stack.Stack) {
	for _, stack := range stacks {
		fmt.Printf("%c", stack.Peek())
	}
	fmt.Println()
}

func main() {
	file, _ := os.Open("adventofcode.com_2022_day_5_input.txt")
	// file, _ := os.Open("example.txt")
	defer file.Close()

	lines := readFile(file)
	stacks := createStacks(lines)
	loopInstructions(lines, stacks)
	peekAllStacks(stacks)
}
