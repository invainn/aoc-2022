package main

import (
	"container/heap"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type ElfInventory struct {
	calorieCounts []int
}

func (ei *ElfInventory) Total() int {
	total := 0

	for _, calorieCount := range ei.calorieCounts {
		total = total + calorieCount
	}

	return total
}

func ReadInput(filename string) []string {
	dat, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	return strings.Split(string(dat), "\n")
}

func main() {
	lines := ReadInput("input")

	intHeap := &IntHeap{}
	heap.Init(intHeap)

	tempElfInv := ElfInventory{}

	for _, line := range lines {
		if line == "" {
			heap.Push(intHeap, tempElfInv.Total())
			tempElfInv = ElfInventory{}
		} else {
			num, err := strconv.Atoi(line)
			if err != nil {
				panic(err)
			}
			tempElfInv.calorieCounts = append(tempElfInv.calorieCounts, num)
		}
	}

	totalCaloriesOfTopElf := heap.Pop(intHeap).(int)
	totalCaloriesOfTopThreeElves := totalCaloriesOfTopElf + heap.Pop(intHeap).(int) + heap.Pop(intHeap).(int)
	fmt.Printf("Total calories of the top elf with the most calories: %d\n", totalCaloriesOfTopElf)
	fmt.Printf("Total calories of the three elves with the most calories: %d\n", totalCaloriesOfTopThreeElves)
}
