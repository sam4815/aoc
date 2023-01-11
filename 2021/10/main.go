package main

import (
	"bufio"
	"log"
	"os"
	"sort"
	"time"
)

func isClosingBracket(bracket string) bool {
	return bracket == "]" || bracket == ">" || bracket == "}" || bracket == ")"
}

func mapOpeningToClosing(bracket string) string {
	return map[string]string{
		"[": "]",
		"{": "}",
		"<": ">",
		"(": ")",
	}[bracket]
}

func mapClosingToOpening(bracket string) string {
	return map[string]string{
		"]": "[",
		"}": "{",
		">": "<",
		")": "(",
	}[bracket]
}

func mapClosingToSyntaxScore(bracket string) int {
	return map[string]int{
		"]": 57,
		"}": 1197,
		">": 25137,
		")": 3,
	}[bracket]
}

func mapClosingToAutocompleteScore(bracket string) int {
	return map[string]int{
		")": 1,
		"]": 2,
		"}": 3,
		">": 4,
	}[bracket]
}

func isValidChunk(str string, index int, required []string) (bool, int, []string) {
	opening := string(str[index])

	for i := index + 1; i < len(str); i++ {
		bracket := string(str[i])

		if isClosingBracket(bracket) {
			return mapClosingToOpening(bracket) == opening, i, required
		} else {
			is_valid, next_index, next_required := isValidChunk(str, i, required)
			if is_valid {
				i = next_index
				required = append(required, next_required...)
			} else {
				return is_valid, next_index, required
			}
		}
	}

	return true, len(str) - 1, append(required, mapOpeningToClosing(opening))
}

func main() {
	start := time.Now()

	f, _ := os.Open("input.txt")
	defer f.Close()

	scanner := bufio.NewScanner(f)
	syntax_error_score, autocomplete_scores := 0, make([]int, 0)

	for scanner.Scan() {
		brackets := scanner.Text()
		required := make([]string, 0)

		is_valid, idx, required := isValidChunk(brackets, 0, required)
		for is_valid && idx != len(brackets)-1 {
			is_valid, idx, required = isValidChunk(brackets, idx+1, required)
		}

		if !is_valid {
			syntax_error_score += mapClosingToSyntaxScore(string(brackets[idx]))
		} else {
			autocomplete_score := 0
			for _, bracket := range required {
				autocomplete_score *= 5
				autocomplete_score += mapClosingToAutocompleteScore(bracket)
			}

			autocomplete_scores = append(autocomplete_scores, autocomplete_score)
		}
	}

	sort.Ints(autocomplete_scores)
	middle_autocomplete_score := autocomplete_scores[len(autocomplete_scores)/2]

	time_elapsed := time.Since(start)

	log.Printf(`
The total syntax error score is %d.
The middle autocomplete score is %d.
Solution generated in %s.`,
		syntax_error_score,
		middle_autocomplete_score,
		time_elapsed,
	)
}
