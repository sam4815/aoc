package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func insertAt(arr []int, index int, val int) {
	for i := len(arr) - 1; i > index; i-- {
		arr[i] = arr[i-1]
	}
	arr[index] = val
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)

	top_three_calorie_counts := make([]int, 3)
	current_elf_calories := 0

	for scanner.Scan() {
		if scanner.Text() == "" {
			for x := 0; x < 3; x++ {
				if current_elf_calories > top_three_calorie_counts[x] {
					insertAt(top_three_calorie_counts, x, current_elf_calories)
					break
				}
			}

			current_elf_calories = 0
		} else {
			calories, err := strconv.Atoi(scanner.Text())
			if err != nil {
				log.Fatal(err)
			}

			current_elf_calories += calories
		}
	}

	top_calorie_count := top_three_calorie_counts[0]
	top_three_calorie_counts_sum := top_three_calorie_counts[0] + top_three_calorie_counts[1] + top_three_calorie_counts[2]

	log.Printf(`
The elf carrying the most calories is carrying %d calories.
The elves carrying the three largest number of calories are carrying a total of %d calories.`,
		top_calorie_count,
		top_three_calorie_counts_sum,
	)
}
