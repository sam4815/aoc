package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	start := time.Now()

	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)

	simple_horizontal_position, simple_depth := 0, 0
	horizontal_position, depth, aim := 0, 0, 0

	for scanner.Scan() {
		command := strings.Split(scanner.Text(), " ")

		direction := command[0]
		magnitude, err := strconv.Atoi(command[1])
		if err != nil {
			log.Fatal(err)
		}

		switch direction {
		case "forward":
			simple_horizontal_position += magnitude
			horizontal_position += magnitude
			depth += (magnitude * aim)
		case "down":
			simple_depth += magnitude
			aim += magnitude
		case "up":
			simple_depth -= magnitude
			aim -= magnitude
		}
	}

	simple_product := simple_horizontal_position * simple_depth
	product := horizontal_position * depth

	time_elapsed := time.Since(start)

	log.Printf(`
Multiplying the submarine's depth and horizontal position following the simple course gives %d.
Multiplying the submarine's depth and horizontal position following the proper course gives %d.
Solution generated in %s.`,
		simple_product,
		product,
		time_elapsed,
	)
}
