package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func createGrid(x int) [][]int {
	grid := make([][]int, x)
	for i := range grid {
		grid[i] = make([]int, x)
	}

	return grid
}

func parseCoordinates(coordinates string) (int, int) {
	split_coordinates := strings.Split(coordinates, ",")
	x, _ := strconv.Atoi(split_coordinates[0])
	y, _ := strconv.Atoi(split_coordinates[1])

	return x, y
}

func drawLine(x1 int, y1 int, x2 int, y2 int, grid *[][]int) {
	(*grid)[y1][x1]++

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

		(*grid)[y1][x1]++
	}
}

func countOverlaps(grid [][]int) int {
	overlaps := 0

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] >= 2 {
				overlaps++
			}
		}
	}

	return overlaps
}

func main() {
	start := time.Now()

	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)

	naive_grid := createGrid(1000)
	full_grid := createGrid(1000)

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " -> ")
		x1, y1 := parseCoordinates(line[0])
		x2, y2 := parseCoordinates(line[1])

		if x1 == x2 || y1 == y2 {
			drawLine(x1, y1, x2, y2, &naive_grid)
		}

		drawLine(x1, y1, x2, y2, &full_grid)
	}

	num_naive_overlaps := countOverlaps(naive_grid)
	num_overlaps := countOverlaps(full_grid)

	time_elapsed := time.Since(start)

	log.Printf(`
The number of points where at least two horizontal and vertical lines overlap is %d.
The number of points where at least two lines overlap is %d.
Solution generated in %s.`,
		num_naive_overlaps,
		num_overlaps,
		time_elapsed,
	)
}
