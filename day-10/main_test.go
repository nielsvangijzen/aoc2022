package main

import (
	"github.com/nielsvangijzen/aoc2022/util"
	"testing"
)

func TestCycles(t *testing.T) {
	input := util.MustInputLines("example1.txt")
	cpu := NewCPU(input, func(cpu *CPU) {})

	err := cpu.run()
	if err != nil {
		t.Fatal(err)
	}

	if cpu.registerX != -1 {
		t.Fatalf("registerX should've been -1 but is %d", cpu.registerX)
	}
}

func TestHook(t *testing.T) {
	input := util.MustInputLines("example2.txt")
	hc := &CounterHook{}

	cpu := NewCPU(input, hc.cycleHook)

	err := cpu.run()
	if err != nil {
		t.Fatal(err)
	}

	if hc.sum != 13140 {
		t.Fatalf("sum of the CounterHook should be 13140 but came out as: %d", hc.sum)
	}
}

func TestDrawHook(t *testing.T) {
	input := util.MustInputLines("example2.txt")
	drawHook := &DrawHook{}

	cpu := NewCPU(input, drawHook.drawHook)

	err := cpu.run()
	if err != nil {
		t.Fatal(err)
	}

	if drawHook.output != "##..##..##..##..##..##..##..##..##..##..\n###...###...###...###...###...###...###.\n####....####....####....####....####....\n#####.....#####.....#####.....#####.....\n######......######......######......####\n#######.......#######.......#######.....\n" {
		t.Fatalf("output didn't return expected output")
	}

}
