package main

import (
	"fmt"
	"github.com/nielsvangijzen/aoc2022/util"
)

func partOne() {
	input := util.MustInputString("day-06/input.txt")

	seqAt := firstUniqueSequence(input, 4)

	fmt.Printf("D06P01: %d\n", seqAt)
}

func partTwo() {
	input := util.MustInputString("day-06/input.txt")

	seqAt := firstUniqueSequence(input, 14)

	fmt.Printf("D06P01: %d\n", seqAt)
}

func firstUniqueSequence(seq string, length int) int {
	// We loop through the string slicing parts of $length
	for i := 0; i < len(seq)-length; i++ {
		substr := seq[i : i+length]

		// Unique returns true if there are no double elements in the slice
		// we cast it to a rune slice because the function needs a slice.
		if util.Unique([]rune(substr)) {
			return i + length
		}
	}

	// Return -1 if not found. In normal instances I'd return an error
	// but for now we leave it at this.
	return -1
}

func main() {
	partOne()
	partTwo()
}
