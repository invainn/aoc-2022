package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"
)

func GetRuneValue(r rune) int {
	if unicode.IsUpper(r) {
		return int(r) - 65 + 27
	} else if unicode.IsLower(r) {
		return int(r) - 97 + 1
	}

	panic(r)
}

func FindRucksackCommonPriorities(rucksacks []string) int {
	sumOfCommonPriorities := 0
	for _, rucksack := range rucksacks {
		set := make(map[byte]bool)
		rucksackLength := len(rucksack) / 2

		for i := 0; i < rucksackLength; i++ {
			set[rucksack[i]] = true
		}

		for i := rucksackLength; i < rucksackLength*2; i++ {
			if _, ok := set[rucksack[i]]; ok {
				sumOfCommonPriorities += GetRuneValue(rune(rucksack[i]))
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
		firstSet := make(map[byte]bool)
		for i := 0; i < len(rucksacks[0]); i++ {
			firstSet[rucksacks[0][i]] = true
		}

		secondSet := make(map[byte]bool)
		for i := 0; i < len(rucksacks[1]); i++ {
			secondSet[rucksacks[1][i]] = true
		}

		for i := 0; i < len(rucksacks[2]); i++ {
			char := rucksacks[2][i]
			if _, firstOk := firstSet[char]; firstOk {
				if _, secondOk := secondSet[char]; secondOk {
					sumOfCommonPriorities += GetRuneValue(rune(char))
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
