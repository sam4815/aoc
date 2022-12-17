package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type Volcano struct {
	valves            []Valve
	players           []Player
	minute            int
	pressure_released int
	id                int
}

type Valve struct {
	label         string
	flow_rate     int
	open          bool
	neighbours    []string
	min_distances []int
	targeted      bool
}

type Player struct {
	position        int
	target_distance int
	target_position int
}

func parseNeighbours(description []string) []string {
	neighbours := make([]string, 0)

	for i := len(description) - 1; i > 0; i-- {
		if len(description[i]) == 2 {
			neighbours = append(neighbours, description[i])
		} else if len(description[i]) == 3 {
			neighbours = append(neighbours, description[i][0:2])
		} else {
			break
		}
	}

	return neighbours
}

func getValveSummary(valves []Valve) (int, int) {
	pressure := 0
	num_open := 0

	for i := 0; i < len(valves); i++ {
		if valves[i].open {
			num_open++
			pressure += valves[i].flow_rate
		}
	}

	return pressure, num_open
}

func copyVolcano(volcano Volcano) Volcano {
	copy := Volcano{minute: volcano.minute, pressure_released: volcano.pressure_released}

	valves := make([]Valve, 0)
	for _, valve := range volcano.valves {
		valves = append(valves, Valve{
			label:         valve.label,
			flow_rate:     valve.flow_rate,
			open:          valve.open,
			neighbours:    valve.neighbours,
			min_distances: valve.min_distances,
		})
	}

	players := make([]Player, 0)
	for _, player := range volcano.players {
		players = append(players, Player{
			position:        player.position,
			target_distance: player.target_distance,
			target_position: player.target_position,
		})
	}

	copy.valves = valves
	copy.players = players

	return copy
}

func findValveIndexByLabel(valves *[]Valve, label string) int {
	for i := 0; i < len(*valves); i++ {
		if (*valves)[i].label == label {
			return i
		}
	}
	return -1
}

func getBellmanFord(valves *[]Valve, s_idx int) []int {
	distances := make([]int, len(*valves))
	queue, valve_idx := []int{s_idx}, 0

	for i := 0; i < len(distances); i++ {
		distances[i] = 100
	}

	distances[s_idx] = 0

	for len(queue) > 0 {
		valve_idx, queue = queue[0], queue[1:]
		valve := (*valves)[valve_idx]

		for _, neighbour := range valve.neighbours {
			neighbour_index := findValveIndexByLabel(valves, neighbour)

			temp_distance := distances[valve_idx] + 1
			if temp_distance < distances[neighbour_index] {
				distances[neighbour_index] = temp_distance
				queue = append(queue, neighbour_index)
			}
		}
	}

	return distances
}

func findMinDistances(valves *[]Valve) {
	for i := 0; i < len(*valves); i++ {
		(*valves)[i].min_distances = getBellmanFord(valves, i)
	}
}

func tick(volcano *Volcano) {
	pressure_sum, num_open := getValveSummary(volcano.valves)

	volcano.minute += 1
	volcano.pressure_released += pressure_sum

	// All valves open; nowhere else for players to move
	if num_open == len(volcano.valves) {
		return
	}

	for i := 0; i < len(volcano.players); i++ {
		if volcano.players[i].target_distance > 0 {
			volcano.players[i].target_distance--
			continue
		}

		// Player is in position: open valve, update position, remove target
		if volcano.players[i].target_distance == 0 {
			target_valve_index := volcano.players[i].target_position

			volcano.valves[target_valve_index].open = true
			volcano.players[i].position = target_valve_index
			volcano.players[i].target_position = -1
		}
	}
}

func setPlayerTarget(volcano *Volcano, player_index int, target_index int) {
	// log.Print("TARGETING ", volcano.valves[target_index].label)
	player_position := volcano.players[player_index].position
	distance := volcano.valves[player_position].min_distances[target_index]

	volcano.players[player_index].target_distance = distance
	volcano.players[player_index].target_position = target_index
	volcano.valves[player_position].targeted = true
}

func obtainTargetIndices(volcano Volcano) []int {
	targets := make([]int, 0)
	for i := 0; i < len(volcano.valves); i++ {
		if volcano.valves[i].open || volcano.valves[i].targeted {
			continue
		}

		targets = append(targets, i)
	}

	return targets
}

func getPlayerValvePermutations(volcano Volcano) []Volcano {
	permutations := make([]Volcano, 0)

	for player_index := 0; player_index < len(volcano.players); player_index++ {
		if volcano.players[player_index].target_position >= 0 {
			continue
		}

		target_valve_indices := obtainTargetIndices(volcano)
		for _, valve_index := range target_valve_indices {
			copy := copyVolcano(volcano)
			setPlayerTarget(&copy, player_index, valve_index)
			sub_variations := getPlayerValvePermutations(copy)
			if len(sub_variations) == 0 {
				permutations = append(permutations, copy)
			} else {
				permutations = append(permutations, sub_variations...)
			}
		}
	}

	return permutations
}

func findBestPath(volcano Volcano) int {
	volcano_queue, curr_volcano := []Volcano{volcano}, volcano

	max_pressure_path := 0
	best_by_minute_map := make(map[int]int)

	for len(volcano_queue) > 0 {
		curr_volcano, volcano_queue = volcano_queue[0], volcano_queue[1:]
		// Simulation complete
		if curr_volcano.minute == 30 {
			if curr_volcano.pressure_released > max_pressure_path {
				max_pressure_path = curr_volcano.pressure_released
			}
			continue
		}

		if curr_volcano.pressure_released+50 < best_by_minute_map[curr_volcano.minute] {
			continue
		} else if curr_volcano.pressure_released > best_by_minute_map[curr_volcano.minute] {
			best_by_minute_map[curr_volcano.minute] = curr_volcano.pressure_released
		}

		possible_paths := getPlayerValvePermutations(curr_volcano)
		for i := 0; i < len(possible_paths); i++ {
			tick(&possible_paths[i])
		}

		if len(possible_paths) == 0 {
			copy := copyVolcano(curr_volcano)
			tick(&copy)
			possible_paths = append(possible_paths, copy)
		}

		volcano_queue = append(volcano_queue, possible_paths...)
	}

	return max_pressure_path
}

func main() {
	start := time.Now()

	f, err := os.Open("input.txt")
	if err != nil {
		fmt.Print(err)
	}
	scanner := bufio.NewScanner(f)

	valves := make([]Valve, 0)

	for scanner.Scan() {
		valve := Valve{}

		valve_description := strings.Split(scanner.Text(), " ")
		rate_description := strings.Split(strings.Split(valve_description[4], "=")[1], ";")
		flow_rate, _ := strconv.Atoi(rate_description[0])

		valve.label = valve_description[1]
		valve.flow_rate = flow_rate

		neighbours := parseNeighbours(valve_description)
		valve.neighbours = neighbours

		if valve.flow_rate == 0 {
			valve.open = true
		}

		valves = append(valves, valve)
	}

	findMinDistances(&valves)

	starting_position := 0
	for i, valve := range valves {
		if valve.label == "AA" {
			starting_position = i
		}
	}

	lonely_volcano := Volcano{
		valves:            valves,
		players:           []Player{{position: starting_position, target_position: -1}},
		pressure_released: 0,
		minute:            0,
		id:                1,
	}

	elephant_volcano := Volcano{
		valves: valves,
		players: []Player{
			{position: starting_position, target_position: -1},
			{position: starting_position, target_position: -1},
		},
		pressure_released: 0,
		minute:            4,
		id:                1,
	}

	max_pressure_alone := findBestPath(lonely_volcano)
	max_pressure_with_elephant := findBestPath(elephant_volcano)

	time_elapsed := time.Since(start)

	log.Printf(`
The most pressure that can be released in 30 minutes alone is %d.
The most pressure that can be released in 30 minutes with an elephant is %d.
Solution generated in %s.`,
		max_pressure_alone,
		max_pressure_with_elephant,
		time_elapsed,
	)
}
