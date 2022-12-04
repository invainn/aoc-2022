package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Interval struct {
	start int
	end   int
}

func (i1 Interval) IsInsideInterval(i2 Interval) bool {
	return i1.start >= i2.start && i1.end <= i2.end
}

func (i1 Interval) DoesOverlapInterval(i2 Interval) bool {
	return i1.start <= i2.end && i1.end >= i2.start
}

func main() {
	dat, err := os.ReadFile("input")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(dat), "\n")
	lines = lines[:len(lines)-1]

	stringPairs := [][]string{}
	for _, line := range lines {
		stringPairs = append(stringPairs, strings.Split(line, ","))
	}

	pairs := [][]Interval{}
	for _, pair := range stringPairs {
		newPair := []Interval{}
		for _, elf := range pair {
			splitElf := strings.Split(elf, "-")

			start, err := strconv.Atoi(splitElf[0])
			if err != nil {
				panic(err)
			}

			end, err := strconv.Atoi(splitElf[1])
			if err != nil {
				panic(err)
			}

			newPair = append(newPair, Interval{start, end})
		}

		pairs = append(pairs, newPair)
	}

	pairsThatAreInsideAnother := 0
	for _, pair := range pairs {
		if pair[0].IsInsideInterval(pair[1]) || pair[1].IsInsideInterval(pair[0]) {
			pairsThatAreInsideAnother += 1
		}
	}
	fmt.Printf("Number of pairs that have a range inside the other: %d\n", pairsThatAreInsideAnother)

	pairsThatOverlap := 0
	for _, pair := range pairs {
		if pair[0].DoesOverlapInterval(pair[1]) {
			pairsThatOverlap += 1
		}
	}
	fmt.Printf("Number of pairs that overlap: %d\n", pairsThatOverlap)
}
