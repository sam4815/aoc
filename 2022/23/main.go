package main

import (
	"bufio"
	"log"
	"math"
	"os"
	"strings"
	"time"
)

type Elf struct {
	curr_position []int
	next_position []int
}

type Scan struct {
	elves []Elf
	grid  [][]string
	round int
	size  int
}

func (scan *Scan) InitializeGrid() {
	grid := make([][]string, scan.size)
	for i := range grid {
		grid[i] = make([]string, scan.size)
		for j := range grid[i] {
			grid[i][j] = "."
		}
	}

	scan.grid = grid
}

func (scan Scan) GetBounds() (int, int, int, int) {
	x_min, y_min := math.MaxInt, math.MaxInt
	x_max, y_max := math.MinInt, math.MinInt

	for y := 0; y < len(scan.grid); y++ {
		for x := 0; x < len(scan.grid); x++ {
			if scan.grid[y][x] != "#" {
				continue
			}

			if y < y_min {
				y_min = y
			}
			if y > y_max {
				y_max = y
			}
			if x < x_min {
				x_min = x
			}
			if x > x_max {
				x_max = x
			}
		}
	}

	return x_min, x_max, y_min, y_max
}

func (scan Scan) CountEmptyCells() int {
	x_min, x_max, y_min, y_max := scan.GetBounds()
	num_empty_cells := (x_max - x_min + 1) * (y_max - y_min + 1)

	for y := y_min; y <= y_max; y++ {
		for x := x_min; x <= x_max; x++ {
			if scan.grid[y][x] == "#" {
				num_empty_cells--
			}
		}
	}

	return num_empty_cells
}

func (scan Scan) GetAdjacentPositions(position []int) [][]int {
	return [][]int{
		{position[0] - 1, position[1] - 1},
		{position[0] - 1, position[1]},
		{position[0] - 1, position[1] + 1},
		{position[0], position[1] - 1},
		{position[0], position[1] + 1},
		{position[0] + 1, position[1] - 1},
		{position[0] + 1, position[1]},
		{position[0] + 1, position[1] + 1},
	}
}

func (scan Scan) HasAdjacentElf(elf Elf) bool {
	for _, position := range scan.GetAdjacentPositions(elf.curr_position) {
		if scan.grid[position[0]][position[1]] == "#" {
			return true
		}
	}
	return false
}

func (scan Scan) HasNorthernElf(elf Elf) bool {
	for _, position := range scan.GetAdjacentPositions(elf.curr_position)[:3] {
		if scan.grid[position[0]][position[1]] == "#" {
			return true
		}
	}
	return false
}

func (scan Scan) HasSouthernElf(elf Elf) bool {
	for _, position := range scan.GetAdjacentPositions(elf.curr_position)[5:] {
		if scan.grid[position[0]][position[1]] == "#" {
			return true
		}
	}
	return false
}

func (scan Scan) HasWesternElf(elf Elf) bool {
	for _, position := range scan.GetAdjacentPositions(elf.curr_position) {
		if scan.grid[position[0]][position[1]] == "#" && position[1] == elf.curr_position[1]-1 {
			return true
		}
	}
	return false
}

func (scan Scan) HasEasternElf(elf Elf) bool {
	for _, position := range scan.GetAdjacentPositions(elf.curr_position) {
		if scan.grid[position[0]][position[1]] == "#" && position[1] == elf.curr_position[1]+1 {
			return true
		}
	}
	return false
}

func (scan *Scan) MarkNextPosition(elf Elf) {
	if elf.next_position == nil {
		return
	}

	if scan.grid[elf.next_position[0]][elf.next_position[1]] == "O" {
		scan.grid[elf.next_position[0]][elf.next_position[1]] = "X"
	} else {
		scan.grid[elf.next_position[0]][elf.next_position[1]] = "O"
	}
}

func (scan Scan) GetNextPositionTest(i int, elf Elf) (func(elf Elf) bool, []int) {
	if i == 0 {
		return scan.HasNorthernElf, []int{elf.curr_position[0] - 1, elf.curr_position[1]}
	} else if i == 1 {
		return scan.HasSouthernElf, []int{elf.curr_position[0] + 1, elf.curr_position[1]}
	} else if i == 2 {
		return scan.HasWesternElf, []int{elf.curr_position[0], elf.curr_position[1] - 1}
	} else {
		return scan.HasEasternElf, []int{elf.curr_position[0], elf.curr_position[1] + 1}
	}
}

func (scan *Scan) SetNextPositions(round int) bool {
	elf_needs_to_move := false

	for i, elf := range scan.elves {
		if !scan.HasAdjacentElf(elf) {
			continue
		}
		elf_needs_to_move = true

		for j := 0; j < 4; j++ {
			test, next_position := scan.GetNextPositionTest((round-1+j)%4, elf)
			if !test(elf) {
				scan.elves[i].next_position = next_position
				break
			}
		}

		scan.MarkNextPosition(scan.elves[i])
	}

	return elf_needs_to_move
}

func (scan *Scan) ClearProposedPositions() {
	for i := 0; i < scan.size; i++ {
		for j := 0; j < scan.size; j++ {
			if scan.grid[i][j] == "X" {
				scan.grid[i][j] = "."
			}
		}
	}
}

func (scan *Scan) MoveToNextPosition(elf_index int, elf Elf) {
	scan.grid[elf.curr_position[0]][elf.curr_position[1]] = "."
	scan.grid[elf.next_position[0]][elf.next_position[1]] = "#"

	scan.elves[elf_index].curr_position = scan.elves[elf_index].next_position
	scan.elves[elf_index].next_position = nil
}

func (scan *Scan) MoveToNextPositions() {
	for i, elf := range scan.elves {
		if elf.next_position == nil {
			continue
		}

		if scan.grid[elf.next_position[0]][elf.next_position[1]] == "O" {
			scan.MoveToNextPosition(i, elf)
		} else if scan.grid[elf.next_position[0]][elf.next_position[1]] == "X" {
			scan.elves[i].next_position = nil
		} else {
			log.Fatal("Wut?")
		}
	}

	scan.ClearProposedPositions()
}

func main() {
	start := time.Now()

	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)

	scan := Scan{size: 200, round: 1}
	scan.InitializeGrid()

	row_num := scan.size / 3

	for scanner.Scan() {
		row := strings.Split(scanner.Text(), "")
		for cell := range row {
			if row[cell] == "#" {
				scan.grid[row_num][cell+scan.size/3] = "#"
				scan.elves = append(scan.elves, Elf{curr_position: []int{row_num, cell + scan.size/3}})
			}
		}
		row_num += 1
	}

	round := 1
	for round <= 10 {
		scan.SetNextPositions(round)
		scan.MoveToNextPositions()
		round += 1
	}

	num_empty_tiles := scan.CountEmptyCells()

	for scan.SetNextPositions(round) {
		scan.MoveToNextPositions()
		round += 1
	}

	time_elapsed := time.Since(start)

	log.Printf(`
The number of empty tiles after 10 rounds is %d.
The first round where no elf moves is %d.
Solution generated in %s.`,
		num_empty_tiles,
		round,
		time_elapsed,
	)
}
