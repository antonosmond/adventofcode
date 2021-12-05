package main

import (
	"adventofcode/pkg/input"
	"fmt"
	"strconv"
)

func main() {
	in := input.Load()
	ints := toInt(in)
	windows := createSlidingWindows(ints, 3)
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

func toInt(in []string) []int {
	out := []int{}
	for _, s := range in {
		i, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		out = append(out, i)
	}
	return out
}
