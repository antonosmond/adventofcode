package main

import (
	"adventofcode/pkg/input"
	"fmt"
)

type Location struct {
	Aim        int
	Depth      int
	Horizontal int
}

func (l *Location) Total() int {
	return l.Horizontal * l.Depth
}

func main() {
	in := input.Load().ToCommands()
	location := sail(in)
	fmt.Println(location.Total())
}

func sail(in []input.Command) *Location {
	loc := &Location{0, 0, 0}
	for _, i := range in {
		switch i.Action {
		case "forward":
			loc.Horizontal += i.Value
			loc.Depth += (i.Value * loc.Aim)
		case "up":
			loc.Aim -= i.Value
		case "down":
			loc.Aim += i.Value
		}
	}
	return loc
}
