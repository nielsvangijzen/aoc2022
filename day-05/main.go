package main

import (
	"fmt"
	"github.com/nielsvangijzen/aoc2022/util"
	"log"
	"strconv"
	"strings"
)

type Crate struct {
	letter string
}

type MoveInstruction struct {
	amount int
	from   int
	to     int
}

func partOne() {
	input := util.MustInputString("day-05/input.txt")
	parts := strings.Split(input, "\n\n")

	stacks, err := parseStacks(parts[0])
	if err != nil {
		log.Fatalln(err)
	}

	moves, err := parseMoves(parts[1])
	if err != nil {
		log.Fatalln(err)
	}

	// We loop over all the parsed moves from start to end
	for _, move := range moves {
		// Because we move one crate at a time, we use a for loop on the amount
		// of moves in this particular instruction.
		for i := 0; i < move.amount; i++ {
			// We pop the current crate from the stack and check if the stack
			// isn't empty
			value, ok := stacks[move.from-1].Pop()
			if !ok {
				log.Fatalln(fmt.Errorf("stacks %d is empty", i))
			}

			// We push the crate onto the new stack
			stacks[move.to-1].Push(value)
		}
	}

	fmt.Print("D05P01: ")
	for _, stack := range stacks {
		crate, ok := stack.Pop()
		if !ok {
			continue
		}

		fmt.Printf("%s", crate.letter)
	}

	fmt.Println()
}

func partTwo() {
	input := util.MustInputString("day-05/input.txt")
	parts := strings.Split(input, "\n\n")

	stacks, err := parseStacks(parts[0])
	if err != nil {
		log.Fatalln(err)
	}

	moves, err := parseMoves(parts[1])
	if err != nil {
		log.Fatalln(err)
	}

	// We loop over all the parsed moves from start to end
	for _, move := range moves {
		// We move a whole stack at one time. So we create an empty stack and
		// load up it up crate by crate.
		localStack := make(util.Stack[Crate], move.amount)
		for i := 0; i < move.amount; i++ {
			// We pop the current crate from the stack and check if the stack
			// isn't empty
			value, ok := stacks[move.from-1].Pop()
			if !ok {
				log.Fatalln(fmt.Errorf("stacks %d is empty", i))
			}

			// We push the crate onto the new stack
			localStack.Push(value)
		}

		for i := 0; i < move.amount; i++ {
			// Logically the stack can't be empty, so we don't care about
			// checking
			crate, _ := localStack.Pop()

			stacks[move.to-1].Push(crate)
		}
	}

	fmt.Print("D05P02: ")
	for _, stack := range stacks {
		crate, ok := stack.Pop()
		if !ok {
			continue
		}

		fmt.Printf("%s", crate.letter)
	}

	fmt.Println()
}

func parseStacks(input string) ([]util.Stack[Crate], error) {
	// We want to separate the input on lines to figure out the crates
	lines := strings.Split(input, "\n")

	// The height will be the length of the above slice minus two. One to account
	// for the index starting at 0. And one to account for the row of numbers at
	// the bottom.
	height := len(lines) - 2

	// Now we figure out the width by taking the last character of the last line.
	// This is the highest number of the number row on the bottom.
	lastLine := lines[height+1]
	lastChar := lastLine[len(lastLine)-1]
	width, err := strconv.Atoi(string(lastChar))
	if err != nil {
		return nil, err
	}

	// Now we create our stack using the known width of stacks
	stack := make([]util.Stack[Crate], width)

	// We start iterating at the bottom to fill up our stack structure.
	// Since we can only push and pop we need the first push to be the
	// bottom most crate.
	for i := height; i >= 0; i-- {
		row := lines[i]
		for j := 0; j < width; j++ {
			// Crates are surrounded by square brackets like so: [A]. If we only want
			// its letter we add one as offset and pick the current width times 4 to
			// get to the next like so: [Z] [M] [P]
			//							0123456789
			letterIndex := 1 + (j * 4)
			if letterIndex >= len(row) {
				break
			}

			// If we hit a space it means there's no crate here.
			letter := row[letterIndex]
			if letter == ' ' {
				continue
			}

			stack[j].Push(Crate{string(letter)})
		}
	}

	return stack, err
}

func parseMoves(input string) ([]MoveInstruction, error) {
	moves := strings.Split(input, "\n")
	instructions := make([]MoveInstruction, len(moves))

	for index, instruction := range moves {
		// We do my trademark replacing instead of using regex. Because
		// regex is evil and should never be used.
		ins := strings.Replace(instruction, "move ", "", -1)
		ins = strings.Replace(ins, "from ", "", -1)
		ins = strings.Replace(ins, "to ", "", -1)

		// Now we should have a nice string with all numbers surrounded by a
		// single space ' '
		parts := strings.Split(ins, " ")

		move, err := strconv.Atoi(parts[0])
		if err != nil {
			return nil, fmt.Errorf("unable parsing move %s: %s", instruction, err)
		}

		from, err := strconv.Atoi(parts[1])
		if err != nil {
			return nil, fmt.Errorf("unable parsing from %s: %s", instruction, err)
		}

		to, err := strconv.Atoi(parts[2])
		if err != nil {
			return nil, fmt.Errorf("unable parsing to %s: %s", instruction, err)
		}

		instructions[index] = MoveInstruction{
			amount: move,
			from:   from,
			to:     to,
		}
	}

	return instructions, nil
}

func main() {
	partOne()
	partTwo()
}
