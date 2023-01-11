package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type Cavern struct {
	flashes     int
	octopuses   [][]Octopus
	all_flashed bool
}

type Octopus struct {
	flashed bool
	energy  int
}

func (cavern Cavern) GetNeighbours(position []int) [][]int {
	valid_positions := make([][]int, 0)
	possible_positions := [][]int{
		{position[0] - 1, position[1] - 1},
		{position[0] - 1, position[1]},
		{position[0] - 1, position[1] + 1},
		{position[0], position[1] - 1},
		{position[0], position[1] + 1},
		{position[0] + 1, position[1] - 1},
		{position[0] + 1, position[1]},
		{position[0] + 1, position[1] + 1},
	}

	for _, position := range possible_positions {
		if position[0] >= 0 && position[1] >= 0 && position[0] < len(cavern.octopuses) && position[1] < len(cavern.octopuses[0]) {
			valid_positions = append(valid_positions, position)
		}
	}

	return valid_positions
}

func (cavern *Cavern) ResetFlashed() {
	flashed_count := 0

	for i := range cavern.octopuses {
		for j := range cavern.octopuses[i] {
			if cavern.octopuses[i][j].flashed {
				cavern.octopuses[i][j].flashed = false
				cavern.octopuses[i][j].energy = 0

				flashed_count += 1
			}
		}
	}

	if flashed_count == len(cavern.octopuses)*len(cavern.octopuses[0]) {
		cavern.all_flashed = true
	}
}

func (cavern *Cavern) FlashOctopus(octopus_position []int) {
	cavern.octopuses[octopus_position[0]][octopus_position[1]].flashed = true
	cavern.flashes += 1

	for _, neighbour := range cavern.GetNeighbours(octopus_position) {
		cavern.IncrementOctopus(neighbour)
	}
}

func (cavern *Cavern) IncrementOctopus(octopus_position []int) {
	if cavern.octopuses[octopus_position[0]][octopus_position[1]].flashed {
		return
	}

	cavern.octopuses[octopus_position[0]][octopus_position[1]].energy += 1

	if cavern.octopuses[octopus_position[0]][octopus_position[1]].energy > 9 {
		cavern.FlashOctopus(octopus_position)
	}
}

func (cavern *Cavern) Step() {
	for i := range cavern.octopuses {
		for j := range cavern.octopuses[i] {
			cavern.IncrementOctopus([]int{i, j})
		}
	}

	cavern.ResetFlashed()
}

func main() {
	start := time.Now()

	f, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(f)
	cavern := Cavern{flashes: 0, octopuses: make([][]Octopus, 0)}

	for scanner.Scan() {
		cavern.octopuses = append(cavern.octopuses, make([]Octopus, 0))

		for _, energy_str := range strings.Split(scanner.Text(), "") {
			energy, _ := strconv.Atoi(energy_str)
			octopus := Octopus{energy: energy, flashed: false}
			cavern.octopuses[len(cavern.octopuses)-1] = append(cavern.octopuses[len(cavern.octopuses)-1], octopus)
		}
	}

	for i := 0; i < 100; i++ {
		cavern.Step()
	}

	flashes_after_100_steps, simultaneous_flash_step := cavern.flashes, 100

	for !cavern.all_flashed {
		cavern.Step()
		simultaneous_flash_step += 1
	}

	time_elapsed := time.Since(start)

	log.Printf(`
After 100 steps, the octopuses have flashed %d times.
All of the octopuses flash at the same time during step %d.
Solution generated in %s.`,
		flashes_after_100_steps,
		simultaneous_flash_step,
		time_elapsed,
	)
}
