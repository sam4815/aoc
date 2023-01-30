package main

import (
	"io/ioutil"
	"log"
	"strings"
	"time"
)

type Seafloor [][]string

func (seafloor Seafloor) Get(i, j int) string {
	return seafloor[i%len(seafloor)][j%len(seafloor[0])]
}

func (seafloor *Seafloor) Set(i, j int, val string) {
	(*seafloor)[i%len(*seafloor)][j%len((*seafloor)[0])] = val
}

func (seafloor Seafloor) Clone() Seafloor {
	clone := make(Seafloor, len(seafloor))
	for i, row := range seafloor {
		clone[i] = make([]string, len(row))
		copy(clone[i], row)
	}

	return clone
}

func (seafloor Seafloor) MoveEast() (Seafloor, bool) {
	cucumber_moved := false
	next_seafloor := seafloor.Clone()

	for i := 0; i < len(seafloor); i++ {
		for j := 0; j < len(seafloor[0]); j++ {
			if seafloor.Get(i, j) == "." || seafloor.Get(i, j) == "v" {
				continue
			}

			if seafloor.Get(i, j+1) == "." {
				next_seafloor.Set(i, j+1, ">")
				next_seafloor.Set(i, j, ".")

				cucumber_moved = true
			}
		}
	}

	return next_seafloor, cucumber_moved
}

func (seafloor Seafloor) MoveSouth() (Seafloor, bool) {
	cucumber_moved := false
	next_seafloor := seafloor.Clone()

	for i := 0; i < len(seafloor); i++ {
		for j := 0; j < len(seafloor[0]); j++ {
			if seafloor.Get(i, j) == "." || seafloor.Get(i, j) == ">" {
				continue
			}

			if seafloor.Get(i+1, j) == "." {
				next_seafloor.Set(i+1, j, "v")
				next_seafloor.Set(i, j, ".")

				cucumber_moved = true
			}
		}
	}

	return next_seafloor, cucumber_moved
}

func (seafloor Seafloor) Step() (Seafloor, bool) {
	shifted_east, cucumber_moved_east := seafloor.MoveEast()
	shifted_south, cucumber_moved_south := shifted_east.MoveSouth()

	return shifted_south, cucumber_moved_east || cucumber_moved_south
}

func main() {
	start := time.Now()
	f, _ := ioutil.ReadFile("input.txt")
	seafloor_str := strings.Split(strings.TrimSpace(string(f)), "\n")
	seafloor := make(Seafloor, 0)
	for _, seafloor_row_str := range seafloor_str {
		seafloor = append(seafloor, strings.Split(seafloor_row_str, ""))
	}

	step, cucumber_moved := 0, true
	for cucumber_moved {
		seafloor, cucumber_moved = seafloor.Step()
		step += 1
	}

	time_elapsed := time.Since(start)

	log.Printf(`
The first step on which no sea cucumbers move is %d.
Solution generated in %s.`,
		step,
		time_elapsed,
	)
}
