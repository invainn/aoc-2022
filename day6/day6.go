package main

import (
	"fmt"
	"os"
	"strings"
)

func FindDistinctCharsSlidingWindow(slidingWindowLength int, datastream []string) int {
	queue := Queue{}
	for i, char := range datastream {
		if i >= slidingWindowLength {
			if !queue.ContainsDuplicates() {
				return i
			}

			queue = queue.Enqueue(char)
			queue = queue.Dequeue()
		} else {
			queue = queue.Enqueue(char)
		}
	}
	panic("could not find distinct chars with sliding window length")
}

func main() {
	dat, err := os.ReadFile("input")
	if err != nil {
		panic(err)
	}

	line := strings.Split(string(dat), "\n")[0]
	datastream := strings.Split(line, "")
	fmt.Println("First start of packet is:", FindDistinctCharsSlidingWindow(4, datastream))
	fmt.Println("First start of message is:", FindDistinctCharsSlidingWindow(14, datastream))
}
