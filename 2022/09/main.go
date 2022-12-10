package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func moveDirectionOne(pos [2]int, direction string) [2]int {
	switch direction {
	case "U":
		pos[0] = pos[0] - 1
	case "D":
		pos[0] = pos[0] + 1
	case "L":
		pos[1] = pos[1] - 1
	case "R":
		pos[1] = pos[1] + 1
	}

	return pos
}

func followHead(h_pos [2]int, t_pos [2]int) [2]int {
	// If H is to the right of T
	if h_pos[1] >= t_pos[1]+2 {
		t_pos[1]++
		if h_pos[0] > t_pos[0] {
			t_pos[0]++
		} else if h_pos[0] < t_pos[0] {
			t_pos[0]--
		}
		return t_pos
	}
	// If H is to the left of T
	if h_pos[1] <= t_pos[1]-2 {
		t_pos[1]--
		if h_pos[0] > t_pos[0] {
			t_pos[0]++
		} else if h_pos[0] < t_pos[0] {
			t_pos[0]--
		}
		return t_pos
	}
	// If H is below T
	if h_pos[0] >= t_pos[0]+2 {
		t_pos[0]++
		if h_pos[1] > t_pos[1] {
			t_pos[1]++
		} else if h_pos[1] < t_pos[1] {
			t_pos[1]--
		}
		return t_pos
	}
	// If H is above T
	if h_pos[0] <= t_pos[0]-2 {
		t_pos[0]--
		if h_pos[1] > t_pos[1] {
			t_pos[1]++
		} else if h_pos[1] < t_pos[1] {
			t_pos[1]--
		}
		return t_pos
	}

	return t_pos
}

func main() {
	start := time.Now()

	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)

	num_visited := 0

	rope := make([][]string, 1000)
	knot_positions := make([][2]int, 10)
	for i := range rope {
		rope[i] = make([]string, 1000)
	}
	for i := range knot_positions {
		knot_positions[i] = [2]int{500, 500}
	}
	rope[knot_positions[0][0]][knot_positions[0][1]] = "#"

	// log.Print(h_pos, t_pos)

	for scanner.Scan() {
		movement := strings.Split(scanner.Text(), " ")
		direction := movement[0]
		magnitude, err := strconv.Atoi(movement[1])
		if err != nil {
			log.Fatal(err)
		}

		for i := 0; i < magnitude; i++ {
			knot_positions[0] = moveDirectionOne(knot_positions[0], direction)

			for j := 1; j < len(knot_positions); j++ {
				knot_positions[j] = followHead(knot_positions[j-1], knot_positions[j])
				if j == len(knot_positions)-1 {
					rope[knot_positions[j][0]][knot_positions[j][1]] = "#"
				}
			}
		}
	}

	for i := range rope {
		for j := range rope[i] {
			if rope[i][j] == "#" {
				num_visited++
			}
		}
	}

	time_elapsed := time.Since(start)

	log.Printf(`
The number of positions visited by the tail is %d.
Solution generated in %s.`,
		num_visited,
		time_elapsed,
	)
}
