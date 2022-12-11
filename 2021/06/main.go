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

	scanner.Scan()
	fish_days_str := strings.Split(scanner.Text(), ",")
	fish_days := make([]int, 9)

	for _, day_str := range fish_days_str {
		day, _ := strconv.Atoi(day_str)
		fish_days[day]++
	}

	for i := 0; i < 256; i++ {
		new_fish_days := make([]int, 9)
		for j := 0; j < 8; j++ {
			new_fish_days[j] = fish_days[j+1]
		}
		new_fish_days[8] += fish_days[0]
		new_fish_days[6] += fish_days[0]

		fish_days = new_fish_days
	}

	num_lanternfish := 0
	for i := 0; i < len(fish_days); i++ {
		num_lanternfish += fish_days[i]
	}

	time_elapsed := time.Since(start)

	log.Printf(`
The number of lanternfish after 256 days is %d.
Solution generated in %s.`,
		num_lanternfish,
		time_elapsed,
	)
}
