package main

import (
	"fmt"
	"github.com/nielsvangijzen/aoc2022/util"
	"log"
	"sort"
	"strconv"
)

type treeGrid [][]int

func parseGrid(input []string) (treeGrid, error) {
	height := len(input)

	grid := make([][]int, height)

	for index, row := range input {
		grid[index] = make([]int, len(row))

		for treeIndex, tree := range row {
			treeHeight, err := strconv.Atoi(string(tree))
			if err != nil {
				return nil, err
			}

			grid[index][treeIndex] = treeHeight
		}
	}

	return grid, nil
}

func (grid treeGrid) IsVisibleAndScore(x, y int) (bool, int) {
	// We check if either of the coordinates is zero or the maximum index on the
	// grid. If so, the tree is on the outside and is visible by default.
	if x == 0 || y == 0 || x == len(grid[0])-1 || y == len(grid)-1 {
		return true, 0
	}

	yVisible, yScore := grid.yVisibleAndScore(x, y)
	xVisible, xScore := grid.xVisibleAndScore(x, y)
	return yVisible || xVisible, xScore * yScore
}

func (grid treeGrid) yVisibleAndScore(x, y int) (bool, int) {
	// We check the height, it doesn't matter if we check the tree with itself
	// because it's not shorter than itself.
	treeHeight := grid[y][x]

	yVisibleLow := true
	yVisibleHigh := true
	yScoreLow := 0
	yScoreHigh := 0

	for currentY := y - 1; currentY >= 0; currentY-- {
		currentTree := grid[currentY][x]
		yScoreLow += 1
		if currentTree >= treeHeight {
			yVisibleLow = false
			break
		}
	}

	for currentY := y + 1; currentY < len(grid); currentY++ {
		currentTree := grid[currentY][x]
		yScoreHigh += 1
		if currentTree >= treeHeight {
			yVisibleHigh = false
			break
		}
	}

	return yVisibleHigh || yVisibleLow, yScoreLow * yScoreHigh
}

func (grid treeGrid) xVisibleAndScore(x, y int) (bool, int) {
	// We check the height, it doesn't matter if we check the tree with itself
	// because it's not shorter than itself.
	treeHeight := grid[y][x]

	xVisibleLow := true
	xVisibleHigh := true
	xScoreLow := 0
	xScoreHigh := 0

	for currentX := x - 1; currentX >= 0; currentX-- {
		currentTree := grid[y][currentX]
		xScoreLow += 1
		if currentTree >= treeHeight {
			xVisibleLow = false
			break
		}
	}

	for currentX := x + 1; currentX < len(grid[y]); currentX++ {
		currentTree := grid[y][currentX]
		xScoreHigh += 1
		if currentTree >= treeHeight {
			xVisibleHigh = false
			break
		}
	}

	return xVisibleLow || xVisibleHigh, xScoreLow * xScoreHigh
}

func partOne() {
	input := util.MustInputLines("day-08/input.txt")
	grid, err := parseGrid(input)
	if err != nil {
		log.Fatal(err)
	}

	visibleCount := 0

	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[0]); x++ {
			visible, _ := grid.IsVisibleAndScore(x, y)
			if visible {
				visibleCount += 1
			}
		}
	}

	fmt.Printf("D08P01: %d\n", visibleCount)
}

func partTwo() {
	input := util.MustInputLines("day-08/input.txt")
	grid, err := parseGrid(input)
	if err != nil {
		log.Fatal(err)
	}

	var scores []int

	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[0]); x++ {
			_, score := grid.IsVisibleAndScore(x, y)
			scores = append(scores, score)
		}
	}

	sort.Ints(scores)

	fmt.Printf("D08P01: %d\n", scores[len(scores)-1])
}

func main() {
	partOne()
	partTwo()
}
