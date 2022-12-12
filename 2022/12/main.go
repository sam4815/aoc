package main

import (
	"bufio"
	"log"
	"math"
	"os"
	"strings"
	"time"
)

func testCell(current_cell []int, proposed_cell []int, forest *[][]string, steps_map *[][]int) bool {
	current_letter := (*forest)[current_cell[0]][current_cell[1]]
	current_cost := (*steps_map)[current_cell[0]][current_cell[1]]
	proposed_letter := (*forest)[proposed_cell[0]][proposed_cell[1]]
	proposed_cost := (*steps_map)[proposed_cell[0]][proposed_cell[1]]

	is_movable := int(current_letter[0])-int(proposed_letter[0]) <= 1
	is_worth_moving := proposed_cost > current_cost+1

	return is_movable && is_worth_moving
}

func markAdjacentCells(cell []int, forest *[][]string, steps_map *[][]int) {
	adjacent_cells := make([][]int, 0)
	curr_steps := (*steps_map)[cell[0]][cell[1]]
	possible_cells := [][]int{
		{cell[0] - 1, cell[1]},
		{cell[0] + 1, cell[1]},
		{cell[0], cell[1] - 1},
		{cell[0], cell[1] + 1},
	}

	for _, test_cell := range possible_cells {
		if test_cell[0] == -1 || test_cell[0] == len(*forest) || test_cell[1] == -1 || test_cell[1] == len((*forest)[0]) {
			continue
		}

		if testCell(cell, test_cell, forest, steps_map) {
			adjacent_cells = append(adjacent_cells, test_cell)
			(*steps_map)[test_cell[0]][test_cell[1]] = curr_steps + 1
		}
	}

	for _, cell := range adjacent_cells {
		markAdjacentCells(cell, forest, steps_map)
	}
}

func main() {
	start := time.Now()

	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)

	forest := make([][]string, 0)
	steps_map := make([][]int, 0)
	start_pos, end_pos := []int{0, 0}, []int{0, 0}
	row_num := 0

	for scanner.Scan() {
		row := make([]string, 0)
		letters := strings.Split(scanner.Text(), "")
		for col_num, letter := range letters {
			if letter == "S" {
				start_pos[0], start_pos[1] = row_num, col_num
			}
			if letter == "E" {
				end_pos[0], end_pos[1] = row_num, col_num
			}
			row = append(row, letter)
		}

		forest = append(forest, row)

		steps_row := make([]int, len(row))
		for i := range steps_row {
			steps_row[i] = math.MaxInt
		}
		steps_map = append(steps_map, steps_row)

		row_num++
	}

	forest[start_pos[0]][start_pos[1]] = "a"
	forest[end_pos[0]][end_pos[1]] = "z"
	steps_map[end_pos[0]][end_pos[1]] = 0

	markAdjacentCells(end_pos, &forest, &steps_map)

	time_elapsed := time.Since(start)

	fewest_steps_from_start := steps_map[start_pos[0]][start_pos[1]]
	fewest_steps_from_a := math.MaxInt

	for i, row := range forest {
		for j, cell := range row {
			if cell == "a" && steps_map[i][j] < fewest_steps_from_a {
				fewest_steps_from_a = steps_map[i][j]
			}
		}
	}

	log.Printf(`
The fewest steps required to reach z from the start is %d.
The fewest steps required to reach z from any a is %d.
Solution generated in %s.`,
		fewest_steps_from_start,
		fewest_steps_from_a,
		time_elapsed,
	)
}
