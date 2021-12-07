package main

import (
	"adventofcode/pkg/input"
	"fmt"
	"math"
	"regexp"
	"strconv"
)

var coordinateSegmentRegexp = regexp.MustCompile("^([0-9]+),([0-9]+) -> ([0-9]+),([0-9]+)")

type Coordinate struct {
	x int
	y int
}

type CoordinateSegment struct {
	Start *Coordinate
	End   *Coordinate
}

func main() {
	in := input.Load()
	cs := toCoordinateSegment(in)
	hydrothermalVentCordinates := map[Coordinate]int{}
	for _, seg := range cs {
		coordinates := seg.plot()
		for _, c := range coordinates {
			hydrothermalVentCordinates[c]++
		}
	}
	count := 0
	for _, v := range hydrothermalVentCordinates {
		if v >= 2 {
			count++
		}
	}
	fmt.Println(count)
}

func toCoordinateSegment(in []string) []CoordinateSegment {
	cs := []CoordinateSegment{}
	for _, i := range in {
		if i == "" {
			continue
		}
		matches := coordinateSegmentRegexp.FindStringSubmatch(i)
		ints := []int{}
		for _, m := range matches[1:] {
			i, _ := strconv.ParseInt(m, 10, 32)
			ints = append(ints, int(i))
		}
		c := CoordinateSegment{
			Start: &Coordinate{
				x: ints[0],
				y: ints[1],
			},
			End: &Coordinate{
				x: ints[2],
				y: ints[3],
			},
		}
		cs = append(cs, c)
	}
	return cs
}

func (c *CoordinateSegment) plot() []Coordinate {

	xShift := 1
	xDiff := c.End.x - c.Start.x
	if xDiff < 0 {
		xShift = -1
	}

	yShift := 1
	yDiff := c.End.y - c.Start.y
	if yDiff < 0 {
		yShift = -1
	}

	// Handle case where start & end could be the same
	if xDiff == 0 && yDiff == 0 {
		return []Coordinate{*c.Start}
	}

	coordinates := []Coordinate{}
	switch true {
	case xDiff != 0 && yDiff != 0:
		// diagonal
		for i := 0; i <= int(math.Abs(float64(xDiff))); i++ {
			coordinates = append(coordinates, Coordinate{
				x: c.Start.x + i*xShift,
				y: c.Start.y + i*yShift,
			})
		}
		break
	case xDiff != 0:
		// horizontal
		for i := 0; i <= int(math.Abs(float64(xDiff))); i++ {
			coordinates = append(coordinates, Coordinate{
				x: c.Start.x + i*xShift,
				y: c.Start.y,
			})
		}
		break
	case yDiff != 0:
		// vertical
		for i := 0; i <= int(math.Abs(float64(yDiff))); i++ {
			coordinates = append(coordinates, Coordinate{
				x: c.Start.x,
				y: c.Start.y + i*yShift,
			})
		}
		break
	}
	return coordinates
}
