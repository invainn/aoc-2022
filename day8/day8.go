package main

import (
	"fmt"
	"os"
	"strings"
)

type Direction int

// I'm pretty sure there's a smarter way to do this, I just don't have it on the top of my head
func CheckVisibilityCardinalDirections(x int, y int, trees [][]string) bool {
	colLength := len(trees)
	rowLength := len(trees[0])

	treeToCheckHeight := trees[x][y]

	// check visibility in north direction
	IsVisibleFromNorth := true
	for i := 0; i < x; i++ {
		if trees[i][y] >= treeToCheckHeight {
			IsVisibleFromNorth = false
		}
	}
	// check visibility in south direction
	IsVisibleFromSouth := true
	for i := x + 1; i < colLength; i++ {
		if trees[i][y] >= treeToCheckHeight {
			IsVisibleFromSouth = false
		}
	}
	// check visibility in north direction
	IsVisibleFromWest := true
	for i := 0; i < y; i++ {
		if trees[x][i] >= treeToCheckHeight {
			IsVisibleFromWest = false
		}
	}
	// check visibility in north direction
	IsVisibleFromEast := true
	for i := y + 1; i < rowLength; i++ {
		if trees[x][i] >= treeToCheckHeight {
			IsVisibleFromEast = false
		}
	}

	return IsVisibleFromNorth || IsVisibleFromSouth || IsVisibleFromWest || IsVisibleFromEast
}

func CalculateScenicScore(x int, y int, trees [][]string) int {
	colLength := len(trees)
	rowLength := len(trees[0])

	treeToCheckHeight := trees[x][y]

	northScenicScore := 0
	for i := x - 1; i >= 0; i-- {
		northScenicScore += 1
		if trees[i][y] >= treeToCheckHeight {
			break
		}
	}

	southScenicScore := 0
	for i := x + 1; i < colLength; i++ {
		southScenicScore += 1
		if trees[i][y] >= treeToCheckHeight {
			break
		}
	}

	westScenicScore := 0
	for i := y - 1; i >= 0; i-- {
		westScenicScore += 1
		if trees[x][i] >= treeToCheckHeight {
			break
		}
	}

	eastScenicScore := 0
	for i := y + 1; i < rowLength; i++ {
		eastScenicScore += 1
		if trees[x][i] >= treeToCheckHeight {
			break
		}
	}

	return northScenicScore * southScenicScore * westScenicScore * eastScenicScore
}

func main() {
	dat, err := os.ReadFile("input")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(dat), "\n")
	lines = lines[:len(lines)-1]

	trees := [][]string{}
	for _, line := range lines {
		trees = append(trees, strings.Split(line, ""))
	}

	treesVisible := 0

	for i := 0; i < len(trees); i++ {
		for j := 0; j < len(trees[0]); j++ {
			if i == 0 || i == len(trees)-1 || j == 0 || j == len(trees[0])-1 || CheckVisibilityCardinalDirections(i, j, trees) {
				treesVisible += 1
			}
		}
	}
	// Part 1
	fmt.Println("Trees visible in grid:", treesVisible)

	maxScenicScore := 0
	for i := 0; i < len(trees); i++ {
		for j := 0; j < len(trees[0]); j++ {
			if i == 0 || i == len(trees)-1 || j == 0 || j == len(trees[0])-1 {
				continue
			}
			scenicScore := CalculateScenicScore(i, j, trees)
			if scenicScore > maxScenicScore {
				maxScenicScore = scenicScore
			}
		}
	}
	// Part 2
	fmt.Println("Max scenic score for a tree:", maxScenicScore)
}
