package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func createGrid(x int) [][]string {
	grid := make([][]string, x)
	for i := range grid {
		grid[i] = make([]string, x)
		for j := range grid[i] {
			grid[i][j] = "."
		}
	}

	return grid
}

func parseCoordinates(coordinates string) (int, int) {
	split_coordinates := strings.Split(coordinates, ",")
	x, _ := strconv.Atoi(split_coordinates[0])
	y, _ := strconv.Atoi(split_coordinates[1])

	return x, y
}

func drawRock(x1 int, y1 int, x2 int, y2 int, grid *[][]string) {
	(*grid)[y1][x1] = "#"

	for x1 != x2 || y1 != y2 {
		if x1 < x2 {
			x1++
		}
		if x2 < x1 {
			x1--
		}
		if y1 < y2 {
			y1++
		}
		if y2 < y1 {
			y1--
		}

		(*grid)[y1][x1] = "#"
	}
}

func updateSandCoordinates(grid *[][]string, sand_coordinates *[]int, new_coordinates *[]int) {
	(*grid)[(*sand_coordinates)[1]][(*sand_coordinates)[0]] = "."
	(*sand_coordinates)[0], (*sand_coordinates)[1] = (*new_coordinates)[0], (*new_coordinates)[1]
	(*grid)[(*sand_coordinates)[1]][(*sand_coordinates)[0]] = "o"
}

func addSand(grid *[][]string) []int {
	sand_coordinates := []int{500, 0}
	(*grid)[sand_coordinates[1]][sand_coordinates[0]] = "o"

	for true {
		down := []int{sand_coordinates[0], sand_coordinates[1] + 1}
		down_left := []int{sand_coordinates[0] - 1, sand_coordinates[1] + 1}
		down_right := []int{sand_coordinates[0] + 1, sand_coordinates[1] + 1}

		if (*grid)[down[1]][down[0]] == "." {
			updateSandCoordinates(grid, &sand_coordinates, &down)
		} else if (*grid)[down_left[1]][down_left[0]] == "." {
			updateSandCoordinates(grid, &sand_coordinates, &down_left)
		} else if (*grid)[down_right[1]][down_right[0]] == "." {
			updateSandCoordinates(grid, &sand_coordinates, &down_right)
		} else {
			break
		}
	}

	return sand_coordinates
}

func main() {
	start := time.Now()

	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)

	finite_grid := createGrid(1000)
	lowest_y := 0

	for scanner.Scan() {
		path := strings.Split(scanner.Text(), " -> ")
		for i := 0; i < len(path)-1; i++ {
			x1, y1 := parseCoordinates(path[i])
			x2, y2 := parseCoordinates(path[i+1])
			drawRock(x1, y1, x2, y2, &finite_grid)

			if y1 > lowest_y {
				lowest_y = y1
			} else if y2 > lowest_y {
				lowest_y = y2
			}
		}
	}

	for i := 0; i < len(finite_grid[0]); i++ {
		finite_grid[lowest_y+2][i] = "#"
	}

	sand_without_floor, sand_with_floor, curr_sand := 0, 0, 0

	for true {
		sand_coordinates := addSand(&finite_grid)

		if sand_coordinates[1] >= lowest_y && sand_without_floor == 0 {
			sand_without_floor = curr_sand
		}

		curr_sand++

		if sand_coordinates[1] == 0 {
			sand_with_floor = curr_sand
			break
		}
	}

	time_elapsed := time.Since(start)

	log.Printf(`
The number of units of sand that come to rest without a floor is %d.
The number of units of sand that come to rest with a floor is %d.
Solution generated in %s.`,
		sand_without_floor,
		sand_with_floor,
		time_elapsed,
	)
}
