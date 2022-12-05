package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Command struct {
	take int
	from int
	to   int
}

func (c Command) OperateOnStacks(stacks []Stack) []Stack {
	for i := 0; i < c.take; i++ {
		poppedStack, poppedValue := stacks[c.from-1].Pop()
		pushedStack := stacks[c.to-1].Push(poppedValue)

		// reconcile
		stacks[c.from-1] = poppedStack
		stacks[c.to-1] = pushedStack
	}

	return stacks
}

func (c Command) OperateOnStacks9001(stacks []Stack) []Stack {
	poppedValues := []string{}
	for i := 0; i < c.take; i++ {
		poppedStack, poppedValue := stacks[c.from-1].Pop()
		poppedValues = append(poppedValues, poppedValue)

		// reconcile
		stacks[c.from-1] = poppedStack
	}

	poppedValues = reverse(poppedValues)

	for _, crate := range poppedValues {
		pushedStack := stacks[c.to-1].Push(crate)

		// reconcile
		stacks[c.to-1] = pushedStack
	}

	return stacks
}

func reverse(s []string) []string {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}

	return s
}

func transpose(m [][]string) [][]string {
	maxRowLen := 0
	// find max length in array of slices
	for _, s := range m {
		if len(s) > maxRowLen {
			maxRowLen = len(s)
		}
	}

	result := make([][]string, maxRowLen)
	for i := range result {
		result[i] = make([]string, len(m))
	}

	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m[i]); j++ {
			result[j][i] = m[i][j]
		}
	}

	return result
}

func main() {
	dat, err := os.ReadFile("input")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(dat), "\n")
	crateIdx := 0
	for idx, line := range lines {
		if line == "" {
			crateIdx = idx
			break
		}
	}

	crates := lines[:crateIdx]
	rawCommands := lines[crateIdx+1:]
	rawCommands = rawCommands[:len(rawCommands)-1]

	crateSlices := [][]string{}
	for _, crate := range crates[:len(crates)-1] {
		crateSlices = append(crateSlices, strings.Split(crate, ""))
	}

	transposedCrateSlices := transpose(crateSlices)
	re, _ := regexp.Compile(`[\[\]\s+]`)
	transposedCrates := []string{}
	for _, slice := range transposedCrateSlices {
		newString := re.ReplaceAllString(strings.Join(slice, ""), "")
		if newString != "" {
			transposedCrates = append(transposedCrates, newString)
		}
	}

	cratesInStacks := []Stack{}
	for _, crate := range transposedCrates {
		stack := Stack.StackFromStrings(Stack{}, reverse(strings.Split(crate, "")))
		cratesInStacks = append(cratesInStacks, stack)
	}

	commands := []Command{}
	re, _ = regexp.Compile(`[0-9]+`)
	for _, rawCommand := range rawCommands {
		strs := re.FindAllString(rawCommand, -1)
		take, _ := strconv.Atoi(strs[0])
		from, _ := strconv.Atoi(strs[1])
		to, _ := strconv.Atoi(strs[2])

		commands = append(commands, Command{take, from, to})
	}

	// grab top of stacks and print out string
	processedCrates := make([]Stack, len(cratesInStacks))
	copy(processedCrates, cratesInStacks)
	for _, command := range commands {
		processedCrates = command.OperateOnStacks(processedCrates)
	}

	cratesOnTop := []string{}
	for _, stk := range processedCrates {
		_, crateOnTop := stk.Pop()
		cratesOnTop = append(cratesOnTop, crateOnTop)
	}

	fmt.Printf("Crates that were on top were: %s\n", strings.Join(cratesOnTop, ""))

	// grab top of stacks and print out string processed by 9001
	processedCrates9001 := make([]Stack, len(cratesInStacks))
	copy(processedCrates9001, cratesInStacks)
	for _, command := range commands {
		processedCrates9001 = command.OperateOnStacks9001(processedCrates9001)
	}

	cratesOnTop9001 := []string{}
	for _, stk := range processedCrates9001 {
		_, crateOnTop := stk.Pop()
		cratesOnTop9001 = append(cratesOnTop9001, crateOnTop)
	}

	fmt.Printf("Crates that were on top after CrateMover9001 did it's job: %s\n", strings.Join(cratesOnTop9001, ""))
}
