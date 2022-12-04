package main

import (
	"fmt"
	"github.com/nielsvangijzen/aoc2022/util"
	"log"
	"strings"
)

func partOne() {
	lines := util.MustInputLines("day-03/input.txt")

	prioritySum := 0

	for _, line := range lines {
		// We'd normally check if the line's length is an even number.
		// But after running once I know that all the lines are
		// the correct length so sue me.
		middle := len(line) / 2
		first := line[:middle]
		second := line[middle:]

		// We first determine the same character, don't ask me
		// why I dó error handling now...
		sameChar, err := findSameCharacter(first, second)
		if err != nil {
			log.Fatalln(err)
		}

		prioritySum += determinePriority(sameChar)
	}

	fmt.Printf("D03P01: %d\n", prioritySum)
}

func partTwo() {
	lines := util.MustInputLines("day-03/input.txt")

	// We first need to separate all the lines into groups of three.
	// We take the amount of lines and divide that by three.
	groupAmount := len(lines) / 3
	groups := make([][]string, groupAmount)

	for i := 0; i < groupAmount; i++ {
		groups[i] = []string{
			lines[0+(i*3)],
			lines[1+(i*3)],
			lines[2+(i*3)],
		}
	}

	priority := 0

	for _, group := range groups {
		sameInFirstTwo := findSameCharacters(group[0], group[1])
		sameInAll, err := findSameCharacter(sameInFirstTwo, group[2])
		if err != nil {
			log.Fatalln(err)
		}

		priority += determinePriority(sameInAll)
	}

	fmt.Printf("D03P02: %d\n", priority)
}

func findSameCharacter(first, second string) (rune, error) {
	for _, char := range first {
		// strings.Index returns -1 if the substring is not present.
		// We therefore check if the outcome is higher or equal to
		// zero.
		if strings.Index(second, string(char)) >= 0 {
			return char, nil
		}
	}

	// We can't return an empty char, so we return a '_'. It shouldn't
	// be used anyway because we generate an error.
	return '_', fmt.Errorf("strings don't share common item")
}

func findSameCharacters(first, second string) (sameChars string) {
	for _, char := range first {
		if strings.Index(second, string(char)) >= 0 {
			sameChars += string(char)
		}
	}

	return
}

func determinePriority(char rune) int {
	// Casting a rune to an int gives its ASCII representation.
	if 'A' >= char || char <= 'Z' {
		// We first check if the character is within 'A' and 'Z'. If so
		// we cast it to its ASCII value and subtract 38. Since A-Z starts at
		// 65 this leaves us with 27 for 'A'. We then count up sequentially.
		return int(char) - 38
	} else {
		// Same thing as above but we subtract 96 because 'a' starts at 97 and
		// 'a' is supposed to receive priority 1
		return int(char) - 96
	}
}

func main() {
	partOne()
	partTwo()
}
