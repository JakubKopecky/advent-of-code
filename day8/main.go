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

func allocTreeScore(height int, width int) *[][]int {
	toReturn := make([][]int, height)
	for i := 0; i < height; i++ {
		toReturn[i] = make([]int, width)
		for j := 0; j < width; j++ {
			toReturn[i][j] = 1
		}
	}
	return &toReturn
}

func goUp(trees [][]byte, startPoint cord, treeScore *[][]int) bool {
	startHeight := trees[startPoint.y][startPoint.x]
	score := 0

	for i := startPoint.y - 1; i >= 0; i-- {
		score++
		if trees[i][startPoint.x] >= startHeight {
			(*treeScore)[startPoint.y][startPoint.x] *= score
			return false
		}
	}
	(*treeScore)[startPoint.y][startPoint.x] *= score
	return true
}

func goDown(trees [][]byte, startPoint cord, treeScore *[][]int) bool {
	startHeight := trees[startPoint.y][startPoint.x]
	score := 0

	for i := startPoint.y + 1; i < len(trees); i++ {
		score++
		if trees[i][startPoint.x] >= startHeight {
			(*treeScore)[startPoint.y][startPoint.x] *= score
			return false
		}
	}
	(*treeScore)[startPoint.y][startPoint.x] *= score
	return true
}

func goLeft(trees [][]byte, startPoint cord, treeScore *[][]int) bool {
	startHeight := trees[startPoint.y][startPoint.x]
	score := 0

	for i := startPoint.x - 1; i >= 0; i-- {
		score++
		if trees[startPoint.y][i] >= startHeight {
			(*treeScore)[startPoint.y][startPoint.x] *= score
			return false
		}
	}
	(*treeScore)[startPoint.y][startPoint.x] *= score
	return true
}

func goRight(trees [][]byte, startPoint cord, treeScore *[][]int) bool {
	startHeight := trees[startPoint.y][startPoint.x]
	score := 0

	for i := startPoint.x + 1; i < len(trees[0]); i++ {
		score++
		if trees[startPoint.y][i] >= startHeight {
			(*treeScore)[startPoint.y][startPoint.x] *= score
			return false
		}
	}
	(*treeScore)[startPoint.y][startPoint.x] *= score
	return true
}

func loopTrees(trees [][]byte, treeScore *[][]int) int {
	count := 0
	for i := 0; i < len(trees); i++ {
		for j := 0; j < len(trees[0]); j++ {
			startPoint := cord{
				y: i,
				x: j,
			}
			up := goUp(trees, startPoint, treeScore)
			down := goDown(trees, startPoint, treeScore)
			left := goLeft(trees, startPoint, treeScore)
			right := goRight(trees, startPoint, treeScore)
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

func findMax(treeScore *[][]int) int {
	max := (*treeScore)[0][0]
	for _, i := range *treeScore {
		for _, x := range i {
			if x > max {
				max = x
			}
		}
	}
	return max
}

func main() {
	file, _ := os.Open("adventofcode.com_2022_day_8_input.txt")
	// file, _ := os.Open("example.txt")
	defer file.Close()
	x := newTrees(file)
	treeScore := allocTreeScore(len(x), len(x[0]))
	fmt.Println(loopTrees(x, treeScore))
	fmt.Println(findMax(treeScore))
}
