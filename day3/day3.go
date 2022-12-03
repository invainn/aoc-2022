package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"
)

func GetRuneValue(char rune) int {
	if unicode.IsUpper(char) {
		return int(char) - 65 + 27
	} else if unicode.IsLower(char) {
		return int(char) - 97 + 1
	}

	panic(char)
}

func FindRucksackCommonPriorities(rucksacks []string) int {
	sumOfCommonPriorities := 0
	for _, rucksack := range rucksacks {
		freqCount := make(map[rune]int)
		rucksackLength := len(rucksack) / 2

		for i := 0; i < rucksackLength; i++ {
			rune := rune(rucksack[i])
			if count, ok := freqCount[rune]; ok {
				freqCount[rune] += count + 1
			} else {
				freqCount[rune] = 1
			}
		}

		for i := rucksackLength; i < rucksackLength*2; i++ {
			rune := rune(rucksack[i])
			if _, ok := freqCount[rune]; ok {
				sumOfCommonPriorities += GetRuneValue(rune)
				break
			}
		}
	}

	return sumOfCommonPriorities
}

func FindRucksackCommonPrioritiesThree(rucksacks []string) int {
	sumOfCommonPriorities := 0
	rucksacksGroupOfThree := [][]string{}
	for i := 0; i < len(rucksacks); i += 3 {
		rucksacksGroupOfThree = append(rucksacksGroupOfThree, []string{
			rucksacks[i],
			rucksacks[i+1],
			rucksacks[i+2],
		})
	}

	for _, rucksacks := range rucksacksGroupOfThree {
		firstFreqCount := make(map[rune]int)
		for i := 0; i < len(rucksacks[0]); i++ {
			rune := rune(rucksacks[0][i])
			if count, ok := firstFreqCount[rune]; ok {
				firstFreqCount[rune] += count + 1
			} else {
				firstFreqCount[rune] = 1
			}
		}

		secondFreqCount := make(map[rune]int)
		for i := 0; i < len(rucksacks[1]); i++ {
			rune := rune(rucksacks[1][i])
			if count, ok := secondFreqCount[rune]; ok {
				secondFreqCount[rune] += count + 1
			} else {
				secondFreqCount[rune] = 1
			}
		}

		for i := 0; i < len(rucksacks[2]); i++ {
			rune := rune(rucksacks[2][i])
			if _, firstOk := firstFreqCount[rune]; firstOk {
				if _, secondOk := secondFreqCount[rune]; secondOk {
					sumOfCommonPriorities += GetRuneValue(rune)
					break
				}
			}
		}
	}

	return sumOfCommonPriorities
}

func main() {
	dat, err := os.ReadFile("input")
	if err != nil {
		panic(err)
	}

	rucksacks := strings.Split(string(dat), "\n")
	rucksacks = rucksacks[:len(rucksacks)-1] // an extra line gets added here if we split by newline

	sumOfCommonPriorities := FindRucksackCommonPriorities(rucksacks)
	fmt.Printf("Sum of common priorities is : %d\n", sumOfCommonPriorities)

	sumOfCommonPriorities = FindRucksackCommonPrioritiesThree(rucksacks)
	fmt.Printf("Sum of common priorities for groups of three is : %d\n", sumOfCommonPriorities)
}
