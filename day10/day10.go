package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Command int

const (
	NOOP Command = iota
	ADDX Command = iota
)

type CommandStatement struct {
	command Command
	value   int
	cycles  int
}

func ShouldDrawPixel(spritePosX, cycleCount int) bool {
	spritePos := []int{
		spritePosX - 1,
		spritePosX,
		spritePosX + 1,
	}
	for _, pos := range spritePos {
		if pos == cycleCount {
			return true
		}
	}
	return false
}

func ShouldCheckSignalStrength(times, cycles int) bool {
	nextCycleCheck := 20 + (times * 40)
	return nextCycleCheck == cycles
}

func main() {
	dat, err := os.ReadFile("input")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(dat), "\n")
	lines = lines[:len(lines)-1]

	statements := Queue[CommandStatement]{}
	for _, line := range lines {
		if line == "noop" {
			statements = statements.Enqueue(CommandStatement{NOOP, 0, 1})
		} else {
			addxStatement := strings.Split(line, " ")
			operand, _ := strconv.Atoi(addxStatement[1])
			statements = statements.Enqueue(CommandStatement{ADDX, operand, 2})
		}
	}

	x := 1
	cycleCount := 0
	crt := [6][40]string{}

	signalStrengthCheckCount := 0
	signalStrengthSum := 0
	for !statements.IsEmpty() {
		cycleCount += 1
		statements[0].cycles -= 1
		if ShouldCheckSignalStrength(signalStrengthCheckCount, cycleCount) {
			signalStrengthCheckCount += 1
			signalStrengthSum += cycleCount * x
		}

		crtRow := (cycleCount - 1) / 40
		crtColumn := (cycleCount - 1) % 40

		if ShouldDrawPixel(x, crtColumn) {
			crt[crtRow][crtColumn] = "#"
		} else {
			crt[crtRow][crtColumn] = "."
		}

		if statements[0].command == ADDX && statements[0].cycles == 0 {
			x += statements[0].value
		}
		if statements[0].cycles == 0 {
			statements, _ = statements.Dequeue()
		}
	}

	// Part 1
	fmt.Println("Signal Strength Sum:", signalStrengthSum)

	// Part 2
	for _, row := range crt {
		fmt.Println(row)
	}
}
