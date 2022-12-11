package main

import (
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type Operator int

const (
	MULT Operator = iota
	ADD  Operator = iota
)

type Monkey struct {
	items []int

	operand  int
	operator Operator

	testNum        int
	testTrueIndex  int
	testFalseIndex int
}

func CalculateMonkeyBusiness(monkeys []Monkey, rounds int, divByThree bool) {
	productOfDivisors := 1
	for _, monkey := range monkeys {
		productOfDivisors *= monkey.testNum
	}

	monkeyInspections := map[int]int{}

	for round := 0; round < rounds; round++ {
		for monkeyIdx := 0; monkeyIdx < len(monkeys); monkeyIdx++ {
			for _, item := range monkeys[monkeyIdx].items {
				monkeyInspections[monkeyIdx] += 1

				operand := monkeys[monkeyIdx].operand
				if operand == 0 {
					operand = item
				}

				newWorryLevel := 0
				if monkeys[monkeyIdx].operator == MULT {
					newWorryLevel = (item * operand) % productOfDivisors
				} else {
					newWorryLevel = (item + operand) % productOfDivisors
				}

				if divByThree {
					newWorryLevel /= 3
				} else {
					newWorryLevel %= productOfDivisors
				}

				if (newWorryLevel % monkeys[monkeyIdx].testNum) == 0 {
					monkeys[monkeys[monkeyIdx].testTrueIndex].items = append(monkeys[monkeys[monkeyIdx].testTrueIndex].items, newWorryLevel)
				} else {
					monkeys[monkeys[monkeyIdx].testFalseIndex].items = append(monkeys[monkeys[monkeyIdx].testFalseIndex].items, newWorryLevel)
				}

				monkeys[monkeyIdx].items = monkeys[monkeyIdx].items[1:]
			}
		}
	}
	monkeyInspectionList := []int{}
	for _, num := range monkeyInspections {
		monkeyInspectionList = append(monkeyInspectionList, num)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(monkeyInspectionList)))
	fmt.Println("Level of monkey business:", monkeyInspectionList[0]*monkeyInspectionList[1])
}

func main() {
	dat, err := os.ReadFile("input")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(dat), "\n")

	monkeys := []Monkey{}

	// Parsing
	for idx, line := range lines {
		re := regexp.MustCompile(`[0-9]+`)
		if strings.Contains(line, "Monkey") {
			monkey := Monkey{
				[]int{},
				0,
				0,
				0,
				0,
				0,
			}

			// items
			itemsLine := lines[idx+1]
			items := re.FindAllString(itemsLine, -1)
			for _, item := range items {
				itemNum, _ := strconv.Atoi(item)
				monkey.items = append(monkey.items, itemNum)
			}

			// operation
			operationLine := lines[idx+2]
			if strings.Contains(operationLine, "old * old") {
				monkey.operand = 0
				monkey.operator = MULT
			} else {
				operandStr := re.FindAllString(operationLine, -1)[0]
				operand, _ := strconv.Atoi(operandStr)
				monkey.operand = operand

				operatorRe := regexp.MustCompile(`[*|+]{1}`)
				operator := operatorRe.FindAllString(operationLine, -1)[0]

				if operator == "*" {
					monkey.operator = MULT
				} else {
					monkey.operator = ADD
				}
			}

			// test
			testLine := lines[idx+3]
			testTrueLine := lines[idx+4]
			testFalseLine := lines[idx+5]

			testLineStr := re.FindAllString(testLine, -1)[0]
			testTrueLineStr := re.FindAllString(testTrueLine, -1)[0]
			testFalseLineStr := re.FindAllString(testFalseLine, -1)[0]

			testLineNum, _ := strconv.Atoi(testLineStr)
			testTrueLineNum, _ := strconv.Atoi(testTrueLineStr)
			testFalseLineNum, _ := strconv.Atoi(testFalseLineStr)

			monkey.testNum = testLineNum
			monkey.testTrueIndex = testTrueLineNum
			monkey.testFalseIndex = testFalseLineNum

			monkeys = append(monkeys, monkey)
		}
	}

	monkeysPart1 := make([]Monkey, len(monkeys))
	copy(monkeysPart1, monkeys)
	CalculateMonkeyBusiness(monkeysPart1, 20, true)

	monkeysPart2 := make([]Monkey, len(monkeys))
	copy(monkeysPart2, monkeys)
	CalculateMonkeyBusiness(monkeysPart2, 10000, false)
}
