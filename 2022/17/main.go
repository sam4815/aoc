package main

import (
	"log"
	"os"
	"strings"
	"time"
)

type Rock struct {
	width     int
	height    int
	formation [][]string
}

type Chamber struct {
	grid               [][]string
	max_height         int
	num_discarded_rows int
}

func getRock(i int) Rock {
	rocks := []Rock{
		{
			width:     4,
			height:    1,
			formation: [][]string{{"@", "@", "@", "@"}},
		},
		{
			width:     3,
			height:    3,
			formation: [][]string{{".", "@", "."}, {"@", "@", "@"}, {".", "@", "."}},
		},
		{
			width:     3,
			height:    3,
			formation: [][]string{{".", ".", "@"}, {".", ".", "@"}, {"@", "@", "@"}},
		},
		{
			width:     1,
			height:    4,
			formation: [][]string{{"@"}, {"@"}, {"@"}, {"@"}},
		},
		{
			width:     2,
			height:    2,
			formation: [][]string{{"@", "@"}, {"@", "@"}},
		},
	}

	return rocks[i]
}

func (chamber *Chamber) HighestRockRow() int {
	highest_row := 0

	for i := 0; i < len(chamber.grid); i++ {
		for j := 0; j < len(chamber.grid[i]); j++ {
			if chamber.grid[i][j] == "#" {
				return i
			}
		}
	}

	return highest_row
}

func (chamber *Chamber) Height() int {
	return len(chamber.grid) - chamber.HighestRockRow() + chamber.num_discarded_rows
}

func (chamber *Chamber) AddEmptyRow() {
	empty_row := make([]string, 7)
	for i := 0; i < len(empty_row); i++ {
		empty_row[i] = "."
	}

	chamber.grid = append([][]string{empty_row}, chamber.grid...)
}

func (chamber *Chamber) ShiftRowsDown() {
	for i := chamber.max_height - 1; i > 0; i-- {
		for j := 0; j < len(chamber.grid[i]); j++ {
			chamber.grid[i][j] = chamber.grid[i-1][j]
		}
	}

	for j := 0; j < len(chamber.grid[0]); j++ {
		chamber.grid[0][j] = "."
	}

	chamber.num_discarded_rows++
}

func (chamber *Chamber) AddEmptyRows(num_rows int) {
	for i := 0; i < num_rows; i++ {
		if len(chamber.grid) < chamber.max_height {
			chamber.AddEmptyRow()
		} else {
			chamber.ShiftRowsDown()
		}
	}
}

func (chamber *Chamber) RemoveEmptyRowFromTop() {
	chamber.grid = chamber.grid[1:]
}

func (chamber *Chamber) RemoveEmptyRows(num_rows int) {
	for i := 0; i < num_rows; i++ {
		chamber.RemoveEmptyRowFromTop()
	}
}

func (chamber *Chamber) DiscardBottomRows() {
	if len(chamber.grid) > chamber.max_height {
		chamber.num_discarded_rows += len(chamber.grid) - chamber.max_height
		chamber.grid = chamber.grid[:chamber.max_height]
	}
}

func (chamber *Chamber) AddRock(rock Rock) {
	highest_row := chamber.HighestRockRow()

	required_rows_from_top := 3 + rock.height
	required_addition := required_rows_from_top - highest_row
	if required_addition < 0 {
		chamber.RemoveEmptyRows(-required_addition)
	} else if required_addition > 0 {
		chamber.AddEmptyRows(required_addition)
	}

	// chamber.DiscardBottomRows()

	for i := 0; i < rock.height; i++ {
		for j := 0; j < rock.width; j++ {
			chamber.grid[i][j+2] = rock.formation[i][j]
		}
	}
}

func (chamber *Chamber) HasActiveRock() bool {
	has_active_rock := false

	for i := len(chamber.grid) - 1; i >= 0; i-- {
		for j := 0; j < len(chamber.grid[i]); j++ {
			if chamber.grid[i][j] == "@" {
				has_active_rock = true
			}
		}
	}

	return has_active_rock
}

func getDirectionFunction(direction string) func(int) int {
	if direction == ">" {
		return func(i int) int { return i + 1 }
	}
	return func(i int) int { return i - 1 }
}

func (chamber *Chamber) PushRock(direction string) {
	direction_function := getDirectionFunction(direction)
	can_move := true

	for i := len(chamber.grid) - 1; i >= 0; i-- {
		for j := 0; j < len(chamber.grid[i]); j++ {
			if chamber.grid[i][j] == "@" {
				target_cell := direction_function(j)

				if target_cell < 0 || target_cell >= len(chamber.grid[i]) || chamber.grid[i][target_cell] == "#" {
					can_move = false
				}
			}
		}
	}

	if !can_move {
		return
	}

	for i := len(chamber.grid) - 1; i >= 0; i-- {
		if direction_function(0) == 1 {
			for j := len(chamber.grid[i]) - 1; j >= 0; j-- {
				if chamber.grid[i][j] == "@" {
					target_cell := direction_function(j)
					chamber.grid[i][target_cell], chamber.grid[i][j] = "@", "."
				}
			}
		} else {
			for j := 0; j < len(chamber.grid[i]); j++ {
				if chamber.grid[i][j] == "@" {
					target_cell := direction_function(j)
					chamber.grid[i][target_cell], chamber.grid[i][j] = "@", "."
				}
			}
		}
	}
}

func (chamber *Chamber) BringRockToRest() {
	for i := len(chamber.grid) - 1; i >= 0; i-- {
		for j := 0; j < len(chamber.grid[i]); j++ {
			if chamber.grid[i][j] == "@" {
				chamber.grid[i][j] = "#"
			}
		}
	}
}

func (chamber *Chamber) Fall() {
	can_fall := true

	for i := len(chamber.grid) - 1; i >= 0; i-- {
		for j := 0; j < len(chamber.grid[i]); j++ {
			if chamber.grid[i][j] == "@" {
				if i == len(chamber.grid)-1 || chamber.grid[i+1][j] == "#" {
					can_fall = false
				}
			}
		}
	}

	if !can_fall {
		chamber.BringRockToRest()
	}

	for i := len(chamber.grid) - 1; i >= 0; i-- {
		for j := 0; j < len(chamber.grid[i]); j++ {
			if chamber.grid[i][j] == "@" {
				chamber.grid[i+1][j], chamber.grid[i][j] = "@", "."
			}
		}
	}
}

func (chamber *Chamber) Print() {
	for i := 0; i < len(chamber.grid); i++ {
		log.Print(chamber.grid[i])
	}
	log.Print("+- - - - - - -+")
}

func main() {
	start := time.Now()

	f, err := os.ReadFile("input.txt")
	if err != nil {
		log.Print(err)
	}

	jets := strings.Split(string(f), "")
	jets_length, cycle_length := len(jets), len(jets)*5
	jet_index, rock_index := 0, 0
	height_after_first_cycle, rocks_after_first_cycle := 0, 0

	chamber := Chamber{max_height: 50}

	small_rocks_height := 0
	target_rock := 1000000000000

	for i := 0; i < target_rock; i++ {
		chamber.AddRock(getRock(rock_index % 5))

		for chamber.HasActiveRock() {
			chamber.PushRock(jets[jet_index%jets_length])
			jet_index++
			chamber.Fall()
			// Check whether we've finished the cycle
			if jet_index%cycle_length == 0 {
				if height_after_first_cycle == 0 {
					height_after_first_cycle = chamber.Height()
					rocks_after_first_cycle = i
					continue
				}

				rocks_processed_in_cycle := i - rocks_after_first_cycle
				height_gained_in_cycle := chamber.Height() - height_after_first_cycle

				num_cycles_required := (target_rock - i) / rocks_processed_in_cycle

				chamber.num_discarded_rows = (num_cycles_required * height_gained_in_cycle) + chamber.num_discarded_rows
				i += (num_cycles_required * rocks_processed_in_cycle)
			}
		}

		if i == 2022 {
			small_rocks_height = chamber.Height()
		}

		rock_index++
	}

	big_rocks_height := chamber.Height()

	time_elapsed := time.Since(start)

	log.Printf(`
The tower of rocks is %d units tall after 2022 rocks have fallen.
The tower of rocks is %d units tall after 1000000000000 rocks have fallen.
Solution generated in %s.`,
		small_rocks_height,
		big_rocks_height,
		time_elapsed,
	)
}
