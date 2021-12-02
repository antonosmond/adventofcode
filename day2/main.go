package main

import (
	"adventofcode/pkg/input"
	"fmt"
)

type Location struct {
	Horizontal int
	Depth      int
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
	loc := &Location{0, 0}
	for _, i := range in {
		switch i.Action {
		case "forward":
			loc.Horizontal += i.Value
		case "up":
			loc.Depth -= i.Value
		case "down":
			loc.Depth += i.Value
		}
	}
	return loc
}
