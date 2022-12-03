package main

import (
	"fmt"
	"github.com/nielsvangijzen/aoc2022/util"
	"strings"
)

const (
	Rock     = "A"
	Paper    = "B"
	Scissors = "C"

	RockResponse     = "X"
	PaperResponse    = "Y"
	ScissorsResponse = "Z"

	ShouldLose = "X"
	ShouldDraw = "Y"
	ShouldWin  = "Z"
)

// Here we map the opponent's pick to your response
// that results in a score
var scoreMap = map[string]map[string]int{
	Rock: {
		RockResponse:     1 + 3, // Draw
		PaperResponse:    2 + 6, // Win
		ScissorsResponse: 3 + 0, // Loss
	},
	Paper: {
		RockResponse:     1 + 0, // Loss
		PaperResponse:    2 + 3, // Draw
		ScissorsResponse: 3 + 6, // Win
	},
	Scissors: {
		RockResponse:     1 + 6, // Win
		PaperResponse:    2 + 0, // Loss
		ScissorsResponse: 3 + 3, // Draw
	},
}

// Here we map the opponent's pick to the strategy
// we map those two to a response
var outcomeMap = map[string]map[string]string{
	Rock: {
		ShouldLose: ScissorsResponse,
		ShouldDraw: RockResponse,
		ShouldWin:  PaperResponse,
	},
	Paper: {
		ShouldLose: RockResponse,
		ShouldDraw: PaperResponse,
		ShouldWin:  ScissorsResponse,
	},
	Scissors: {
		ShouldLose: PaperResponse,
		ShouldDraw: ScissorsResponse,
		ShouldWin:  RockResponse,
	},
}

func partOne() {
	input := util.MustInputString("day-02/input.txt")

	games := strings.Split(input, "\n")
	score := 0

	for _, game := range games {
		picks := strings.Split(game, " ")
		score += scoreMap[picks[0]][picks[1]]
	}

	fmt.Printf("D02P01: %d\n", score)
}

// partTwo is almost the same as partOne, but we
// look up our pick in the outcomeMap prior to
// looking up the score.
func partTwo() {
	input := util.MustInputString("day-02/input.txt")

	games := strings.Split(input, "\n")
	score := 0

	for _, game := range games {
		picks := strings.Split(game, " ")
		shouldPick := outcomeMap[picks[0]][picks[1]]
		score += scoreMap[picks[0]][shouldPick]
	}

	fmt.Printf("D02P01: %d\n", score)
}

func main() {
	partOne()
	partTwo()
}
