package main

import (
	"adventofcode/pkg/input"
	"fmt"
	"strconv"
	"strings"
)

type Command struct {
	Action string
	Value  int
}

type Location struct {
	Aim        int
	Depth      int
	Horizontal int
}

func (l *Location) Total() int {
	return l.Horizontal * l.Depth
}

func main() {
	in := input.Load()
	commands := toCommands(in)
	location := sail(commands)
	fmt.Println(location.Total())
}

func sail(in []Command) *Location {
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

func toCommands(in []string) []Command {
	commands := []Command{}
	for _, s := range in {
		parts := strings.Split(s, " ")
		value, err := strconv.Atoi(parts[1])
		if err != nil {
			panic(err)
		}
		commands = append(commands, Command{
			Action: parts[0],
			Value:  value,
		})
	}
	return commands
}
