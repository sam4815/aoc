package main

import (
	"bufio"
	"log"
	"os"
)

func toPriority(i int) int {
	if i >= 97 {
		return i - 96
	} else if i >= 65 {
		return i - 38
	}
	return 0
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)

	total_priority_of_duplicates, total_priority_of_badges, current_group_index := 0, 0, 0
	encountered_in_rucksack := map[string]int{}

	for scanner.Scan() {
		line := scanner.Text()
		encountered_in_first_compartment := map[string]bool{}
		halfway := len(line) / 2

		for x := 0; x < len(line); x++ {
			item_type := string(line[x])

			if x < halfway {
				encountered_in_first_compartment[item_type] = true
			} else if encountered_in_first_compartment[item_type] {
				total_priority_of_duplicates += toPriority(int(line[x]))
				encountered_in_first_compartment = map[string]bool{}
			}

			if encountered_in_rucksack[item_type] == current_group_index {
				encountered_in_rucksack[item_type] = current_group_index + 1
			}

			if encountered_in_rucksack[item_type] == 3 {
				total_priority_of_badges += toPriority(int(line[x]))
				encountered_in_rucksack = map[string]int{}
			}
		}

		current_group_index = (current_group_index + 1) % 3
	}

	log.Printf(`
The sum of the duplicate item priorities is %d.
The sum of the badge priorities is %d.`,
		total_priority_of_duplicates,
		total_priority_of_badges,
	)
}
