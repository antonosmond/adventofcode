package main

import (
	"adventofcode/pkg/input"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type DiagnosticReport struct {
	bitLength   int
	epsilonRate *int64
	gammaRate   *int64
	values      []int
}

func main() {
	in := input.Load()
	report := toDiagnosticReport(in)
	gamma := report.GammaRate()
	fmt.Printf("Gamma Rate\t\t: %d\n", gamma)
	epsilon := report.EpsilonRate()
	fmt.Printf("Epsilon Rate\t\t: %d\n", epsilon)
	fmt.Printf("Power Consumption\t: %d\n", gamma*epsilon)
	oxygen := report.OxygenGeneratorRating()
	fmt.Printf("Oxygen Generator Rating\t: %d\n", oxygen)
	co2 := report.CO2ScrubberRating()
	fmt.Printf("CO2 Scrubber Rating\t: %d\n", co2)
	fmt.Printf("Life Support Rating\t: %d\n", oxygen*co2)
}

func toDiagnosticReport(in []string) DiagnosticReport {
	report := DiagnosticReport{}
	report.bitLength = len(in[0])
	for _, s := range in {
		v, err := strconv.ParseInt(s, 2, 32)
		if err != nil {
			panic(err)
		}
		report.values = append(report.values, int(v))
	}
	sort.Ints(report.values)
	return report
}

func (r *DiagnosticReport) GammaRate() int64 {
	if r.gammaRate != nil {
		return *r.gammaRate
	}
	g := ""
	for i := 0; i < r.bitLength; i++ {
		bit := "0"
		if commonBit(r.values, i) {
			bit = "1"
		}
		g = bit + g
	}
	gamma, err := strconv.ParseInt(g, 2, 64)
	if err != nil {
		panic(err)
	}
	r.gammaRate = &gamma
	return *r.gammaRate
}

func (r *DiagnosticReport) EpsilonRate() int64 {
	if r.epsilonRate != nil {
		return *r.epsilonRate
	}
	// The epsilon rate is the inverse of the gamma rate i.e. the gamma rate
	// as a binary with all the bits flipped
	e := bitFlip(r.GammaRate(), r.bitLength)
	r.epsilonRate = &e
	return *r.epsilonRate
}

func (r *DiagnosticReport) OxygenGeneratorRating() int {
	list := r.values
	for i := r.bitLength - 1; i >= 0; i-- {
		if len(list) == 1 {
			break
		}
		list = filter(list, i, commonBit(list, i))
	}
	return list[0]
}

func (r *DiagnosticReport) CO2ScrubberRating() int {
	list := r.values
	for i := r.bitLength - 1; i >= 0; i-- {
		if len(list) == 1 {
			break
		}
		list = filter(list, i, !commonBit(list, i))
	}
	return list[0]
}

func on(i int, pos int) bool {
	return i&(1<<pos) > 0
}

// Flipping the bits can be achieved
// by XORing each bit with 1
func bitFlip(i int64, bitLength int) int64 {
	bits := strings.Repeat("1", bitLength)
	b, err := strconv.ParseInt(bits, 2, 32)
	if err != nil {
		panic(err)
	}
	return i ^ b
}

func commonBit(in []int, pos int) bool {
	count := 0
	for _, v := range in {
		if on(int(v), pos) {
			count++
		}
	}
	return count >= len(in)-count
}

func filter(in []int, pos int, bit bool) []int {
	filtered := []int{}
	for _, v := range in {
		if on(v, pos) == bit {
			filtered = append(filtered, v)
		}
	}
	return filtered
}
