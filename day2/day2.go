package main

import (
	"fmt"
	"os"
	"strings"
)

var WIN_POINTS = map[string]int{
	"AX": 4,
	"BX": 1,
	"CX": 7,

	"AY": 8,
	"BY": 5,
	"CY": 2,

	"AZ": 3,
	"BZ": 9,
	"CZ": 6,
}

var ROUND_POINTS = map[string]int{
	"AX": 3,
	"BX": 1,
	"CX": 2,

	"AY": 4,
	"BY": 5,
	"CY": 6,

	"AZ": 8,
	"BZ": 9,
	"CZ": 7,
}

func main() {
	dat, err := os.ReadFile("input")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(dat), "\n")

	rpsRounds := []string{}
	for _, line := range lines {
		rpsRounds = append(rpsRounds, strings.ReplaceAll(line, " ", ""))
	}

	roundPoints := []int{}
	roundMatchPoints := []int{}
	for _, round := range rpsRounds {
		roundPoints = append(roundPoints, WIN_POINTS[round])
		roundMatchPoints = append(roundMatchPoints, ROUND_POINTS[round])
	}

	totalScore := 0
	for _, round := range roundPoints {
		totalScore = totalScore + round
	}

	totalMatchScore := 0
	for _, round := range roundMatchPoints {
		totalMatchScore = totalMatchScore + round
	}

	fmt.Printf("Total Score using RPS strategy: %d\n", totalScore)
	fmt.Printf("Total Score using WLD strategy: %d\n", totalMatchScore)
}
