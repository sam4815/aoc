package main

import (
	"bufio"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

type BeaconSensorPair struct {
	beacon []int
	sensor []int
	radius int
}

func createGrid(x, y int) [][]string {
	grid := make([][]string, y)
	for i := range grid {
		grid[i] = make([]string, x)
		for j := range grid[i] {
			grid[i][j] = "."
		}
	}

	return grid
}

func parseCoordinates(coordinates_string string) []int {
	x_y_string := strings.Split(coordinates_string, "at")[1]
	x_y_slice := strings.Split(x_y_string, ", ")
	x_string, y_string := x_y_slice[0], x_y_slice[1]
	x_coordinate := strings.Split(x_string, "=")[1]
	y_coordinate := strings.Split(y_string, "=")[1]

	x, _ := strconv.Atoi(x_coordinate)
	y, _ := strconv.Atoi(y_coordinate)

	return []int{x, y}
}

func calculateManhattanDistance(x, y []int) int {
	diff_x := int(math.Abs(float64(x[0]) - float64(y[0])))
	diff_y := int(math.Abs(float64(x[1]) - float64(y[1])))
	return diff_x + diff_y
}

func checkForDistressSignal(coordinates []int, beacon_sensor_pairs []BeaconSensorPair, search_bounds []int) bool {
	if coordinates[0] < search_bounds[0] || coordinates[0] > search_bounds[1] {
		return false
	}
	if coordinates[1] < search_bounds[0] || coordinates[1] > search_bounds[1] {
		return false
	}

	is_distress_signal := true

	for _, pair := range beacon_sensor_pairs {
		distance := calculateManhattanDistance(coordinates, pair.sensor)
		if distance <= pair.radius {
			is_distress_signal = false
			break
		}
	}

	return is_distress_signal
}

func checkOuterRings(pairs []BeaconSensorPair, search_bounds []int) []int {
	for _, pair := range pairs {
		sensor, radius := pair.sensor, pair.radius

		for y := sensor[1] - radius; y <= sensor[1]+radius; y++ {
			diff_man := int(math.Abs(float64(y) - float64(sensor[1])))
			x_span := radius - diff_man

			if y == sensor[1]-radius {
				above := []int{sensor[0], y - 1}
				if checkForDistressSignal(above, pairs, search_bounds) {
					return above
				}
			} else if y == sensor[1]+radius {
				below := []int{sensor[0], y + 1}
				if checkForDistressSignal(below, pairs, search_bounds) {
					return below
				}
			} else {
				left := []int{sensor[0] - x_span - 1, y}
				if checkForDistressSignal(left, pairs, search_bounds) {
					return left
				}
				right := []int{sensor[0] + x_span + 1, y}
				if checkForDistressSignal(right, pairs, search_bounds) {
					return right
				}
			}
		}
	}

	return []int{0, 0}
}

func markGrid(coordinates []int, grid *[][]string, character string) {
	if coordinates[1] < 0 || coordinates[1] >= len(*grid) {
		return
	}
	if coordinates[0] < 0 || coordinates[0] >= len((*grid)[coordinates[1]]) {
		return
	}
	// Don't overwrite S or B with #
	if character == "#" && ((*grid)[coordinates[1]][coordinates[0]] == "B" || (*grid)[coordinates[1]][coordinates[0]] == "S") {
		return
	}

	(*grid)[coordinates[1]][coordinates[0]] = character
}

func markSensorAndBeacon(beacon_sensor_pair BeaconSensorPair, grid *[][]string) {
	sensor := beacon_sensor_pair.sensor
	beacon := beacon_sensor_pair.beacon
	radius := beacon_sensor_pair.radius
	grid_height := len(*grid)

	for y := sensor[1] - radius; y <= sensor[1]+radius; y++ {
		if y < 0 || y >= grid_height {
			continue
		}

		diff_man := int(math.Abs(float64(y) - float64(sensor[1])))
		x_span := radius - diff_man

		for x := sensor[0] - x_span; x <= sensor[0]+x_span; x++ {
			markGrid([]int{x, y}, grid, "#")
		}
	}

	markGrid(sensor, grid, "S")
	markGrid(beacon, grid, "B")
}

func calculateXBounds(beacon_sensor_pairs []BeaconSensorPair) (int, int) {
	min_x := math.MaxInt
	max_x := math.MinInt

	for i := 0; i < len(beacon_sensor_pairs); i++ {
		if beacon_sensor_pairs[i].sensor[0] < min_x {
			min_x = beacon_sensor_pairs[i].sensor[0]
		} else if beacon_sensor_pairs[i].sensor[0] > max_x {
			max_x = beacon_sensor_pairs[i].sensor[0]
		} else if beacon_sensor_pairs[i].beacon[0] < min_x {
			min_x = beacon_sensor_pairs[i].beacon[0]
		} else if beacon_sensor_pairs[i].beacon[0] > max_x {
			max_x = beacon_sensor_pairs[i].beacon[0]
		}
	}

	return min_x, max_x
}

func main() {
	start := time.Now()

	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)

	beacon_sensors := make([]BeaconSensorPair, 0)

	for scanner.Scan() {
		reading := strings.Split(scanner.Text(), ":")
		sensor_string, beacon_string := reading[0], reading[1]
		sensor, beacon := parseCoordinates(sensor_string), parseCoordinates(beacon_string)
		radius := calculateManhattanDistance(sensor, beacon)

		beacon_sensor_pair := BeaconSensorPair{sensor: sensor, beacon: beacon, radius: radius}
		beacon_sensors = append(beacon_sensors, beacon_sensor_pair)
	}

	min_x, max_x := calculateXBounds(beacon_sensors)
	x_offset := -min_x
	grid_width := (max_x - min_x) + x_offset + 1

	grid := createGrid(grid_width, 1)
	target_y, y_offset := 2000000, -2000000

	for _, pair := range beacon_sensors {
		offset_beacon := []int{pair.beacon[0] + x_offset, pair.beacon[1] + y_offset}
		offset_sensor := []int{pair.sensor[0] + x_offset, pair.sensor[1] + y_offset}
		offset_pair := BeaconSensorPair{sensor: offset_sensor, beacon: offset_beacon, radius: pair.radius}

		markSensorAndBeacon(offset_pair, &grid)
	}

	num_empty_positions := 0
	for y := range grid {
		if y-y_offset == target_y {
			for x := 0; x < len(grid[y]); x++ {
				if grid[y][x] == "#" {
					num_empty_positions++
				}
			}
		}
	}

	search_bounds := []int{0, 4000000}
	// Instead of checking the entire grid, since there's only one possible location for
	// the distress signal, it follows that it borders one of the beacon-sensor circles
	distress_signal := checkOuterRings(beacon_sensors, search_bounds)
	tuning_frequency := distress_signal[0]*4000000 + distress_signal[1]

	time_elapsed := time.Since(start)

	log.Printf(`
The number of positions that cannot contain a beacon in row 2000000 is %d.
The tuning frequency is %d.
Solution generated in %s.`,
		num_empty_positions,
		tuning_frequency,
		time_elapsed,
	)
}
