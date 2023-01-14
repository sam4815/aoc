package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"time"
)

type TargetArea struct {
	x_min, x_max, y_min, y_max int
}

type Point struct{ x, y int }
type Velocity struct{ x, y int }

type Probe struct {
	position Point
	velocity Velocity
	max_y    int
}

func (probe Probe) Inside(area TargetArea) bool {
	return probe.position.x <= area.x_max &&
		probe.position.x >= area.x_min &&
		probe.position.y <= area.y_max &&
		probe.position.y >= area.y_min
}

func (probe *Probe) Step() {
	probe.position.x += probe.velocity.x
	probe.position.y += probe.velocity.y

	if probe.position.y > probe.max_y {
		probe.max_y = probe.position.y
	}

	if probe.velocity.x > 0 {
		probe.velocity.x -= 1
	} else if probe.velocity.x < 0 {
		probe.velocity.x += 1
	}

	probe.velocity.y -= 1
}

func (probe Probe) TargetWithinRange(area TargetArea) bool {
	return probe.position.y >= area.y_min && probe.position.x <= area.x_max
}

func main() {
	start := time.Now()
	f, _ := ioutil.ReadFile("input.txt")

	var x_min, x_max, y_min, y_max int
	fmt.Sscanf(string(f), "target area: x=%d..%d, y=%d..%d", &x_min, &x_max, &y_min, &y_max)
	target_area := TargetArea{x_min, x_max, y_min, y_max}
	num_hit_target, highest_y := 0, math.MinInt

	for x := 0; x < 200; x++ {
		for y := -200; y < 200; y++ {
			probe := Probe{position: Point{0, 0}, velocity: Velocity{x, y}}
			for probe.TargetWithinRange(target_area) {
				probe.Step()

				if probe.Inside(target_area) {
					num_hit_target += 1
					if probe.max_y > highest_y {
						highest_y = probe.max_y
					}

					break
				}
			}
		}
	}

	time_elapsed := time.Since(start)

	log.Printf(`
The highest y position that the probe can reach is %d.
There are %d distinct velocities capable of reaching the target area.
Solution generated in %s.`,
		highest_y,
		num_hit_target,
		time_elapsed,
	)
}
