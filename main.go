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
	depth := 0
	count := 0
	for idx, i := range in {
		// only check for increase if this is NOT the first value
		if idx != 0 {
			if i > depth {
				count++
			}
		}
		depth = i
	}
	fmt.Println(count)
}
