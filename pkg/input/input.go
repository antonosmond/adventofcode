package input

import (
	"bufio"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

var passwordRegexp = regexp.MustCompile(`^([0-9]+)-([0-9]+) ([a-z]): (.*)`)

type Inputs []string

type Command struct {
	Action string
	Value  int
}

type Password struct {
	Policy PasswordPolicy
	Value  string
}

type PasswordPolicy struct {
	Letter string
	Min    int
	Max    int
}

func Load() Inputs {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	file, err := os.Open(filepath.Join(wd, "input.txt"))
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

func (in Inputs) ToPasswords() []Password {
	passwords := []Password{}
	for _, p := range in {
		groups := passwordRegexp.FindStringSubmatch(p)
		min, err := strconv.Atoi(groups[1])
		if err != nil {
			panic(err)
		}
		max, err := strconv.Atoi(groups[2])
		if err != nil {
			panic(err)
		}
		passwords = append(passwords, Password{
			Policy: PasswordPolicy{
				Min:    min,
				Max:    max,
				Letter: groups[3],
			},
			Value: groups[4],
		})
	}
	return passwords
}
