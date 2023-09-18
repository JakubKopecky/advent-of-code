package main

import (
	"bufio"
	"fmt"
	"os"
)

type cord struct {
	y int
	x int
}

func newTrees(file *os.File) [][]byte {
	reader := bufio.NewReader(file)
	trees := [][]byte{}
	line := []byte{}
	for {
		b, err := reader.ReadByte()

		if b == '\n' {
			trees = append(trees, line)
			line = []byte{}
			continue
		}

		line = append(line, b)

		if err != nil {
			break
		}
	}
	return trees
}

func goUp(trees [][]byte, startPoint cord) bool {
	startHeight := trees[startPoint.y][startPoint.x]

	for i := startPoint.y - 1; i >= 0; i-- {
		if trees[i][startPoint.x] >= startHeight {
			return false
		}
	}
	return true
}

func goDown(trees [][]byte, startPoint cord) bool {
	startHeight := trees[startPoint.y][startPoint.x]

	for i := startPoint.y + 1; i < len(trees); i++ {
		if trees[i][startPoint.x] >= startHeight {
			return false
		}
	}
	return true
}

func goLeft(trees [][]byte, startPoint cord) bool {
	startHeight := trees[startPoint.y][startPoint.x]

	for i := startPoint.x - 1; i >= 0; i-- {
		if trees[startPoint.y][i] >= startHeight {
			return false
		}
	}
	return true
}

func goRight(trees [][]byte, startPoint cord) bool {
	startHeight := trees[startPoint.y][startPoint.x]

	for i := startPoint.x + 1; i < len(trees[0]); i++ {
		if trees[startPoint.y][i] >= startHeight {
			return false
		}
	}
	return true
}

func loopTrees(trees [][]byte) int {
	count := 0
	for i := 0; i < len(trees); i++ {
		for j := 0; j < len(trees[0]); j++ {
			startPoint := cord{
				y: i,
				x: j,
			}
			up := goUp(trees, startPoint)
			down := goDown(trees, startPoint)
			left := goLeft(trees, startPoint)
			right := goRight(trees, startPoint)
			if up || down || left || right {
				count++
			}
		}
	}
	return count
}

func showTrees(trees [][]byte) {
	for _, i := range trees {
		fmt.Println(i)
	}
}

func main() {
	file, _ := os.Open("adventofcode.com_2022_day_8_input.txt")
	// file, _ := os.Open("example.txt")
	defer file.Close()
	x := newTrees(file)
	fmt.Println(loopTrees(x))
}
