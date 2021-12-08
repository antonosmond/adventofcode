package main

import (
	"adventofcode/pkg/input"
	"fmt"
	"strconv"
	"strings"
)

type Fish int8

var fish []*Fish

func (f *Fish) Spawn() {
	*f--
	if *f < 0 {
		*f = 6
		fish = append(fish, NewFish())
	}
}

func main() {
	in := input.Load()
	toFish(in)
	fmt.Println(Simulate(256))
}

func toFish(in []string) {
	for _, line := range in {
		if line == "" {
			continue
		}
		s := strings.Split(line, ",")
		for _, num := range s {
			i, err := strconv.ParseInt(num, 10, 8)
			if err != nil {
				panic(err)
			}
			f := Fish(i)
			fish = append(fish, &f)
		}
	}
}

func Simulate(days int) int {
	// fmt.Printf("Initial state: ")
	// for _, f := range fish {
	// 	fmt.Printf("%d,", *f)
	// }
	// fmt.Println()
	for d := 0; d < days; d++ {
		countBefore := len(fish)
		fmt.Printf("Day %d: ", d)
		for _, f := range fish {
			f.Spawn()
		}
		countAfter := len(fish)
		diff := countAfter - countBefore
		fmt.Printf("%d (+%d)\n", countAfter, diff)
	}
	return len(fish)
}

func NewFish() *Fish {
	f := Fish(8)
	return &f
}
