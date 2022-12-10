package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func initializeStacks(stacks_input []string) [][]string {
	num_stacks := len(stacks_input[0])/4 + 1
	stacks := make([][]string, num_stacks)

	for i := 0; i < len(stacks_input)-1; i++ {
		for j := 0; j < num_stacks; j++ {
			char := string(stacks_input[i][j*4+1])
			if char != " " {
				stacks[j] = append([]string{char}, stacks[j]...)
			}
		}
	}

	return stacks
}

func toInt(str string) int {
	int, err := strconv.Atoi(str)
	if err != nil {
		log.Fatal(err)
	}

	return int
}

func main() {
	start := time.Now()

	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)

	stacks_input := make([]string, 0)
	stacks_9000 := make([][]string, 0)
	stacks_9001 := make([][]string, 0)
	stacks_initialized := false

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			stacks_9000 = initializeStacks(stacks_input)
			stacks_9001 = initializeStacks(stacks_input)
			stacks_initialized = true
			continue
		}

		if !stacks_initialized {
			stacks_input = append(stacks_input, line)
			continue
		}

		instruction := strings.Split(line, " ")
		num_move, from, to := toInt(instruction[1]), toInt(instruction[3])-1, toInt(instruction[5])-1

		for i := 0; i < num_move; i++ {
			fromIdx := len(stacks_9000[from]) - 1
			stacks_9000[to] = append(stacks_9000[to], stacks_9000[from][fromIdx])
			stacks_9000[from] = stacks_9000[from][:fromIdx]
		}

		fromIdx := len(stacks_9001[from]) - num_move
		stacks_9001[to] = append(stacks_9001[to], stacks_9001[from][fromIdx:]...)
		stacks_9001[from] = stacks_9001[from][:fromIdx]
	}

	num_stacks := len(stacks_9000)
	top_items_9000 := make([]string, num_stacks)
	top_items_9001 := make([]string, num_stacks)

	for i := 0; i < num_stacks; i++ {
		top_item_9000 := stacks_9000[i][len(stacks_9000[i])-1]
		top_items_9000[i] = top_item_9000

		top_item_9001 := stacks_9001[i][len(stacks_9001[i])-1]
		top_items_9001[i] = top_item_9001
	}

	time_elapsed := time.Since(start)

	log.Printf(`
The crates that end up on the top of each stack using the CrateMover 9000 are %s.
The crates that end up on the top of each stack using the CrateMover 9001 are %s.
Solution generated in %s.`,
		strings.Join(top_items_9000, ""),
		strings.Join(top_items_9001, ""),
		time_elapsed,
	)
}
