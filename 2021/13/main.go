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

type Paper [][]string

func (paper *Paper) Init() {
	size := len(*paper)
	for i := 0; i < size; i++ {
		(*paper)[i] = make([]string, size)
		for j := 0; j < size; j++ {
			(*paper)[i][j] = "."
		}
	}
}

func (paper *Paper) Set(coordinates []int, val string) {
	(*paper)[coordinates[0]][coordinates[1]] = val
}

func (paper Paper) Get(coordinates []int) string {
	return paper[coordinates[0]][coordinates[1]]
}

func (paper *Paper) FoldX(x int) {
	for i := 0; i < len(*paper); i++ {
		for j := 0; j < x; j++ {
			distance := x - j
			if paper.Get([]int{i, x + distance}) == "#" {
				paper.Set([]int{i, j}, "#")
			}
		}
		(*paper)[i] = (*paper)[i][:x]
	}
}

func (paper *Paper) FoldY(y int) {
	for i := 0; i < y; i++ {
		distance := y - i
		if distance+y >= len(*paper) {
			continue
		}

		for j := 0; j < len((*paper)[i]); j++ {
			if paper.Get([]int{distance + y, j}) == "#" {
				paper.Set([]int{i, j}, "#")
			}
		}
	}
	*paper = (*paper)[:y]
}

func (paper Paper) CountDots() int {
	dot_count := 0
	for _, row := range paper {
		for _, val := range row {
			if val == "#" {
				dot_count += 1
			}
		}
	}
	return dot_count
}

func (paper Paper) Stringify() string {
	str := ""
	for _, row := range paper {
		str += strings.Join(row, "")
		str += "\n"
	}
	return str
}

func main() {
	start := time.Now()

	f, _ := os.Open("input.txt")
	defer f.Close()
	scanner := bufio.NewScanner(f)
	paper := Paper(make([][]string, 1500))
	paper.Init()

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), ",")
		if len(line) <= 1 {
			break
		}
		x, _ := strconv.Atoi(line[0])
		y, _ := strconv.Atoi(line[1])
		paper.Set([]int{y, x}, "#")
	}

	first_fold_dot_count := 0

	for scanner.Scan() {
		var direction byte
		var coordinate int
		fmt.Sscanf(scanner.Text(), "fold along %c=%d", &direction, &coordinate)
		if string(direction) == "x" {
			paper.FoldX(coordinate)
		} else {
			paper.FoldY(coordinate)
		}

		if first_fold_dot_count == 0 {
			first_fold_dot_count = paper.CountDots()
		}
	}

	time_elapsed := time.Since(start)

	log.Printf(`
After the first fold, there are %d dots visible.
After all folds, the paper looks like this:
%s
Solution generated in %s.`,
		first_fold_dot_count,
		paper.Stringify(),
		time_elapsed,
	)
}
