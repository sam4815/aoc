package main

import (
	"bufio"
	"log"
	"os"
	"strings"
	"time"
)

type CaveSystem map[string][]string

func isBigCave(cave string) bool {
	return int(cave[0]) <= 90
}

func hasVisitedSmallCaveTwice(path []string) bool {
	cave_map := make(map[string]bool)
	for _, visited_cave := range path[1:] {
		if isBigCave(visited_cave) {
			continue
		} else if cave_map[visited_cave] {
			return true
		} else {
			cave_map[visited_cave] = true
		}
	}
	return false
}

func hasVisitedCave(path []string, cave string, num_visits_allowed int) bool {
	if hasVisitedSmallCaveTwice(path) {
		num_visits_allowed -= 1
	}

	count := 0
	for _, visited_cave := range path[1:] {
		if visited_cave == cave {
			count += 1
			if count >= num_visits_allowed {
				return true
			}
		}
	}
	return false
}

func isValidMove(path []string, cave string, num_visits_allowed int) bool {
	if isBigCave(cave) || cave == "end" {
		return true
	} else if cave == "start" {
		return false
	}

	return !hasVisitedCave(path, cave, num_visits_allowed)
}

func findPaths(system CaveSystem, num_visits_allowed int) [][]string {
	paths, path_queue, curr_path := make([][]string, 0), [][]string{{"start"}}, []string{"start"}

	for len(path_queue) > 0 {
		curr_path, path_queue = path_queue[0], path_queue[1:]
		curr_cave := curr_path[len(curr_path)-1]

		if curr_cave == "end" {
			paths = append(paths, curr_path)
			continue
		}

		for _, connected_cave := range system[curr_cave] {
			if !isValidMove(curr_path, connected_cave, num_visits_allowed) {
				continue
			}

			clone_path := make([]string, len(curr_path))
			copy(clone_path, curr_path)

			clone_path = append(clone_path, connected_cave)
			path_queue = append(path_queue, clone_path)
		}
	}

	return paths
}

func main() {
	start := time.Now()
	f, _ := os.Open("input.txt")
	defer f.Close()
	scanner := bufio.NewScanner(f)

	cave_system := CaveSystem{}

	for scanner.Scan() {
		path := strings.Split(scanner.Text(), "-")
		from, to := path[0], path[1]

		if _, exists := cave_system[from]; exists {
			cave_system[from] = append(cave_system[from], to)
		} else {
			cave_system[from] = []string{to}
		}

		if _, exists := cave_system[to]; exists {
			cave_system[to] = append(cave_system[to], from)
		} else {
			cave_system[to] = []string{from}
		}
	}

	paths_with_single_small_visit := findPaths(cave_system, 1)
	paths_with_two_small_visits := findPaths(cave_system, 2)

	time_elapsed := time.Since(start)

	log.Printf(`
The number of distinct paths when small caves can only be visited once is %d.
The number of distinct paths when a single small cave can be visited twice is %d.
Solution generated in %s.`,
		len(paths_with_single_small_visit),
		len(paths_with_two_small_visits),
		time_elapsed,
	)
}
