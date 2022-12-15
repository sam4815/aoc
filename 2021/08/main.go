// 0:      1:      2:      3:      4:
//  aaaa    ....    aaaa    aaaa    ....
// b    c  .    c  .    c  .    c  b    c
// b    c  .    c  .    c  .    c  b    c
//  ....    ....    dddd    dddd    dddd
// e    f  .    f  e    .  .    f  .    f
// e    f  .    f  e    .  .    f  .    f
//  gggg    ....    gggg    gggg    ....

//  5:      6:      7:      8:      9:
//  aaaa    aaaa    aaaa    aaaa    aaaa
// b    .  b    .  .    c  b    c  b    c
// b    .  b    .  .    c  b    c  b    c
//  dddd    dddd    ....    dddd    dddd
// .    f  e    f  .    f  e    f  .    f
// .    f  e    f  .    f  e    f  .    f
//  gggg    gggg    ....    gggg    gggg

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

func subtract(a, b string) string {
	b_map := map[string]bool{}
	result := ""

	for _, char := range b {
		b_map[string(char)] = true
	}

	for _, char := range a {
		if b_map[string(char)] != true {
			result += string(char)
		}
	}

	return result
}

func has(a, b string) bool {
	for _, char := range a {
		if string(char) == b {
			return true
		}
	}

	return false
}

func equal(a, b string) bool {
	if len(a) != len(b) {
		return false
	}

	a_map := map[string]bool{}

	for _, char := range a {
		a_map[string(char)] = true
	}

	for _, char := range b {
		if a_map[string(char)] != true {
			return false
		}
	}

	return true
}

func main() {
	start := time.Now()

	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)

	num_unique_numbers, sum_output_numbers := 0, 0

	for scanner.Scan() {
		entry := strings.Split(scanner.Text(), " | ")

		numbers := entry[0]
		output := entry[1]

		number_map := map[int]string{}
		segment_map := map[string]string{}
		segment_occurrences := map[string]int{}

		// Simple cases where number of segments is unique
		for _, segments := range strings.Split(numbers, " ") {
			value_len := len(segments)

			if value_len == 2 {
				number_map[1] = segments
			}
			if value_len == 4 {
				number_map[4] = segments
			}
			if value_len == 3 {
				number_map[7] = segments
			}
			if value_len == 7 {
				number_map[8] = segments
			}
			for _, segment := range segments {
				segment_occurrences[string(segment)]++
			}
		}

		// Calculate segments using number of occurrences and subtraction
		segment_map["a"] = subtract(number_map[7], number_map[1])

		for segment, occurences := range segment_occurrences {
			if occurences == 9 {
				segment_map["f"] = segment
			}
			if occurences == 4 {
				segment_map["e"] = segment
			}
			if occurences == 6 {
				segment_map["b"] = segment
			}
			if occurences == 8 && segment != segment_map["a"] {
				segment_map["c"] = segment
			}
		}

		segment_map["d"] = subtract(subtract(number_map[4], number_map[1]), segment_map["b"])

		// Calcuate numbers by subtracting segments from found numbers
		number_map[9] = subtract(number_map[8], segment_map["e"])
		number_map[3] = subtract(number_map[9], segment_map["b"])
		number_map[6] = subtract(number_map[8], segment_map["c"])
		number_map[5] = subtract(number_map[6], segment_map["e"])
		number_map[0] = subtract(number_map[8], segment_map["d"])
		number_map[2] = subtract(subtract(number_map[8], segment_map["b"]), segment_map["f"])

		output_numbers := ""

		for _, value := range strings.Split(output, " ") {
			for number, segments := range number_map {
				if equal(value, segments) {
					output_numbers += fmt.Sprintf("%d", number)
				}
			}

			value_len := len(value)
			for _, num_segments := range []int{2, 4, 3, 7} {
				if value_len == num_segments {
					num_unique_numbers++
				}
			}
		}

		number, _ := strconv.Atoi(output_numbers)
		sum_output_numbers += number
	}

	time_elapsed := time.Since(start)

	log.Printf(`
The digits with a unique number of segments occur %d times.
The sum of all output numbers is %d.
Solution generated in %s.`,
		num_unique_numbers,
		sum_output_numbers,
		time_elapsed,
	)
}
