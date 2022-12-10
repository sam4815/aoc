package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
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

	num_single_measurement_increases := -1
	num_window_measurement_increases := -3
	prev, prev_window, window_sum := 0, 0, 0
	window_a, window_b, window_c := 0, 0, 0

	for scanner.Scan() {
		measurement, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}

		window_sum = window_sum + measurement - window_a

		if measurement > prev {
			num_single_measurement_increases += 1
		}
		if window_sum > prev_window {
			num_window_measurement_increases += 1
		}

		window_a = window_b
		window_b = window_c
		window_c = measurement
		prev = measurement
		prev_window = window_sum
	}

	time_elapsed := time.Since(start)

	log.Printf(`
The number of measurements larger than the previous is %d.
The number of window measurements larger than the previous is %d.
Solution generated in %s.`,
		num_single_measurement_increases,
		num_window_measurement_increases,
		time_elapsed,
	)
}
