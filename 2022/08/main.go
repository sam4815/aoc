package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"time"
)

func checkUp(i int, j int, trees [][]int) (bool, int) {
	tree := trees[i][j]
	visible := true
	num_trees_seen := 0

	for x := i - 1; x >= 0; x-- {
		num_trees_seen++
		if trees[x][j] >= tree {
			visible = false
			break
		}
	}

	return visible, num_trees_seen
}

func checkDown(i int, j int, trees [][]int) (bool, int) {
	tree := trees[i][j]
	visible := true
	num_trees_seen := 0

	for x := i + 1; x < len(trees); x++ {
		num_trees_seen++
		if trees[x][j] >= tree {
			visible = false
			break
		}
	}

	return visible, num_trees_seen
}

func checkLeft(i int, j int, trees [][]int) (bool, int) {
	tree := trees[i][j]
	visible := true
	num_trees_seen := 0

	for x := j - 1; x >= 0; x-- {
		num_trees_seen++
		if trees[i][x] >= tree {
			visible = false
			break
		}
	}

	return visible, num_trees_seen
}

func checkRight(i int, j int, trees [][]int) (bool, int) {
	tree := trees[i][j]
	visible := true
	num_trees_seen := 0

	for x := j + 1; x < len(trees[i]); x++ {
		num_trees_seen++
		if trees[i][x] >= tree {
			visible = false
			break
		}
	}

	return visible, num_trees_seen
}

func main() {
	start := time.Now()

	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)

	num_visible := 0
	highest_scenic_score := 0
	trees := make([][]int, 0)

	for scanner.Scan() {
		buffer := scanner.Text()
		tree_row := make([]int, len(buffer))

		for pos := range buffer {
			tree_height, err := strconv.Atoi(string(buffer[pos]))
			if err != nil {
				log.Fatal(err)
			}

			tree_row[pos] = tree_height
		}

		trees = append(trees, tree_row)
	}

	// All trees on outside are visible
	num_visible += (2 * len(trees)) + (2 * (len(trees[0]) - 2))

	for i := 1; i < len(trees)-1; i++ {
		for j := 1; j < len(trees[i])-1; j++ {
			up_visible, up_num_visible := checkUp(i, j, trees)
			down_visible, down_num_visible := checkDown(i, j, trees)
			left_visible, left_num_visible := checkLeft(i, j, trees)
			right_visible, right_num_visible := checkRight(i, j, trees)

			if up_visible || down_visible || left_visible || right_visible {
				num_visible += 1
			}

			scenic_score := up_num_visible * down_num_visible * left_num_visible * right_num_visible
			if scenic_score > highest_scenic_score {
				highest_scenic_score = scenic_score
			}
		}
	}

	time_elapsed := time.Since(start)

	log.Printf(`
The number of visible trees is %d.
The highest possible scenic score is %d.
Solution generated in %s.`,
		num_visible,
		highest_scenic_score,
		time_elapsed,
	)
}
