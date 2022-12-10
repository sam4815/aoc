package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func markBoards(boards [][][]int, char int) [][][]int {
	for i := 0; i < len(boards); i++ {
		for j := 0; j < len(boards[i]); j++ {
			for k := 0; k < len(boards[i][j]); k++ {
				if boards[i][j][k] == char {
					boards[i][j][k] = -1
				}
			}
		}
	}

	return boards
}

func returnCompleteBoardIndex(boards [][][]int) []int {
	completed := make([]int, 0)

	for i := 0; i < len(boards); i++ {
		for j := 0; j < len(boards[i]); j++ {
			row_complete := true
			column_complete := true

			for k := 0; k < len(boards[i][j]); k++ {
				if boards[i][j][k] != -1 {
					row_complete = false
				}
				if boards[i][k][j] != -1 {
					column_complete = false
				}
			}

			if row_complete || column_complete {
				completed = append(completed, i)
			}
		}
	}

	return completed
}

func sumBoard(board [][]int) int {
	sum := 0

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] != -1 {
				sum += board[i][j]
			}
		}
	}

	return sum
}

func emptyBoard() [][]int {
	board := make([][]int, 5)
	for i := range board {
		board[i] = make([]int, 5)
	}

	return board
}

func main() {
	start := time.Now()

	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)

	drawn_numbers := make([]int, 0)
	boards := make([][][]int, 0)
	board := make([][]int, 0)
	first_score, last_score := 0, 0

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			continue
		}

		if len(drawn_numbers) == 0 {
			drawn_numbers_strings := strings.Split(line, ",")
			for _, str := range drawn_numbers_strings {
				num, _ := strconv.Atoi(str)
				drawn_numbers = append(drawn_numbers, num)
			}
			continue
		}

		row := make([]int, 5)
		for i := 0; i < 5; i++ {
			chars := line[(i * 3):(i*3 + 2)]
			num, _ := strconv.Atoi(strings.Trim(string(chars), " "))
			row[i] = num
		}

		board = append(board, row)

		if len(board) == 5 {
			boards = append(boards, board)
			board = make([][]int, 0)
		}
	}

	for i := 0; i < len(drawn_numbers); i++ {
		boards = markBoards(boards, drawn_numbers[i])
		completed_board_indices := returnCompleteBoardIndex(boards)

		for _, idx := range completed_board_indices {
			score := drawn_numbers[i] * sumBoard(boards[idx])
			boards[idx] = emptyBoard()

			if first_score == 0 {
				first_score = score
			} else {
				last_score = score
			}
		}
	}

	time_elapsed := time.Since(start)

	log.Printf(`
The score of the winning board is %d.
The score of the last board to win is %d.
Solution generated in %s.`,
		first_score,
		last_score,
		time_elapsed,
	)
}
