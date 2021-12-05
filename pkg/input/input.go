package input

import (
	"bufio"
	"os"
	"path/filepath"
)

func Load(in ...string) []string {
	path := "input.txt"
	if len(in) > 0 {
		path = in[0]
	}
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	file, err := os.Open(filepath.Join(wd, path))
	if err != nil {
		panic(err)
	}
	defer file.Close()
	inputs := []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		inputs = append(inputs, scanner.Text())
	}
	return inputs
}
