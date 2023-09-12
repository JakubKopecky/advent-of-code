package main

import (
	"bufio"
	"fmt"
	"os"
)

// A - X - rock - 1 point
// B - Y - paper - 2 points
// C - Z - scissors - 3 points
// 6 points if win
// 0 if lose
// 3 if draw

type game struct {
	playerA rune
	playerB rune
}

func createGame(line string) game {
	return game{
		playerA: []rune(line)[0],
		playerB: []rune(line)[2],
	}
}

func outcomePoints(game game) int {
	if game.playerA == game.playerB-23 {
		return 3
	}
	if game.playerB == 'X' && game.playerA == 'C' {
		return 6
	}
	if game.playerB == 'Y' && game.playerA == 'A' {
		return 6
	}
	if game.playerB == 'Z' && game.playerA == 'B' {
		return 6
	}
	return 0
}

func selectionPointsPlayerB(game game) int {
	return int(game.playerB - 87)
}

func selectionPointsPlayerA(game game) int {
	return int(game.playerA - 64)
}

// part2
// X - need lose
// Y - need draw
// Z - need win

func doOutcome(game game) int {
	if game.playerB == 'Y' {
		return selectionPointsPlayerA(game) + 3
	}
	if game.playerB == 'Z' {
	}
	return 0
}

func main() {
	totalWinScore := 0
	totalWinScore2 := 0

	file, _ := os.Open("adventofcode.com_2022_day_2_input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// part1
		round := createGame(scanner.Text())
		totalWinScore += outcomePoints(round)
		totalWinScore += selectionPointsPlayerB(round)

		// part2
		round = createGame(scanner.Text())
		totalWinScore2 += doOutcome(round)
	}

	fmt.Println("total score is", totalWinScore)
	fmt.Println("total score for part2 is", totalWinScore2)
}
