package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func getInput(path string) ([]int, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	var depths []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		i, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return nil, err
		}
		depths = append(depths, i)
	}
	return depths, nil
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Expected 1 arg: path to input")
		os.Exit(1)
	}
	in, err := getInput(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	windows := createSlidingWindows(in, 3)
	increases := countIncreases(windows)
	fmt.Printf("Result: %d\n", increases)
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
