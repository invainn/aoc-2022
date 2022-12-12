package main

import (
	"fmt"
	"os"
	// "strconv"
	"strings"
)

type Coords struct {
	x int
	y int
}

func min(nums ...int) int {
	minNum := 100000000000000000

	for _, num := range nums {
		if num < minNum {
			minNum = num
		}
	}

	return minNum
}

func DeltaAtMostOne(src, dst int) bool {
	dXY := dst - src

	return dXY < 2
}

func GetValidCardinalDirections(pos Coords, visited map[Coords]Coords, heightMap [][]int) []Coords {
	validDirections := []Coords{}
	currentPosChar := heightMap[pos.y][pos.x]

	southPosition := Coords{pos.x, pos.y + 1}
	if southPosition.y < len(heightMap) && DeltaAtMostOne(currentPosChar, heightMap[southPosition.y][southPosition.x]) {
		if _, ok := visited[southPosition]; !ok {
			validDirections = append(validDirections, southPosition)
		}
	}

	northPosition := Coords{pos.x, pos.y - 1}
	if northPosition.y >= 0 && DeltaAtMostOne(currentPosChar, heightMap[northPosition.y][northPosition.x]) {
		if _, ok := visited[northPosition]; !ok {
			validDirections = append(validDirections, northPosition)
		}
	}

	eastPosition := Coords{pos.x + 1, pos.y}
	if eastPosition.x < len(heightMap[0]) && DeltaAtMostOne(currentPosChar, heightMap[eastPosition.y][eastPosition.x]) {
		if _, ok := visited[eastPosition]; !ok {
			validDirections = append(validDirections, eastPosition)
		}
	}

	westPosition := Coords{pos.x - 1, pos.y}
	if westPosition.x >= 0 && DeltaAtMostOne(currentPosChar, heightMap[westPosition.y][westPosition.x]) {
		if _, ok := visited[westPosition]; !ok {
			validDirections = append(validDirections, westPosition)
		}
	}

	return validDirections
}

func FindShortestPathToEndBFS(startPos Coords, endPos Coords, heightMap [][]int) int {
	nextPos := Queue[Coords]{startPos}
	visited := make(map[Coords]Coords)
	InQueue := make(map[Coords]bool)

	for !nextPos.IsEmpty() {
		newQueue, currentPos := nextPos.Dequeue()
		nextPos = newQueue

		if currentPos == endPos {
			break
		}

		validDirections := GetValidCardinalDirections(currentPos, visited, heightMap)
		for _, dir := range validDirections {
			if _, ok := InQueue[dir]; !ok {
				visited[dir] = Coords{currentPos.x, currentPos.y}

				nextPos = nextPos.Enqueue(dir)
				InQueue[dir] = true
			}
		}
	}

	// Backtrack
	stepsTaken := 0
	currentPos := endPos
	for currentPos != startPos {
		currentPos = visited[currentPos]
		stepsTaken += 1
	}

	return stepsTaken
}

func DeltaAtMostOneForValue(dst, src int) bool {
	dXY := dst - src

	return dXY < 2
}

func GetValidCardinalDirectionsForValue(pos Coords, visited map[Coords]Coords, heightMap [][]int) []Coords {
	validDirections := []Coords{}
	currentPosChar := heightMap[pos.y][pos.x]

	southPosition := Coords{pos.x, pos.y + 1}
	if southPosition.y < len(heightMap) && DeltaAtMostOneForValue(currentPosChar, heightMap[southPosition.y][southPosition.x]) {
		if _, ok := visited[southPosition]; !ok {
			validDirections = append(validDirections, southPosition)
		}
	}

	northPosition := Coords{pos.x, pos.y - 1}
	if northPosition.y >= 0 && DeltaAtMostOneForValue(currentPosChar, heightMap[northPosition.y][northPosition.x]) {
		if _, ok := visited[northPosition]; !ok {
			validDirections = append(validDirections, northPosition)
		}
	}

	eastPosition := Coords{pos.x + 1, pos.y}
	if eastPosition.x < len(heightMap[0]) && DeltaAtMostOneForValue(currentPosChar, heightMap[eastPosition.y][eastPosition.x]) {
		if _, ok := visited[eastPosition]; !ok {
			validDirections = append(validDirections, eastPosition)
		}
	}

	westPosition := Coords{pos.x - 1, pos.y}
	if westPosition.x >= 0 && DeltaAtMostOneForValue(currentPosChar, heightMap[westPosition.y][westPosition.x]) {
		if _, ok := visited[westPosition]; !ok {
			validDirections = append(validDirections, westPosition)
		}
	}

	return validDirections
}

func FindShortestPathToValue(startPos Coords, value int, heightMap [][]int) int {
	nextPos := Queue[Coords]{startPos}
	visited := make(map[Coords]Coords)
	InQueue := make(map[Coords]bool)

	endPos := Coords{0, 0}
	for !nextPos.IsEmpty() {
		newQueue, currentPos := nextPos.Dequeue()
		nextPos = newQueue

		if heightMap[currentPos.y][currentPos.x] == value {
			endPos = currentPos
			break
		}

		validDirections := GetValidCardinalDirectionsForValue(currentPos, visited, heightMap)
		for _, dir := range validDirections {
			if _, ok := InQueue[dir]; !ok {
				visited[dir] = Coords{currentPos.x, currentPos.y}

				nextPos = nextPos.Enqueue(dir)
				InQueue[dir] = true
			}
		}
	}

	// Backtrack
	stepsTaken := 0
	currentPos := endPos
	for currentPos != startPos {
		currentPos = visited[currentPos]
		stepsTaken += 1
	}

	return stepsTaken
}

func FindShortestPathToEndDFS(startPos Coords, endPos Coords, heightMap [][]int, visited map[Coords]bool) int {
	visited[startPos] = true
	if startPos == endPos {
		return 0
	}

	nextPositions := []Coords{}
	currentPosChar := heightMap[startPos.y][startPos.x]

	southPosition := Coords{startPos.x, startPos.y + 1}
	if southPosition.y < len(heightMap) && DeltaAtMostOne(currentPosChar, heightMap[southPosition.y][southPosition.x]) {
		if _, ok := visited[southPosition]; !ok {
			nextPositions = append(nextPositions, southPosition)
		}
	}

	northPosition := Coords{startPos.x, startPos.y - 1}
	if northPosition.y >= 0 && DeltaAtMostOne(currentPosChar, heightMap[northPosition.y][northPosition.x]) {
		if _, ok := visited[northPosition]; !ok {
			nextPositions = append(nextPositions, northPosition)
		}
	}

	eastPosition := Coords{startPos.x + 1, startPos.y}
	if eastPosition.x < len(heightMap[0]) && DeltaAtMostOne(currentPosChar, heightMap[eastPosition.y][eastPosition.x]) {
		if _, ok := visited[eastPosition]; !ok {
			nextPositions = append(nextPositions, eastPosition)
		}
	}

	westPosition := Coords{startPos.x - 1, startPos.y}
	if westPosition.x >= 0 && DeltaAtMostOne(currentPosChar, heightMap[westPosition.y][westPosition.x]) {
		if _, ok := visited[westPosition]; !ok {
			nextPositions = append(nextPositions, westPosition)
		}
	}

	calculatedPositions := []int{}
	shortPathChan := make(chan int)
	for _, pos := range nextPositions {
		copiedMap := make(map[Coords]bool)
		for k, v := range visited {
			copiedMap[k] = v
		}
		go func(pos Coords, endPos Coords, heightMap [][]int, copiedMap map[Coords]bool) {
			shortestPath := 1 + FindShortestPathToEndDFS(pos, endPos, heightMap, copiedMap)
			shortPathChan <- shortestPath
		}(pos, endPos, heightMap, copiedMap)
	}

	for i := 0; i < len(nextPositions); i++ {
		shortestPath := <-shortPathChan
		calculatedPositions = append(calculatedPositions, shortestPath)
	}

	close(shortPathChan)

	return min(calculatedPositions...)
}

func main() {
	dat, err := os.ReadFile("input")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(dat), "\n")
	lines = lines[:len(lines)-1]

	heightMap := make([][]int, len(lines))
	for y := range lines {
		heightMap[y] = make([]int, len(lines[0]))
	}

	startCoords := Coords{0, 0}
	endCoords := Coords{0, 0}

	for y, row := range lines {
		for x, col := range row {
			if col == 'S' {
				startCoords = Coords{x, y}
				heightMap[y][x] = int('a')
			} else if col == 'E' {
				endCoords = Coords{x, y}
				heightMap[y][x] = int('z') + 1
			} else {
				heightMap[y][x] = int(col)
			}
		}
	}

	fmt.Println(FindShortestPathToEndBFS(startCoords, endCoords, heightMap))
	fmt.Println(FindShortestPathToValue(endCoords, int('a'), heightMap))
}
