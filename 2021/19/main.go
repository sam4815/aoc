package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
	"time"
)

type Point struct {
	x, y, z int
}

type Beacons []Point

type Scanner struct {
	id        int
	origin    Point
	beacons   Beacons
	rotations []string
}

func (point_a Point) Equals(point_b Point) bool {
	return point_a.x == point_b.x && point_a.y == point_b.y && point_a.z == point_b.z
}

func (point_a Point) Add(point_b Point) Point {
	x, y, z := point_a.x+point_b.x, point_a.y+point_b.y, point_a.z+point_b.z
	return Point{x, y, z}
}

func (point_a Point) ManhattanDistance(point_b Point) int {
	x_diff := point_a.x - point_b.x
	y_diff := point_a.y - point_b.y
	z_diff := point_a.z - point_b.z
	return x_diff + y_diff + z_diff
}

func (point Point) Clone() Point {
	return Point{x: point.x, y: point.y, z: point.z}
}

func (scanner Scanner) Clone() Scanner {
	clone := Scanner{id: scanner.id, origin: scanner.origin.Clone(), rotations: scanner.rotations, beacons: make(Beacons, 0)}
	for _, beacon := range scanner.beacons {
		clone.beacons = append(clone.beacons, beacon.Clone())
	}
	return clone
}

func (scanner Scanner) Rotations() []Scanner {
	rotations := make([]Scanner, 0)

	for x := 0; x < 4; x++ {
		for z := 0; z < 4; z++ {
			new_scanner := scanner.Clone()
			new_scanner.RotateZ(z)
			new_scanner.RotateX(x)
			rotations = append(rotations, new_scanner)
		}

		for y := 1; y < 4; y += 2 {
			new_scanner := scanner.Clone()
			new_scanner.RotateY(y)
			new_scanner.RotateX(x)
			rotations = append(rotations, new_scanner)
		}
	}

	return rotations
}

func (scanner Scanner) Translations(point Point) []Scanner {
	scanners := make([]Scanner, 0)

	for _, beacon := range scanner.beacons {
		beacons := make([]Point, 0)
		x_diff, y_diff, z_diff := point.x-beacon.x, point.y-beacon.y, point.z-beacon.z
		translation := Point{x_diff, y_diff, z_diff}

		for _, beacon := range scanner.beacons {
			beacons = append(beacons, beacon.Add(translation))
		}

		new_scanner := scanner.Clone()
		new_scanner.beacons = beacons
		new_scanner.origin = scanner.origin.Add(translation)
		scanners = append(scanners, new_scanner)
	}

	return scanners
}

func (point Point) MatrixMul(m [][]int) Point {
	x := m[0][0]*point.x + m[0][1]*point.y + m[0][2]*point.z
	y := m[1][0]*point.x + m[1][1]*point.y + m[1][2]*point.z
	z := m[2][0]*point.x + m[2][1]*point.y + m[2][2]*point.z
	return Point{x, y, z}
}

func (point Point) RotateX() Point {
	return point.MatrixMul([][]int{{1, 0, 0}, {0, 0, -1}, {0, 1, 0}})
}

func (point Point) RotateY() Point {
	return point.MatrixMul([][]int{{0, 0, 1}, {0, 1, 0}, {-1, 0, 0}})
}

func (point Point) RotateZ() Point {
	return point.MatrixMul([][]int{{0, -1, 0}, {1, 0, 0}, {0, 0, 1}})
}

func (scanner *Scanner) RotateX(num_times int) {
	for i := 0; i < num_times; i++ {
		for beacon_index, beacon := range scanner.beacons {
			scanner.beacons[beacon_index] = beacon.RotateX()
		}
		scanner.origin = scanner.origin.RotateX()
		scanner.rotations = append(scanner.rotations, "x")
	}
}

func (scanner *Scanner) RotateY(num_times int) {
	for i := 0; i < num_times; i++ {
		for beacon_index, beacon := range scanner.beacons {
			scanner.beacons[beacon_index] = beacon.RotateY()
		}
		scanner.origin = scanner.origin.RotateY()
		scanner.rotations = append(scanner.rotations, "y")
	}
}

func (scanner *Scanner) RotateZ(num_times int) {
	for i := 0; i < num_times; i++ {
		for beacon_index, beacon := range scanner.beacons {
			scanner.beacons[beacon_index] = beacon.RotateZ()
		}
		scanner.origin = scanner.origin.RotateZ()
		scanner.rotations = append(scanner.rotations, "z")
	}
}

func (scanner *Scanner) Add(point Point) {
	scanner.origin = scanner.origin.Add(point)
	for i, beacon := range scanner.beacons {
		scanner.beacons[i] = beacon.Add(point)
	}
}

func (scanner *Scanner) Rotate(rotations []string) {
	for _, direction := range rotations {
		switch direction {
		case "x":
			scanner.RotateX(1)
		case "y":
			scanner.RotateY(1)
		case "z":
			scanner.RotateZ(1)
		}
	}
}

func (beacons_a Beacons) Match(beacons_b Beacons) bool {
	count := 0
	for _, beacon_a := range beacons_a {
		for _, beacon_b := range beacons_b {
			if beacon_a.Equals(beacon_b) {
				count += 1
				break
			}
		}
	}

	return count >= 12
}

func (scanner_a Scanner) Overlaps(scanner_b Scanner) (bool, Scanner) {
	for _, beacon := range scanner_a.beacons {
		for _, rotated_scanner := range scanner_b.Rotations() {
			for _, translated_scanner := range rotated_scanner.Translations(beacon) {
				if translated_scanner.beacons.Match(scanner_a.beacons) {
					return true, translated_scanner
				}
			}
		}
	}

	return false, scanner_a
}

func (scanner Scanner) Stringify() string {
	return fmt.Sprintf("ID: %d, Origin: %d, rotations: %s", scanner.id, scanner.origin, scanner.rotations)
}

func main() {
	start := time.Now()
	f, _ := ioutil.ReadFile("input.txt")
	scanner_strings := strings.Split(strings.TrimSpace(string(f)), "\n\n")
	scanners := make([]Scanner, 0)

	for i, str := range scanner_strings {
		beacons := make([]Point, 0)
		for _, beacon_str := range strings.Split(str, "\n")[1:] {
			coordinates := strings.Split(beacon_str, ",")
			x, _ := strconv.Atoi(coordinates[0])
			y, _ := strconv.Atoi(coordinates[1])
			z, _ := strconv.Atoi(coordinates[2])
			beacons = append(beacons, Point{x, y, z})
		}
		scanners = append(scanners, Scanner{id: i, origin: Point{0, 0, 0}, beacons: beacons})
	}

	overlapping_scanners := make([][]Scanner, 0)

	for _, scanner_a := range scanners {
		for _, scanner_b := range scanners {
			overlaps, scanner := scanner_a.Overlaps(scanner_b)
			if overlaps {
				overlapping_scanners = append(overlapping_scanners, []Scanner{scanner_a, scanner})
			}
		}
	}

	completed_scanner_map := map[int]Scanner{scanners[0].id: scanners[0]}

	for len(completed_scanner_map) < len(scanner_strings) {
		for _, pair := range overlapping_scanners {
			var reference_zero, next_reference Scanner
			if ref, exists := completed_scanner_map[pair[0].id]; exists {
				reference_zero, next_reference = ref.Clone(), pair[1].Clone()
			} else {
				continue
			}

			next_reference.Rotate(reference_zero.rotations)
			next_reference.Add(reference_zero.origin)

			completed_scanner_map[next_reference.id] = next_reference
		}
	}

	unique_beacons := make(map[Point]bool)
	for _, completed := range completed_scanner_map {
		for _, beacon := range completed.beacons {
			unique_beacons[beacon] = true
		}
	}
	num_unique_beacons := len(unique_beacons)

	max_manhattan_distance := 0
	for _, scanner_a := range completed_scanner_map {
		for _, scanner_b := range completed_scanner_map {
			manhattan_distance := scanner_a.origin.ManhattanDistance(scanner_b.origin)
			if manhattan_distance > max_manhattan_distance {
				max_manhattan_distance = manhattan_distance
			}
		}
	}

	time_elapsed := time.Since(start)

	log.Printf(`
There are %d beacons.
The maximum Manhattan distance is %d units.
Solution generated in %s.`,
		num_unique_beacons,
		max_manhattan_distance,
		time_elapsed,
	)
}
