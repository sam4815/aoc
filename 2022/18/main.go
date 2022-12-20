package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type Scan struct {
	vals    [][][]int
	pockets [][]int
	size    int
}

func createScan(size int) Scan {
	x := make([][][]int, size)
	for i := range x {
		y := make([][]int, size)
		x[i] = y
		for j := range y {
			x[i][j] = make([]int, size)
		}
	}

	return Scan{vals: x, size: size}
}

func (scan Scan) ForEach(lambda func(coordinates []int)) {
	for z := 0; z < scan.size; z++ {
		for y := 0; y < scan.size; y++ {
			for x := 0; x < scan.size; x++ {
				lambda([]int{x, y, z})
			}
		}
	}
}

func (scan Scan) GetVal(coordinates []int) int {
	return scan.vals[coordinates[0]][coordinates[1]][coordinates[2]]
}

func (scan *Scan) SetVal(x int, y int, z int, val int) {
	scan.vals[x][y][z] = val
}

func addX(coordinates []int) []int {
	return []int{coordinates[0] + 1, coordinates[1], coordinates[2]}
}
func subX(coordinates []int) []int {
	return []int{coordinates[0] - 1, coordinates[1], coordinates[2]}
}
func addY(coordinates []int) []int {
	return []int{coordinates[0], coordinates[1] + 1, coordinates[2]}
}
func subY(coordinates []int) []int {
	return []int{coordinates[0], coordinates[1] - 1, coordinates[2]}
}
func addZ(coordinates []int) []int {
	return []int{coordinates[0], coordinates[1], coordinates[2] + 1}
}
func subZ(coordinates []int) []int {
	return []int{coordinates[0], coordinates[1], coordinates[2] - 1}
}

func (scan Scan) GetAdjacentNeighbourCoordinates(coordinates []int) [][]int {
	adjacent_neighours := make([][]int, 0)

	if coordinates[0] > 0 {
		adjacent_neighours = append(adjacent_neighours, subX(coordinates))
	}
	if coordinates[1] > 0 {
		adjacent_neighours = append(adjacent_neighours, subY(coordinates))
	}
	if coordinates[2] > 0 {
		adjacent_neighours = append(adjacent_neighours, subZ(coordinates))
	}

	if coordinates[0] < scan.size-1 {
		adjacent_neighours = append(adjacent_neighours, addX(coordinates))
	}
	if coordinates[1] < scan.size-1 {
		adjacent_neighours = append(adjacent_neighours, addY(coordinates))
	}
	if coordinates[2] < scan.size-1 {
		adjacent_neighours = append(adjacent_neighours, addZ(coordinates))
	}

	return adjacent_neighours
}

func (scan Scan) TestPocket(coordinates []int) bool {
	is_pocket := true

	for _, fun := range []func(coordinates []int) []int{addX, subX, addY, subY, addZ, subZ} {
		step := []int{coordinates[0], coordinates[1], coordinates[2]}
		for true {
			next_step := fun(step)
			i, j, k := next_step[0], next_step[1], next_step[2]
			// If this is the end of the grid, it's not a pocket
			if i < 0 || j < 0 || k < 0 {
				is_pocket = false
				break
			}
			if i >= scan.size || j >= scan.size || k >= scan.size {
				is_pocket = false
				break
			}
			// If this is the droplet, stoplet
			if scan.GetVal(next_step) == 1 {
				break
			}
			step[0], step[1], step[2] = i, j, k
		}
	}

	return is_pocket
}

func (scan *Scan) FindPockets() {
	pockets := make([][]int, 0)

	scan.ForEach(func(coordinates []int) {
		if scan.GetVal(coordinates) != 0 {
			return
		}

		is_pocket := scan.TestPocket(coordinates)

		if is_pocket {
			pockets = append(pockets, coordinates)
		}
	})

	scan.pockets = pockets
}

func (scan Scan) IsPocket(coordinates []int) bool {
	is_pocket := false

	for _, pockets := range scan.pockets {
		if pockets[0] == coordinates[0] && pockets[1] == coordinates[1] && pockets[2] == coordinates[2] {
			is_pocket = true
		}
	}

	return is_pocket
}

func parseCoordinates(coordinates_string string) []int {
	coordinates := make([]int, 3)
	split_coordinates := strings.Split(coordinates_string, ",")
	for i := range coordinates {
		num, _ := strconv.Atoi(split_coordinates[i])
		coordinates[i] = num
	}

	return coordinates
}

func main() {
	start := time.Now()

	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)

	scan := createScan(25)
	surface_area, exterior_surface_area := 0, 0

	for scanner.Scan() {
		coordinates := parseCoordinates(scanner.Text())
		x, y, z := coordinates[0], coordinates[1], coordinates[2]
		scan.SetVal(x, y, z, 1)
	}

	scan.FindPockets()

	scan.ForEach(func(coordinates []int) {
		if scan.GetVal(coordinates) != 1 {
			return
		}

		faces_touching_droplet := 0
		faces_touching_pocket := 0
		for _, neighbour := range scan.GetAdjacentNeighbourCoordinates(coordinates) {
			if scan.GetVal(neighbour) == 1 {
				faces_touching_droplet += 1
			}
			if scan.IsPocket(neighbour) {
				faces_touching_pocket += 1
			}
		}

		surface_area += 6 - faces_touching_droplet
		exterior_surface_area += 6 - faces_touching_droplet - faces_touching_pocket
	})

	time_elapsed := time.Since(start)

	log.Printf(`
The surface area of the lava droplet is %d.
The exterior surface area of the lava droplet is %d.
Solution generated in %s.`,
		surface_area,
		exterior_surface_area,
		time_elapsed,
	)
}
