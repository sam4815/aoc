package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type Blizzard struct {
	direction     string
	curr_position []int
}

type Valley struct {
	grid      [][]string
	blizzards []Blizzard
	width     int
	height    int
}

type Path struct {
	curr_position     []int
	next_position     []int
	minute            int
	has_reached_end   bool
	has_reached_start bool
}

type Moment struct {
	y_pos         int
	x_pos         int
	cycle         int
	reached_end   bool
	reached_start bool
}

func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func LCM(a, b int) int {
	return a * b / GCD(a, b)
}

func (valley Valley) Clone() Valley {
	copy := valley
	copy.grid = make([][]string, copy.height)
	copy.blizzards = make([]Blizzard, len(valley.blizzards))

	for i, row := range valley.grid {
		copy.grid[i] = make([]string, copy.width)
		for j, cell := range row {
			copy.grid[i][j] = cell
		}
	}

	for i, blizzard := range valley.blizzards {
		copy.blizzards[i] = Blizzard{
			curr_position: make([]int, 2),
			direction:     blizzard.direction,
		}
		copy.blizzards[i].curr_position[0], copy.blizzards[i].curr_position[1] = valley.blizzards[i].curr_position[0], valley.blizzards[i].curr_position[1]
	}

	return copy
}

func (path Path) Clone() Path {
	copy := Path{minute: path.minute, curr_position: make([]int, 2), next_position: make([]int, 2), has_reached_end: path.has_reached_end, has_reached_start: path.has_reached_start}
	copy.curr_position[0], copy.curr_position[1] = path.curr_position[0], path.curr_position[1]

	return copy
}

func (path Path) GetAdjacentPlayerPositions(valleys []Valley) [][]int {
	positions := make([][]int, 0)
	possible_positions := [][]int{
		{path.curr_position[0], path.curr_position[1] + 1},
		{path.curr_position[0] + 1, path.curr_position[1]},
		{path.curr_position[0] - 1, path.curr_position[1]},
		{path.curr_position[0], path.curr_position[1] - 1},
	}

	if path.has_reached_end && !path.has_reached_start {
		possible_positions = [][]int{
			{path.curr_position[0] - 1, path.curr_position[1]},
			{path.curr_position[0], path.curr_position[1] - 1},
			{path.curr_position[0], path.curr_position[1] + 1},
			{path.curr_position[0] + 1, path.curr_position[1]},
		}
	}

	for _, position := range possible_positions {
		if position[0] >= 0 && position[0] < valleys[0].height && valleys[0].GetVal(position) != "#" {
			positions = append(positions, position)
		}
	}

	return positions
}

func (path Path) WillPositionHaveBlizzard(position []int, valleys []Valley) bool {
	if position[0] == 0 || position[0] == valleys[0].height-1 {
		return false
	}

	next_valley := valleys[(path.minute+1)%len(valleys)]
	return next_valley.GetVal(position) != "."
}

func (valley Valley) GetNextBlizzardPosition(blizzard Blizzard) []int {
	next_position := []int{blizzard.curr_position[0], blizzard.curr_position[1]}

	switch blizzard.direction {
	case "^":
		next_position[0] -= 1
	case "v":
		next_position[0] += 1
	case "<":
		next_position[1] -= 1
	case ">":
		next_position[1] += 1
	}

	return valley.GetWrappedPosition(next_position)
}

func (valley *Valley) RemoveBlizzardFromPosition(blizzard Blizzard) {
	if len(valley.GetVal(blizzard.curr_position)) == 1 {
		valley.SetVal(blizzard.curr_position, ".")
	} else {
		valley.SetVal(blizzard.curr_position, strings.Replace(valley.GetVal(blizzard.curr_position), blizzard.direction, "", 1))
	}
}

func (valley *Valley) AddBlizzardToPosition(blizzard Blizzard) {
	if valley.GetVal(blizzard.curr_position) == "." {
		valley.SetVal(blizzard.curr_position, blizzard.direction)
	} else {
		valley.SetVal(blizzard.curr_position, valley.GetVal(blizzard.curr_position)+blizzard.direction)
	}
}

func (valley *Valley) GetWrappedPosition(position []int) []int {
	wrapped_position := []int{position[0], position[1]}

	if wrapped_position[0] == 0 {
		wrapped_position[0] = valley.height - 2
	} else if wrapped_position[0] == valley.height-1 {
		wrapped_position[0] = 1
	}

	if wrapped_position[1] == 0 {
		wrapped_position[1] = valley.width - 2
	} else if wrapped_position[1] == valley.width-1 {
		wrapped_position[1] = 1
	}

	return wrapped_position
}

func (valley Valley) ForEach(lambda func(coordinates []int)) {
	for y := 1; y < valley.height-1; y++ {
		for x := 1; x < valley.width-1; x++ {
			lambda([]int{y, x})
		}
	}
}

func (valley Valley) GetVal(coordinates []int) string {
	return valley.grid[coordinates[0]][coordinates[1]]
}

func (valley *Valley) SetVal(coordinates []int, value string) {
	valley.grid[coordinates[0]][coordinates[1]] = value
}

func (valley *Valley) Print() {
	for i, row := range valley.grid {
		fmt_row := make([]string, len(row))
		for j := range row {
			if len(valley.grid[i][j]) >= 2 {
				fmt_row[j] = strconv.Itoa(len(valley.grid[i][j]))
			} else {
				fmt_row[j] = valley.grid[i][j]
			}
		}
		log.Print(fmt_row)
	}
}

func HasReachedEnd(path Path, valley Valley) bool {
	return path.curr_position[0] == valley.height-1
}

func HasReachedStart(path Path, valley Valley) bool {
	return path.curr_position[0] == 0 && path.has_reached_end
}

func HasReachedEndAgain(path Path, valley Valley) bool {
	return path.curr_position[0] == valley.height-1 && path.has_reached_start
}

func (path *Path) Tick(valley Valley) {
	if path.next_position[1] != 0 {
		path.curr_position[0], path.curr_position[1] = path.next_position[0], path.next_position[1]
		path.next_position[1] = 0
	}

	if HasReachedEnd(*path, valley) {
		path.has_reached_end = true
	}

	if HasReachedStart(*path, valley) {
		path.has_reached_start = true
	}

	path.minute += 1
}

func (valley *Valley) Tick() {
	for i := range valley.blizzards {
		next_position := valley.GetNextBlizzardPosition(valley.blizzards[i])
		valley.RemoveBlizzardFromPosition(valley.blizzards[i])
		valley.blizzards[i].curr_position[0], valley.blizzards[i].curr_position[1] = next_position[0], next_position[1]
		valley.AddBlizzardToPosition(valley.blizzards[i])
	}
}

func (p Path) GetPossiblePaths(valleys []Valley) []Path {
	original_path := p.Clone()
	paths := make([]Path, 0)
	num_minutes_waited := 0

	for {
		for _, position := range original_path.GetAdjacentPlayerPositions(valleys) {
			if original_path.WillPositionHaveBlizzard(position, valleys) {
				continue
			}

			branching_path := original_path.Clone()
			branching_path.next_position[0], branching_path.next_position[1] = position[0], position[1]
			branching_path.Tick(valleys[0])

			paths = append([]Path{branching_path}, paths...)
		}

		if original_path.WillPositionHaveBlizzard(original_path.curr_position, valleys) || num_minutes_waited > valleys[0].width {
			break
		}

		num_minutes_waited += 1
		original_path.Tick(valleys[0])
	}

	return paths
}

func (path Path) FindShortestDistance(valleys []Valley, goal func(path Path, valley Valley) bool) int {
	path_queue, curr_path := []Path{path}, path

	shortest_path := 0
	visited_map := make(map[Moment]int)

	for len(path_queue) > 0 {
		curr_path, path_queue = path_queue[0], path_queue[1:]

		if goal(curr_path, valleys[0]) {
			if curr_path.minute < shortest_path || shortest_path == 0 {
				shortest_path = curr_path.minute
			}
			log.Print("SHORTEST IS: ", shortest_path)
			continue
		}

		moment := Moment{
			y_pos:         curr_path.curr_position[0],
			x_pos:         curr_path.curr_position[1],
			cycle:         curr_path.minute % len(valleys),
			reached_end:   curr_path.has_reached_end,
			reached_start: curr_path.has_reached_start,
		}

		if visited_map[moment] == 0 || visited_map[moment] > curr_path.minute {
			visited_map[moment] = curr_path.minute
		} else {
			continue
		}

		if curr_path.minute > 1500 {
			continue
		}

		if shortest_path != 0 && curr_path.minute >= shortest_path {
			continue
		}

		possible_paths := curr_path.GetPossiblePaths(valleys)
		path_queue = append(possible_paths, path_queue...)
	}

	return shortest_path
}

func main() {
	start := time.Now()

	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)

	valley := Valley{grid: make([][]string, 0), blizzards: make([]Blizzard, 0)}
	path := Path{curr_position: []int{0, 1}, next_position: []int{0, 0}}

	for scanner.Scan() {
		valley.grid = append(valley.grid, strings.Split(scanner.Text(), ""))
	}

	valley.height = len(valley.grid)
	valley.width = len(valley.grid[0])

	valley.ForEach(func(coordinates []int) {
		if valley.GetVal(coordinates) != "." {
			blizzard := Blizzard{direction: valley.GetVal(coordinates), curr_position: coordinates}
			valley.blizzards = append(valley.blizzards, blizzard)
		}
	})

	valley_permutations := LCM(valley.height-2, valley.width-2)

	valleys := make([]Valley, valley_permutations)
	valleys[0] = valley

	for i := 1; i < valley_permutations; i++ {
		clone := valleys[i-1].Clone()
		clone.Tick()
		valleys[i] = clone
	}

	min_minutes_end := path.FindShortestDistance(valleys, HasReachedEnd)
	min_minutes_end_start_end := path.FindShortestDistance(valleys, HasReachedEndAgain)

	time_elapsed := time.Since(start)

	log.Printf(`
The fewest number of minutes required to reach the end is %d.
The fewest number of minutes required to reach the end, start, and end again is %d.
Solution generated in %s.`,
		min_minutes_end,
		min_minutes_end_start_end,
		time_elapsed,
	)
}
