package main

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

type Cave struct {
	smoke  [][]int
	height int
	width  int
}

func (cave Cave) ForEach(lambda func(coordinates []int)) {
	for y := 0; y < cave.height; y++ {
		for x := 0; x < cave.width; x++ {
			lambda([]int{y, x})
		}
	}
}

func (cave Cave) GetVal(coordinates []int) int {
	return cave.smoke[coordinates[0]][coordinates[1]]
}

func (scan *Cave) SetVal(coordinates []int, val int) {
	scan.smoke[coordinates[0]][coordinates[1]] = val
}

func (cave Cave) GetAdjacentNeighbourCoordinates(coordinates []int) [][]int {
	adjacent_neighours := make([][]int, 0)

	if coordinates[0] > 0 {
		coordinates := []int{coordinates[0] - 1, coordinates[1]}
		adjacent_neighours = append(adjacent_neighours, coordinates)
	}
	if coordinates[1] > 0 {
		coordinates := []int{coordinates[0], coordinates[1] - 1}
		adjacent_neighours = append(adjacent_neighours, coordinates)
	}

	if coordinates[0] < cave.height-1 {
		coordinates := []int{coordinates[0] + 1, coordinates[1]}
		adjacent_neighours = append(adjacent_neighours, coordinates)
	}
	if coordinates[1] < cave.width-1 {
		coordinates := []int{coordinates[0], coordinates[1] + 1}
		adjacent_neighours = append(adjacent_neighours, coordinates)
	}

	return adjacent_neighours
}

func (cave *Cave) GetBasinSize(coordinates []int) int {
	explore_queue, curr_point := [][]int{coordinates}, coordinates
	basin_size := 0

	cave.SetVal(curr_point, 9)

	for len(explore_queue) > 0 {
		curr_point, explore_queue = explore_queue[0], explore_queue[1:]
		basin_size += 1

		for _, neighbour := range cave.GetAdjacentNeighbourCoordinates(curr_point) {
			if cave.GetVal(neighbour) < 9 {
				cave.SetVal(neighbour, 9)
				explore_queue = append(explore_queue, neighbour)
			}
		}
	}

	return basin_size
}

func main() {
	start := time.Now()

	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)

	cave := Cave{}
	risk_level_sum, basin_sizes, lowest_points := 0, make([]int, 0), make([][]int, 0)

	for scanner.Scan() {
		smoke_row := strings.Split(scanner.Text(), "")
		smoke_ints := make([]int, 0)

		for _, char := range smoke_row {
			val, _ := strconv.Atoi(char)
			smoke_ints = append(smoke_ints, val)
		}

		cave.smoke = append(cave.smoke, smoke_ints)
	}

	cave.height, cave.width = len(cave.smoke), len(cave.smoke[0])

	cave.ForEach(func(coordinates []int) {
		is_smallest := true

		for _, neighbour := range cave.GetAdjacentNeighbourCoordinates(coordinates) {
			if cave.GetVal(neighbour) <= cave.GetVal(coordinates) {
				is_smallest = false
			}
		}

		if is_smallest {
			lowest_points = append(lowest_points, coordinates)
			risk_level_sum += cave.GetVal(coordinates) + 1
		}
	})

	for _, coordinates := range lowest_points {
		basin_sizes = append(basin_sizes, cave.GetBasinSize(coordinates))
	}

	sort.Slice(basin_sizes, func(a, b int) bool {
		return basin_sizes[b] < basin_sizes[a]
	})
	basin_product := basin_sizes[0] * basin_sizes[1] * basin_sizes[2]

	time_elapsed := time.Since(start)

	log.Printf(`
The sum of the risk levels is %d.
The product of the basin sizes is %d.
Solution generated in %s.`,
		risk_level_sum,
		basin_product,
		time_elapsed,
	)
}
