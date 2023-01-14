package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type Cavern [][]int
type PathQueue []Path

type Path struct {
	risk_level    int
	curr_position []int
}

func (cavern Cavern) Get(position []int) int {
	return cavern[position[0]][position[1]]
}

func (cavern *Cavern) Set(position []int, val int) {
	(*cavern)[position[0]][position[1]] = val
}

func (path Path) Clone() Path {
	clone := Path{risk_level: path.risk_level, curr_position: make([]int, 0)}
	copy(clone.curr_position, path.curr_position)
	return clone
}

func (cavern Cavern) Expand() Cavern {
	full_cavern := make(Cavern, len(cavern)*5)
	partial_cavern_len := len(cavern)

	for i := 0; i < len(full_cavern); i++ {
		full_cavern[i] = make([]int, len(full_cavern))

		for j := 0; j < len(full_cavern); j++ {
			to_add := (i / partial_cavern_len) + (j / partial_cavern_len)
			value := (cavern.Get([]int{i % partial_cavern_len, j % partial_cavern_len}) + to_add)
			value = ((value - 1) % 9) + 1

			full_cavern.Set([]int{i, j}, value)
		}
	}

	return full_cavern
}

func (cavern Cavern) GetDumbRiskLevels() Cavern {
	dumb_risk_levels := make(Cavern, len(cavern))

	for i := 0; i < len(cavern); i++ {
		dumb_risk_levels[i] = make([]int, len(cavern))

		for j := 0; j < len(cavern); j++ {
			curr_coords := []int{i, j}

			if i == 0 && j == 0 {
				dumb_risk_levels.Set(curr_coords, cavern.Get(curr_coords))
			} else if j == 0 {
				dumb_risk_levels.Set(curr_coords, cavern.Get(curr_coords)+dumb_risk_levels.Get([]int{i - 1, j}))
			} else {
				dumb_risk_levels.Set(curr_coords, cavern.Get(curr_coords)+dumb_risk_levels.Get([]int{i, j - 1}))
			}
		}
	}

	return dumb_risk_levels
}

func (path Path) GetPossibleMoves(cavern Cavern) [][]int {
	possible_moves := make([][]int, 0)
	if path.curr_position[0] < len(cavern)-1 {
		possible_moves = append(possible_moves, []int{path.curr_position[0] + 1, path.curr_position[1]})
	}
	if path.curr_position[1] < len(cavern)-1 {
		possible_moves = append(possible_moves, []int{path.curr_position[0], path.curr_position[1] + 1})
	}
	if path.curr_position[0] > 0 {
		possible_moves = append(possible_moves, []int{path.curr_position[0] - 1, path.curr_position[1]})
	}
	if path.curr_position[1] > 0 {
		possible_moves = append(possible_moves, []int{path.curr_position[0], path.curr_position[1] - 1})
	}
	return possible_moves
}

// Insert the element so that the queue is ordered by lowest cost first
func (queue *PathQueue) Push(path Path) {
	index := len(*queue)
	for i := 0; i < len(*queue); i++ {
		if path.risk_level <= (*queue)[i].risk_level {
			index = i
			break
		}
	}

	if index == len(*queue) {
		*queue = append(*queue, path)
	} else {
		*queue = append((*queue)[:index+1], (*queue)[index:]...)
		(*queue)[index] = path
	}
}

func (cavern Cavern) FindLowestRiskPath() int {
	smallest_levels := cavern.GetDumbRiskLevels()
	starting_path := Path{curr_position: []int{0, 0}, risk_level: 0}
	path_queue, curr_path := PathQueue{starting_path}, starting_path

	for len(path_queue) > 0 {
		curr_path, path_queue = path_queue[0], path_queue[1:]

		if smallest_levels.Get(curr_path.curr_position) <= curr_path.risk_level {
			continue
		} else {
			smallest_levels.Set(curr_path.curr_position, curr_path.risk_level)
		}

		for _, move := range curr_path.GetPossibleMoves(cavern) {
			new_path := curr_path.Clone()
			new_path.curr_position = move
			new_path.risk_level += cavern.Get(move)
			path_queue.Push(new_path)
		}
	}

	return smallest_levels[len(cavern)-1][len(cavern)-1]
}

func main() {
	start := time.Now()
	f, _ := os.Open("input.txt")
	defer f.Close()
	scanner := bufio.NewScanner(f)

	partial_cavern := make(Cavern, 0)
	for scanner.Scan() {
		row := make([]int, 0)
		for _, risk_string := range strings.Split(scanner.Text(), "") {
			risk_level, _ := strconv.Atoi(risk_string)
			row = append(row, risk_level)
		}
		partial_cavern = append(partial_cavern, row)
	}

	full_cavern := partial_cavern.Expand()

	lowest_total_risk_partial_map := partial_cavern.FindLowestRiskPath()
	lowest_total_risk_full_map := full_cavern.FindLowestRiskPath()

	time_elapsed := time.Since(start)

	log.Printf(`
The lowest possible total risk for the partial cavern is %d,
The lowest possible total risk for the full cavern is %d,
Solution generated in %s.`,
		lowest_total_risk_partial_map,
		lowest_total_risk_full_map,
		time_elapsed,
	)
}
