package main

import (
	"bufio"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

func factorial(i int) int {
	return (i * (i + 1)) / 2
}

func identity(i int) int {
	return i
}

func calculateCost(target_cost int, cost_function func(int) int, positions *[]int) int {
	total_cost := 0

	for i := 0; i < len(*positions); i++ {
		diff := target_cost - (*positions)[i]
		if diff > 0 {
			total_cost += cost_function(diff)
		} else {
			total_cost += cost_function(0 - diff)
		}
	}

	return total_cost
}

func main() {
	start := time.Now()

	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)

	scanner.Scan()
	crab_position_strings := strings.Split(scanner.Text(), ",")
	crab_positions := make([]int, 0)

	for _, position_str := range crab_position_strings {
		position, _ := strconv.Atoi(position_str)
		crab_positions = append(crab_positions, position)
	}

	smallest_constant_cost := math.MaxInt
	smallest_growing_cost := math.MaxInt

	sort.Ints(crab_positions)
	smallest := crab_positions[0]
	largest := crab_positions[len(crab_positions)-1]

	for i := smallest; i < largest; i++ {
		constant_cost := calculateCost(i, identity, &crab_positions)
		if constant_cost < smallest_constant_cost {
			smallest_constant_cost = constant_cost
		}

		growing_cost := calculateCost(i, factorial, &crab_positions)
		if growing_cost < smallest_growing_cost {
			smallest_growing_cost = growing_cost
		}
	}

	time_elapsed := time.Since(start)

	log.Printf(`
The least expensive position if each move costs one costs %d fuel.
The least expensive position if each move grows in cost costs %d fuel.
Solution generated in %s.`,
		smallest_constant_cost,
		smallest_growing_cost,
		time_elapsed,
	)
}
