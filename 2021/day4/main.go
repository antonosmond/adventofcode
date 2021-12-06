package main

import (
	"adventofcode/pkg/input"
	"fmt"
	"strconv"
	"strings"
)

type Bingo struct {
	DrawNumbers []int
	Boards      []Board
}

type Board struct {
	Rows []Row
}

type Row struct {
	Cells []*Cell
}

type Cell struct {
	Number int
	Column int
	Marked bool
}

func main() {
	in := input.Load()
	bingo := toBingo(in)
	for _, i := range bingo.DrawNumbers {
		for b := 0; b < len(bingo.Boards); b++ {
			board := bingo.Boards[b]
			if board.Mark(i) {
				score := board.Score(i)
				fmt.Printf("BINGO! Score: %d\n", score)
				bingo.Boards = append(bingo.Boards[:b], bingo.Boards[b+1:]...)
				b--
				// os.Exit(0)
			}
		}
	}
}

func (b *Board) AddRow(nums []int) {
	r := Row{
		Cells: []*Cell{},
	}
	for c, n := range nums {
		r.Cells = append(r.Cells, &Cell{
			Number: n,
			Column: c,
			Marked: false,
		})
	}
	b.Rows = append(b.Rows, r)
}

func toBingo(in []string) Bingo {
	// The first row is the draw numbers
	bingo := Bingo{
		DrawNumbers: stringsToInts(strings.ReplaceAll(in[0], ",", " ")),
		Boards:      []Board{},
	}
	boardIndex := 0
	for _, line := range in[1:] {
		// each time we get an empty line, start a new board and update the index
		if line == "" {
			bingo.Boards = append(bingo.Boards, Board{})
			boardIndex = len(bingo.Boards) - 1
			continue
		}
		bingo.Boards[boardIndex].AddRow(stringsToInts(line))
	}
	return bingo
}

func stringsToInts(s string) []int {
	row := strings.Split(s, " ")
	ints := []int{}
	for _, n := range row {
		if n == "" {
			continue
		}
		i, err := strconv.ParseInt(n, 10, 32)
		if err != nil {
			panic(err)
		}
		ints = append(ints, int(i))
	}
	return ints
}

func (b *Board) Mark(i int) bool {
	for _, r := range b.Rows {
		for _, c := range r.Cells {
			if c.Number == i {
				c.Marked = true
				return b.HasWin()
			}
		}
	}
	return false
}

func (b *Board) HasWin() bool {
	markedCol := make([]int, 5)
	for _, r := range b.Rows {
		marked := 0
		for i, c := range r.Cells {
			if c.Marked {
				marked++
				markedCol[i]++
			}
		}
		if marked == 5 {
			return true
		}
	}
	for _, marked := range markedCol {
		if marked == 5 {
			return true
		}
	}
	return false
}

func (b *Board) Score(winningNumber int) int {
	sum := 0
	for _, r := range b.Rows {
		for _, c := range r.Cells {
			if !c.Marked {
				sum += c.Number
			}
		}
	}
	return sum * winningNumber
}
