package main

import (
	"fmt"
	"github.com/nielsvangijzen/aoc2022/util"
	"log"
	"sort"
	"strconv"
	"strings"
)

func partOne() {
	input := util.MustInputString("day-01/input.txt")

	caloriesPerElf := calculateTotals(input)

	fmt.Printf("D01P01: %d\n", util.MaxValue(caloriesPerElf))
}

func partTwo() {
	// We take the same approach as last time
	input := util.MustInputString("day-01/input.txt")

	caloriesPerElf := calculateTotals(input)

	// Now we sort the slice leaving big numbers at the end
	sort.Ints(caloriesPerElf)

	lastThree := caloriesPerElf[len(caloriesPerElf)-3:]

	fmt.Printf("D01P02: %d", util.Sum(lastThree))
}

func calculateTotals(input string) (caloriesPerElf []int) {
	// Split on double linebreaks to separate input into elves
	elves := strings.Split(input, "\n\n")

	for _, elf := range elves {
		foodItems := strings.Split(elf, "\n")
		totalCalories := 0

		for _, item := range foodItems {
			if item == "" {
				continue
			}
			calories, err := strconv.Atoi(item)
			if err != nil {
				log.Fatalln(err)
			}

			totalCalories += calories
		}

		caloriesPerElf = append(caloriesPerElf, totalCalories)
	}

	return
}

func main() {
	partOne()
	partTwo()
}
