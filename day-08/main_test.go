package main

import (
	"github.com/nielsvangijzen/aoc2022/util"
	"log"
	"testing"
)

func TestIsVisible(t *testing.T) {
	input := util.MustInputLines("example.txt")
	grid, err := parseGrid(input)
	if err != nil {
		log.Fatal(err)
	}

	// 30373
	// 25512
	// 65332
	// 33549
	// 35390

	// shouldBeVisible is an array that reflects whether a tree is visible
	// on the treemap from the example.
	shouldBeVisible := [][]bool{
		{true, true, true, true, true},
		{true, true, true, false, true},
		{true, true, false, true, true},
		{true, false, true, false, true},
		{true, true, true, true, true},
	}

	for y, row := range shouldBeVisible {
		for x, vis := range row {
			visible, _ := grid.IsVisibleAndScore(x, y)

			if visible != vis {
				t.Fatalf("IsVisibleAndScore doesn't return %v at position x:%d,y:%d\n", vis, x, y)
			}
		}
	}
}

func TestScore(t *testing.T) {
	input := util.MustInputLines("example.txt")
	grid, err := parseGrid(input)
	if err != nil {
		log.Fatal(err)
	}

	_, score := grid.IsVisibleAndScore(2, 1)
	if score != 4 {
		t.Fatalf("tree 2,1 should have score of 4 but gets %d", score)
	}

	_, score = grid.IsVisibleAndScore(2, 3)
	if score != 8 {
		t.Fatalf("tree 2,3 should have score of 8 but gets %d", score)
	}
}
