package main

import (
	"fmt"
	"github.com/nielsvangijzen/aoc2022/util"
	"log"
	"strconv"
	"strings"
)

type CPU struct {
	// State
	registerX int
	cycle     int

	// Input
	input []string
	index int

	cycleHook func(*CPU)
}

func NewCPU(input []string, cycleHook func(*CPU)) *CPU {
	return &CPU{
		registerX: 1,
		cycle:     1,

		input: input,
		index: 0,

		cycleHook: cycleHook,
	}
}

func (cpu *CPU) run() error {
	for cpu.index < len(cpu.input) {
		// We split the current line on a single space. We now have the left and right
		// part of the instruction.
		parts := strings.Split(cpu.input[cpu.index], " ")

		// parts[0] will be the instruction
		switch parts[0] {
		case "noop":
			noop(cpu)
		case "addx":
			amount, err := strconv.Atoi(parts[1])
			if err != nil {
				return err
			}
			addx(cpu, amount)
		}

		cpu.index += 1
	}

	return nil
}

func cycle(cpu *CPU) {
	// We run the cycleHook inserted into the CPU and the increase the cycle counter
	// by one.
	cpu.cycleHook(cpu)
	cpu.cycle += 1
}

// noop just executes one cycle.
func noop(cpu *CPU) {
	cycle(cpu)
}

// addx executes two cycles prior to increasing the registerX by the given
// amount.
func addx(cpu *CPU, amount int) {
	// Addx
	cycle(cpu)
	cycle(cpu)
	cpu.registerX += amount
}

type CounterHook struct {
	sum int
}

func (ch *CounterHook) cycleHook(cpu *CPU) {
	if cpu.cycle == 20 || (cpu.cycle+20)%40 == 0 {
		ch.sum += cpu.cycle * cpu.registerX
	}
}

func partOne() {
	input := util.MustInputLines("day-10/input.txt")

	// Because we have no way of reading a return value by using the hook system. We
	// create a struct that contains a sum. We then instantiate that struct by
	// pointer so the hook can increase its sum.
	ch := &CounterHook{}
	cpu := NewCPU(input, ch.cycleHook)

	err := cpu.run()
	if err != nil {
		log.Fatal(err)
	}

	// After running, the sum in the counter hook should contain the sum of signal
	// strengths.
	fmt.Printf("D10P01: %d\n", ch.sum)
}

type DrawHook struct {
	output string
}

func (dh *DrawHook) drawHook(cpu *CPU) {
	// The position of the sprite will reside in registerX
	spritePosition := cpu.registerX

	// We determine the position by dividing the cycle-1 count by 40 and taking the
	// remainder.
	cyclePosition := (cpu.cycle - 1) % 40
	drawn := false

	// We create a for loop in order to check the current position -1, 0, 1
	for i := -1; i < 2; i++ {
		if cyclePosition+i == spritePosition {
			dh.output += "#"
			drawn = true
		}
	}

	// If nothing was draw, we output a dot.
	if !drawn {
		dh.output += "."
	}

	// When we've drawn 40 characters, we add a line break.
	if cpu.cycle%40 == 0 {
		dh.output += "\n"
	}
}

func partTwo() {
	input := util.MustInputLines("day-10/input.txt")

	// We use the same struct and pointer trick as the in part one. But this one
	// holds a string for output instead.
	drawHook := &DrawHook{}
	cpu := NewCPU(input, drawHook.drawHook)

	err := cpu.run()
	if err != nil {
		log.Fatal(err)
	}

	// We print the 8 capital letters.
	fmt.Printf("D10P02:\n%s", drawHook.output)
}

func main() {
	partOne()
	partTwo()
}
