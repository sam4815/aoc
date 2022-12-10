package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func rangeToInt(str string) (int, int, int) {
	bounds := strings.Split(str, "-")

	left, leftErr := strconv.Atoi(bounds[0])
	if leftErr != nil {
		log.Fatal(leftErr)
	}
	right, rightErr := strconv.Atoi(bounds[1])
	if rightErr != nil {
		log.Fatal(rightErr)
	}

	return left, right, right - left
}

func main() {
	start := time.Now()

	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)

	num_fully_overlapping_pairs := 0
	num_partially_overlapping_pairs := 0

	for scanner.Scan() {
		ranges := strings.Split(scanner.Text(), ",")
		range1Left, range1Right, range1Size := rangeToInt(ranges[0])
		range2Left, range2Right, range2Size := rangeToInt(ranges[1])

		if range1Size > range2Size {
			if range1Left <= range2Left && range1Right >= range2Right {
				num_fully_overlapping_pairs++
			}
		} else if range2Left <= range1Left && range2Right >= range1Right {
			num_fully_overlapping_pairs++
		}

		if range1Left <= range2Left && range1Right >= range2Left {
			num_partially_overlapping_pairs++
		} else if range2Left <= range1Left && range2Right >= range1Left {
			num_partially_overlapping_pairs++
		}
	}

	time_elapsed := time.Since(start)

	log.Printf(`
The number of completely overlapping pairs is %d.
The number of partially overlapping pairs is %d.
Solution generated in %s.`,
		num_fully_overlapping_pairs,
		num_partially_overlapping_pairs,
		time_elapsed,
	)
}
