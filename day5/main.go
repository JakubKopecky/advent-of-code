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

func splitInstruction(line string) (int, int, int) {
	split := strings.Split(line, " ")
	amount, _ := strconv.Atoi(split[1])
	from, _ := strconv.Atoi(split[3])
	to, _ := strconv.Atoi(split[5])
	return amount, from, to
}

func doInstruction(line string, stacks []*stack.Stack) {
	_, from, to := splitInstruction(line)

	x := stacks[from-1].Pop()
	stacks[to-1].Push(x)
}

func loopInstructions(lines []string, stacks []*stack.Stack) {
	for _, line := range lines {
		if isInstructionRow(line) {
			number, _, _ := splitInstruction(line)
			for number > 0 {
				doInstruction(line, stacks)
				number--
			}
		}
	}
}

func loopInstructionsPart2(lines []string, stacks []*stack.Stack) {
	for _, line := range lines {
		if isInstructionRow(line) {
			doInstructionPart2(line, stacks)
		}
	}
}

func doInstructionPart2(line string, stacks []*stack.Stack) {
	amount, from, to := splitInstruction(line)
	moved := []interface{}{}
	for i := 0; i < amount; i++ {
		moved = append(moved, stacks[from-1].Pop())
	}
	for i := amount - 1; i >= 0; i-- {
		stacks[to-1].Push(moved[i])
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
	defer file.Close()

	lines := readFile(file)

	stacks := createStacks(lines)
	loopInstructions(lines, stacks)
	peekAllStacks(stacks)

	stacks2 := createStacks(lines)
	loopInstructionsPart2(lines, stacks2)
	peekAllStacks(stacks2)
}
