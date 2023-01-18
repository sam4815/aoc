package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

type Point struct {
	x, y, z int
}

type Cuboid struct {
	swb, seb, net                            Point
	x_min, x_max, y_min, y_max, z_min, z_max int
}

type Cuboids []Cuboid

func (point Point) Clone() Point {
	return Point{point.x, point.y, point.z}
}

func MaxInt(a, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}

func MinInt(a, b int) int {
	if a <= b {
		return a
	} else {
		return b
	}
}

func (cuboid Cuboid) Valid() bool {
	return cuboid.x_min < cuboid.x_max && cuboid.y_min < cuboid.y_max && cuboid.z_min < cuboid.z_max
}

func CreateCuboid(min Point, max Point) Cuboid {
	return Cuboid{
		swb:   Point{min.x, min.y, min.z},
		seb:   Point{max.x, min.y, min.z},
		net:   Point{max.x, max.y, max.z},
		x_min: min.x,
		x_max: max.x,
		y_min: min.y,
		y_max: max.y,
		z_min: min.z,
		z_max: max.z,
	}
}

func (cuboids Cuboids) Add(cuboid Cuboid) Cuboids {
	next_cuboids := make(Cuboids, 0)

	for _, existing_cuboid := range cuboids {
		if existing_cuboid.Intersects(cuboid) || cuboid.Intersects(existing_cuboid) {
			chunks := existing_cuboid.Chunks(cuboid)
			next_cuboids = append(next_cuboids, chunks...)
		} else {
			next_cuboids = append(next_cuboids, existing_cuboid)
		}
	}

	return append(next_cuboids, cuboid)
}

func (cuboids Cuboids) Sub(cuboid Cuboid) Cuboids {
	next_cuboids := make(Cuboids, 0)

	for _, existing_cuboid := range cuboids {
		if existing_cuboid.Intersects(cuboid) || cuboid.Intersects(existing_cuboid) {
			chunks := existing_cuboid.Chunks(cuboid)
			next_cuboids = append(next_cuboids, chunks...)
		} else {
			next_cuboids = append(next_cuboids, existing_cuboid)
		}
	}

	return next_cuboids
}

func (cuboid Cuboid) MoveY(y_min, y_max int) Cuboid {
	min, max := cuboid.swb.Clone(), cuboid.net.Clone()
	min.y, max.y = y_min, y_max
	return CreateCuboid(min, max)
}

func (cuboid Cuboid) MoveZ(z_min, z_max int) Cuboid {
	min, max := cuboid.swb.Clone(), cuboid.net.Clone()
	min.z, max.z = z_min, z_max
	return CreateCuboid(min, max)
}

func (cuboid_a Cuboid) Clip(cuboid_b Cuboid) Cuboid {
	min, max := cuboid_a.swb.Clone(), cuboid_a.net.Clone()

	min.x = MaxInt(min.x, cuboid_b.x_min)
	min.y = MaxInt(min.y, cuboid_b.y_min)
	min.z = MaxInt(min.z, cuboid_b.z_min)
	max.x = MinInt(max.x, cuboid_b.x_max)
	max.y = MinInt(max.y, cuboid_b.y_max)
	max.z = MinInt(max.z, cuboid_b.z_max)
	return CreateCuboid(min, max)
}

func (cuboid_a Cuboid) Chunks(cub Cuboid) Cuboids {
	cuboid_b := cub.Clip(cuboid_a)

	if cuboid_b.Volume() == 0 {
		return Cuboids{cuboid_a}
	}
	// Bottom
	a11 := CreateCuboid(cuboid_a.swb, cuboid_b.swb)
	a12 := CreateCuboid(a11.seb, cuboid_b.seb)
	a13 := CreateCuboid(a12.seb, Point{cuboid_a.x_max, cuboid_b.y_min, cuboid_b.z_min})
	a21 := a11.MoveY(cuboid_b.y_min, cuboid_b.y_max)
	a22 := a12.MoveY(cuboid_b.y_min, cuboid_b.y_max)
	a23 := a13.MoveY(cuboid_b.y_min, cuboid_b.y_max)
	a31 := a11.MoveY(cuboid_b.y_max, cuboid_a.y_max)
	a32 := a12.MoveY(cuboid_b.y_max, cuboid_a.y_max)
	a33 := a13.MoveY(cuboid_b.y_max, cuboid_a.y_max)
	// Middle
	b11 := a11.MoveZ(cuboid_b.z_min, cuboid_b.z_max)
	b12 := a12.MoveZ(cuboid_b.z_min, cuboid_b.z_max)
	b13 := a13.MoveZ(cuboid_b.z_min, cuboid_b.z_max)
	b21 := a21.MoveZ(cuboid_b.z_min, cuboid_b.z_max)
	b22 := a22.MoveZ(cuboid_b.z_min, cuboid_b.z_max)
	b23 := a23.MoveZ(cuboid_b.z_min, cuboid_b.z_max)
	b31 := a31.MoveZ(cuboid_b.z_min, cuboid_b.z_max)
	b32 := a32.MoveZ(cuboid_b.z_min, cuboid_b.z_max)
	b33 := a33.MoveZ(cuboid_b.z_min, cuboid_b.z_max)
	// Top
	c11 := b11.MoveZ(cuboid_b.z_max, cuboid_a.z_max)
	c12 := b12.MoveZ(cuboid_b.z_max, cuboid_a.z_max)
	c13 := b13.MoveZ(cuboid_b.z_max, cuboid_a.z_max)
	c21 := b21.MoveZ(cuboid_b.z_max, cuboid_a.z_max)
	c22 := b22.MoveZ(cuboid_b.z_max, cuboid_a.z_max)
	c23 := b23.MoveZ(cuboid_b.z_max, cuboid_a.z_max)
	c31 := b31.MoveZ(cuboid_b.z_max, cuboid_a.z_max)
	c32 := b32.MoveZ(cuboid_b.z_max, cuboid_a.z_max)
	c33 := b33.MoveZ(cuboid_b.z_max, cuboid_a.z_max)

	chunks := Cuboids{a11, a12, a13, a21, a22, a23, a31, a32, a33, b11, b12, b13, b21, b23, b31, b32, b33, c11, c12, c13, c21, c22, c23, c31, c32, c33}

	filtered_chunks := make(Cuboids, 0)
	for _, chunk := range chunks {
		if chunk.Valid() {
			filtered_chunks = append(filtered_chunks, chunk)
		}
	}

	return filtered_chunks
}

func (cuboid_a Cuboid) Intersects(cuboid_b Cuboid) bool {
	x_max_intersects := cuboid_a.x_max >= cuboid_b.x_min && cuboid_a.x_max < cuboid_b.x_max
	x_min_intersects := cuboid_a.x_min >= cuboid_b.x_min && cuboid_a.x_min < cuboid_b.x_max
	x_encloses := cuboid_a.x_min <= cuboid_b.x_max && cuboid_a.x_max >= cuboid_b.x_max
	x_enclosed := cuboid_a.x_min >= cuboid_b.x_max && cuboid_a.x_max <= cuboid_b.x_max

	y_max_intersects := cuboid_a.y_max >= cuboid_b.y_min && cuboid_a.y_max < cuboid_b.y_max
	y_min_intersects := cuboid_a.y_min >= cuboid_b.y_min && cuboid_a.y_min < cuboid_b.y_max
	y_encloses := cuboid_a.y_min <= cuboid_b.y_max && cuboid_a.y_max >= cuboid_b.y_max
	y_enclosed := cuboid_a.y_min >= cuboid_b.y_max && cuboid_a.y_max <= cuboid_b.y_max

	z_max_intersects := cuboid_a.z_max >= cuboid_b.z_min && cuboid_a.z_max < cuboid_b.z_max
	z_min_intersects := cuboid_a.z_min >= cuboid_b.z_min && cuboid_a.z_min < cuboid_b.z_max
	z_encloses := cuboid_a.z_min <= cuboid_b.z_max && cuboid_a.z_max >= cuboid_b.z_max
	z_enclosed := cuboid_a.z_min >= cuboid_b.z_max && cuboid_a.z_max <= cuboid_b.z_max

	x_intersects := x_max_intersects || x_min_intersects || x_encloses || x_enclosed
	y_intersects := y_max_intersects || y_min_intersects || y_encloses || y_enclosed
	z_intersects := z_max_intersects || z_min_intersects || z_encloses || z_enclosed

	return x_intersects && y_intersects && z_intersects
}

func (cuboid Cuboid) Volume() int {
	x_diff := cuboid.x_max - cuboid.x_min
	y_diff := cuboid.y_max - cuboid.y_min
	z_diff := cuboid.z_max - cuboid.z_min

	return x_diff * y_diff * z_diff
}

func (cuboids Cuboids) Volume() int {
	volume := 0
	for _, cuboid := range cuboids {
		volume += cuboid.Volume()
	}
	return volume
}

func main() {
	start := time.Now()
	f, _ := os.Open("input.txt")
	defer f.Close()
	scanner := bufio.NewScanner(f)

	cuboids, init_cube_count := make(Cuboids, 0), 0

	for scanner.Scan() {
		step_str := strings.Split(scanner.Text(), " ")
		var x_min, x_max, y_min, y_max, z_min, z_max int
		fmt.Sscanf(step_str[1], "x=%d..%d,y=%d..%d,z=%d..%d", &x_min, &x_max, &y_min, &y_max, &z_min, &z_max)
		min := Point{x_min, y_min, z_min}
		max := Point{x_max + 1, y_max + 1, z_max + 1}
		cuboid := CreateCuboid(min, max)

		if x_max > 50 && init_cube_count == 0 {
			init_cube_count = cuboids.Volume()
		}

		if step_str[0] == "on" {
			cuboids = cuboids.Add(cuboid)
		} else {
			cuboids = cuboids.Sub(cuboid)
		}
	}

	total_cube_count := cuboids.Volume()
	time_elapsed := time.Since(start)

	log.Printf(`
After executing the reboot steps, %d cubes are on in the initialization region.
After executing the reboot steps, %d cubes are on total.
Solution generated in %s.`,
		init_cube_count,
		total_cube_count,
		time_elapsed,
	)
}
