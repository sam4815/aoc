package main

import (
	"bufio"
	"log"
	"math"
	"os"
	"strings"
	"time"
)

type SequenceOccurrences map[string]int
type InsertionRules map[string][]string

func (occurrences SequenceOccurrences) Insert(rules InsertionRules) SequenceOccurrences {
	next_occurrences := make(map[string]int)
	for pair, quantity := range occurrences {
		if len(pair) == 1 {
			next_occurrences[pair] += quantity
			continue
		}

		for _, resulting_pair := range rules[pair] {
			next_occurrences[resulting_pair] += quantity
		}
		next_occurrences[string(rules[pair][0][1])] += quantity
	}
	return next_occurrences
}

func (occurrences *SequenceOccurrences) AddTemplate(template string) {
	for i := 0; i < len(template); i++ {
		(*occurrences)[string(template[i])] += 1
		if i+1 >= len(template) {
			continue
		}

		pair := string(template[i]) + string(template[i+1])
		(*occurrences)[pair] += 1
	}
}

func (occurrences SequenceOccurrences) CalculateMaximumDifference() int {
	min, max := math.MaxInt, math.MinInt
	for sequence, quantity := range occurrences {
		if len(sequence) != 1 {
			continue
		}

		if quantity > max {
			max = quantity
		}
		if quantity < min {
			min = quantity
		}
	}
	return max - min
}

func main() {
	start := time.Now()

	f, _ := os.Open("input.txt")
	defer f.Close()
	scanner := bufio.NewScanner(f)

	quantities := make(SequenceOccurrences)
	insertion_rules := make(InsertionRules)

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		} else if len(quantities) == 0 {
			quantities.AddTemplate(line)
			continue
		}

		insertion_rule := strings.Split(line, " -> ")
		pair, char := strings.Split(insertion_rule[0], ""), insertion_rule[1]
		pair_start, pair_end := pair[0], pair[1]
		insertion_rules[insertion_rule[0]] = []string{pair_start + char, char + pair_end}
	}

	max_difference_after_10 := 0

	for i := 0; i < 40; i++ {
		if i == 10 {
			max_difference_after_10 = quantities.CalculateMaximumDifference()
		}
		quantities = quantities.Insert(insertion_rules)
	}

	max_difference_after_40 := quantities.CalculateMaximumDifference()
	time_elapsed := time.Since(start)

	log.Printf(`
The difference between the most and least common element after 10 steps is %d.
The difference between the most and least common element after 40 steps is %d.
Solution generated in %s.
		`,
		max_difference_after_10,
		max_difference_after_40,
		time_elapsed,
	)
}
