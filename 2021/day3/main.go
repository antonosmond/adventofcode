package main

import (
	"adventofcode/pkg/input"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	in := input.Load().ToString()
	data := [][]string{}
	for _, s := range in {
		data = append(data, strings.Split(s, ""))
	}
	transposed := transpose(data)
	newData := []string{}
	for _, t := range transposed {
		newData = append(newData, strings.Join(t, ""))
	}
	gammaRateB := ""
	epsilonRateB := ""
	for _, i := range newData {
		on := strings.Count(i, "1")
		off := strings.Count(i, "0")
		if on > off {
			gammaRateB += "1"
			epsilonRateB += "0"
		} else {
			gammaRateB += "0"
			epsilonRateB += "1"
		}
	}
	gammaRate, err := strconv.ParseInt(gammaRateB, 2, 64)
	if err != nil {
		panic(err)
	}
	epsilonRate, err := strconv.ParseInt(epsilonRateB, 2, 64)
	if err != nil {
		panic(err)
	}
	fmt.Println(gammaRate * epsilonRate)
}

func transpose(slice [][]string) [][]string {
	xl := len(slice[0])
	yl := len(slice)
	result := make([][]string, xl)
	for i := range result {
		result[i] = make([]string, yl)
	}
	for i := 0; i < xl; i++ {
		for j := 0; j < yl; j++ {
			result[i][j] = slice[j][i]
		}
	}
	return result
}
