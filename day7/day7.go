package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type FileTreeNode struct {
	name     string
	size     int
	children []*FileTreeNode
}

func (ftn FileTreeNode) IsDirectory() bool {
	return ftn.size == 0
}

func (ftn FileTreeNode) String() string {
	return ftn.name
}

func (ftn FileTreeNode) PrintChildren() []string {
	result := []string{}

	for _, child := range ftn.children {
		result = append(result, (*child).name)
	}

	return result
}

func (ftn FileTreeNode) PrintDirectory() {
	for _, ptr := range ftn.children {
		child := *ptr
		if child.IsDirectory() {
			child.PrintDirectory()
		}
	}
}

func (ftn FileTreeNode) GetDirectorySizesHelper(results *[]int) int {
	directorySize := 0
	for _, ptr := range ftn.children {
		child := *ptr
		if child.IsDirectory() {
			childDirectorySize := child.GetDirectorySizesHelper(results)
			*results = append(*results, childDirectorySize)
			directorySize += childDirectorySize
		} else {
			directorySize += child.size
		}
	}
	return directorySize
}

func (ftn FileTreeNode) GetDirectorySizes() []int {
	results := []int{}

	directorySize := 0
	for _, ptr := range ftn.children {
		child := *ptr
		if child.IsDirectory() {
			childDirectorySize := child.GetDirectorySizesHelper(&results)
			results = append(results, childDirectorySize)
			directorySize += childDirectorySize
		} else {
			directorySize += child.size
		}
	}

	return append(results, directorySize)
}

func main() {
	dat, err := os.ReadFile("input")
	if err != nil {
		panic(err)
	}

	directoryTracker := Stack[*FileTreeNode]{}

	lines := strings.Split(string(dat), "\n")
	lines = lines[:len(lines)-1]
	// parse filesystem
	for idx, line := range lines {
		output := strings.Split(line, " ")
		if output[0] == "$" {
			command := output[1]

			if command == "cd" {
				operand := output[2]

				if operand == ".." {
					directoryTracker, _ = directoryTracker.Pop()
				} else {
					newNode := FileTreeNode{
						name: output[2],
						size: 0,
					}

					currentNode := directoryTracker.Peek()
					if currentNode != nil {
						(*currentNode).children = append((*currentNode).children, &newNode)
					}
					directoryTracker = directoryTracker.Push(&newNode)
				}
			}

			if command == "ls" {
				for i := idx + 1; i <= len(lines)-1 && lines[i][0] != '$'; i++ {
					file := strings.Split(lines[i], " ")
					if file[0] != "dir" {
						size, _ := strconv.Atoi(file[0])
						newNode := FileTreeNode{
							name: file[1],
							size: size,
						}

						currentNode := directoryTracker.Peek()
						(*currentNode).children = append((*currentNode).children, &newNode)
					}
				}
			}
		}
	}
	// Go to root
	for len(directoryTracker) != 1 {
		directoryTracker, _ = directoryTracker.Pop()
	}

	// Part 1
	directorySizes := (*directoryTracker.Peek()).GetDirectorySizes()
	sumOfDirectoriesLessThan100000 := 0
	for _, size := range directorySizes {
		if size <= 100000 {
			sumOfDirectoriesLessThan100000 += size
		}
	}

	fmt.Printf("Sum of total sizes of directories that are at most size 100k: %d\n", sumOfDirectoriesLessThan100000)

	// Part 2
	// This can be done much faster, I am just so tired
	unusedSpace := 70000000 - directorySizes[len(directorySizes)-1] // last thing to get appended is the root
	sort.Ints(directorySizes)

	for _, size := range directorySizes {
		if (unusedSpace + size) >= 30000000 {
			fmt.Printf("Size of smallest directory that will fulfill 30m unused space req: %d\n", size)
			break
		}
	}
}
