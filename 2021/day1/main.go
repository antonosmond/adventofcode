package main

import (
	"adventofcode/pkg/input"
	"fmt"
)

func main() {
	in := input.Load().ToInt()
	windows := createSlidingWindows(in, 3)
	result := countIncreases(windows)
	fmt.Println(result)
}

func createSlidingWindows(in []int, windowSize int) []int {
	windows := []int{}
	for start, _ := range in {
		sum := 0
		end := start + windowSize
		for end > len(in) {
			end--
		}
		for _, ii := range in[start:end] {
			sum += ii
		}
		windows = append(windows, sum)
	}
	return windows
}

func countIncreases(in []int) int {
	depth := 0
	count := 0
	for idx, i := range in {
		if idx != 0 {
			if i > depth {
				count++
			}
		}
		depth = i
	}
	return count
}
