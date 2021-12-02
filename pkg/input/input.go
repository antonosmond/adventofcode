package input

import (
	"bufio"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type Inputs []string

type Command struct {
	Action string
	Value  int
}

func Load() Inputs {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	file, err := os.Open(filepath.Join(wd, "input"))
	if err != nil {
		panic(err)
	}
	defer file.Close()
	inputs := Inputs{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		inputs = append(inputs, scanner.Text())
	}
	return inputs
}

func (in Inputs) ToInt() []int {
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

func (in Inputs) ToCommands() []Command {
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
