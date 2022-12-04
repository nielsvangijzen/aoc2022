package main

import (
	"fmt"
	"github.com/nielsvangijzen/aoc2022/util"
	"log"
	"strconv"
	"strings"
)

type Section struct {
	Start int
	End   int
}

// NewSection The description mus be in the form of: 43-55
func NewSection(description string) (Section, error) {
	startAndEnd := strings.Split(description, "-")
	start, err := strconv.Atoi(startAndEnd[0])
	if err != nil {
		return Section{}, err
	}

	end, err := strconv.Atoi(startAndEnd[1])
	if err != nil {
		return Section{}, err
	}

	return Section{Start: start, End: end}, nil
}

func partOne() {
	pairs := util.MustInputLines("day-04/input.txt")

	amountContaining := determineOverlap(pairs, sectionContainsOther)

	fmt.Printf("D04P01: %d\n", amountContaining)
}

func partTwo() {
	pairs := util.MustInputLines("day-04/input.txt")

	amountOverlapping := determineOverlap(pairs, sectionsOverlap)

	fmt.Printf("D04P02: %d\n", amountOverlapping)
}

// determineOverlap for both parts of today's challenge we need to do a lot of same things.
// The only thing which really changes is the way we determine if two ranges overlap.
// This function takes a function that takes 2 sections and spits out a bool.
// We can make two different overlap functions and have this function do
// the same work.
func determineOverlap(pairs []string, overlapFunc func(Section, Section) bool) (amountContaining int) {
	for _, pair := range pairs {
		sections := strings.Split(pair, ",")
		sectionA, err := NewSection(sections[0])
		if err != nil {
			log.Fatalln(err)
		}

		sectionB, err := NewSection(sections[1])
		if err != nil {
			log.Fatalln(err)
		}

		if overlapFunc(sectionA, sectionB) {
			amountContaining += 1
		}
	}

	return
}

func sectionContainsOther(sectionA, sectionB Section) bool {
	if sectionA.Start <= sectionB.Start && sectionA.End >= sectionB.End {
		return true
	}

	if sectionB.Start <= sectionA.Start && sectionB.End >= sectionA.End {
		return true
	}

	return false
}

// If the maximum start value is smaller than the minimal end value.
// We know that the two sections overlap somehow.
func sectionsOverlap(sectionA, sectionB Section) bool {
	maxStart := util.MaxValue([]int{sectionA.Start, sectionB.Start})
	minEnd := util.MinValue([]int{sectionA.End, sectionB.End})

	return maxStart <= minEnd
}

func main() {
	partOne()
	partTwo()
}
