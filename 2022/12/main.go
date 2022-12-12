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

	is_movable := int(proposed_letter[0])-int(current_letter[0]) <= 1
	is_worth_moving := proposed_cost > current_cost+1

	return is_movable && is_worth_moving
}

func markAdjacentCells(cell []int, forest *[][]string, steps_map *[][]int) {
	adjacent_cells := make([][]int, 0)
	curr_steps := (*steps_map)[cell[0]][cell[1]]
	// Up
	if cell[0] > 0 {
		up_cell := []int{cell[0] - 1, cell[1]}

		if testCell(cell, up_cell, forest, steps_map) {
			adjacent_cells = append(adjacent_cells, up_cell)
			(*steps_map)[up_cell[0]][up_cell[1]] = curr_steps + 1
		}
	}
	// Down
	if cell[0] < len(*forest)-1 {
		down_cell := []int{cell[0] + 1, cell[1]}

		if testCell(cell, down_cell, forest, steps_map) {
			adjacent_cells = append(adjacent_cells, down_cell)
			(*steps_map)[down_cell[0]][down_cell[1]] = curr_steps + 1
		}
	}
	// Left
	if cell[1] > 0 {
		left_cell := []int{cell[0], cell[1] - 1}

		if testCell(cell, left_cell, forest, steps_map) {
			adjacent_cells = append(adjacent_cells, left_cell)
			(*steps_map)[left_cell[0]][left_cell[1]] = curr_steps + 1
		}
	}
	// Right
	if cell[1] < len((*forest)[0])-1 {
		right_cell := []int{cell[0], cell[1] + 1}

		if testCell(cell, right_cell, forest, steps_map) {
			adjacent_cells = append(adjacent_cells, right_cell)
			(*steps_map)[right_cell[0]][right_cell[1]] = curr_steps + 1
		}
	}

	for _, cell := range adjacent_cells {
		markAdjacentCells(cell, forest, steps_map)
	}
}

func findFewestSteps(start []int, end []int, forest [][]string, steps_map [][]int) int {
	steps_map[start[0]][start[1]] = 0
	markAdjacentCells(start, &forest, &steps_map)
	return steps_map[end[0]][end[1]]
}

func clone2DSlice[T any](src [][]T) [][]T {
	dst := make([][]T, len(src))
	for i := range dst {
		dst_row := make([]T, len(src[i]))
		for j := range dst_row {
			dst_row[j] = src[i][j]
		}
		dst[i] = dst_row
	}

	return dst
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

	forest[end_pos[0]][end_pos[1]] = "z"

	fewest_steps_from_start := math.MaxInt
	fewest_steps_from_a := math.MaxInt

	for i := range forest {
		for j := range forest[i] {
			if forest[i][j] == "S" {
				forest_copy := clone2DSlice(forest)
				steps_map_copy := clone2DSlice(steps_map)

				forest_copy[i][j] = "a"
				fewest_steps_from_start = findFewestSteps(start_pos, end_pos, forest_copy, steps_map_copy)
			} else if forest[i][j] == "a" {
				fewest_steps := findFewestSteps([]int{i, j}, end_pos, clone2DSlice(forest), clone2DSlice(steps_map))
				if fewest_steps < fewest_steps_from_a {
					fewest_steps_from_a = fewest_steps
				}
			}
		}
	}

	// log.Print(start_pos, end_pos)
	// for i := range forest {
	// 	log.Printf("%v", steps_map[i])
	// }

	time_elapsed := time.Since(start)
	log.Printf(`
The fewest steps required to reach z from the start is %d.
The fewest steps required to reach z from any a is %d.
Solution generated in %s.`,
		fewest_steps_from_start,
		fewest_steps_from_a,
		time_elapsed,
	)
}
