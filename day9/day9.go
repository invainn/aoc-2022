package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

var DIRECTION_MAP = map[string]Coords{
	"R": {1, 0},
	"L": {-1, 0},
	"U": {0, 1},
	"D": {0, -1},
}

type Coords struct {
	x int
	y int
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}

func (c1 Coords) Add(c2 Coords) Coords {
	return Coords{
		x: c1.x + c2.x,
		y: c1.y + c2.y,
	}
}

func (c1 Coords) IsEqualTo(c2 Coords) bool {
	return c1.x == c2.x && c1.y == c2.y
}

func (c1 Coords) IsNextTo(c2 Coords) bool {
	adjacentPositions := []Coords{
		c1.Add(DIRECTION_MAP["U"].Add(DIRECTION_MAP["L"])),
		c1.Add(DIRECTION_MAP["U"]),
		c1.Add(DIRECTION_MAP["U"].Add(DIRECTION_MAP["R"])),
		c1.Add(DIRECTION_MAP["R"]),
		c1.Add(DIRECTION_MAP["D"].Add(DIRECTION_MAP["R"])),
		c1.Add(DIRECTION_MAP["D"]),
		c1.Add(DIRECTION_MAP["D"].Add(DIRECTION_MAP["L"])),
		c1.Add(DIRECTION_MAP["L"]),
		c1,
	}

	for _, pos := range adjacentPositions {
		if c2.IsEqualTo(pos) {
			return true
		}
	}

	return false
}

func getMagnitude(x int) int {
	if x == 0 {
		return 0
	}

	return 1
}

func main() {
	dat, err := os.ReadFile("input")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(dat), "\n")
	lines = lines[:len(lines)-1]

	// Part 1
	head := Coords{0, 0}
	tail := Coords{0, 0}

	lastHeadPosition := Coords{0, 0}
	tailPositionsVisited := map[Coords]bool{{0, 0}: true}
	for _, line := range lines {
		move := strings.Split(line, " ")

		direction := move[0]
		repeatTimes, _ := strconv.Atoi(move[1])

		for i := 0; i < repeatTimes; i++ {
			lastHeadPosition = Coords{head.x, head.y}
			head = head.Add(DIRECTION_MAP[direction])

			if !tail.IsNextTo(head) {
				tail = Coords{lastHeadPosition.x, lastHeadPosition.y}
				if _, ok := tailPositionsVisited[tail]; !ok {
					tailPositionsVisited[tail] = true
				}
			}
		}
	}

	fmt.Println("Number of positions tail visited:", len(tailPositionsVisited))

	// Part 2
	knots := []Coords{}
	for i := 0; i < 10; i++ {
		knots = append(knots, Coords{0, 0})
	}

	tailPositionsVisited = map[Coords]bool{{0, 0}: true}
	orderedVisited := []Coords{{0, 0}}
	for _, line := range lines {
		move := strings.Split(line, " ")

		direction := move[0]
		repeatTimes, _ := strconv.Atoi(move[1])

		for i := 0; i < repeatTimes; i++ {
			for knotIdx := 0; knotIdx < len(knots); knotIdx++ {
				if knotIdx == 0 {
					knots[knotIdx] = knots[knotIdx].Add(DIRECTION_MAP[direction])
				} else {
					dX := knots[knotIdx-1].x - knots[knotIdx].x
					dY := knots[knotIdx-1].y - knots[knotIdx].y

					posToMoveTo := Coords{
						x: int(math.Copysign(float64(getMagnitude(dX)), float64(dX))),
						y: int(math.Copysign(float64(getMagnitude(dY)), float64(dY))),
					}

					if Abs(dX) > 1 || Abs(dY) > 1 {
						knots[knotIdx] = knots[knotIdx].Add(posToMoveTo)
					}
				}
			}
			if _, ok := tailPositionsVisited[knots[len(knots)-1]]; !ok {
				tailPositionsVisited[knots[len(knots)-1]] = true
				orderedVisited = append(orderedVisited, knots[len(knots)-1])
			}
		}
	}
	fmt.Println("Number of positions index 9 tail visited:", len(tailPositionsVisited))

}
