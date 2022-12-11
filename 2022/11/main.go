package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

type WorryOp func(int) int
type ThrowOp func(int) int

type Monkey struct {
	items         []int
	inspect_count int
	divisor       int
	worry_op      WorryOp
	throw_op      ThrowOp
}

func parseWorryOp(operation string) WorryOp {
	operation_string := strings.Split(operation, " = old ")[1]
	operation_slice := strings.Split(operation_string, " ")
	operator, operand_string := operation_slice[0], operation_slice[1]

	if operand_string == "old" {
		switch operator {
		case "+":
			return func(i int) int { return i + i }
		case "*":
			return func(i int) int { return i * i }
		}
	}

	operand, _ := strconv.Atoi(operand_string)

	switch operator {
	case "+":
		return func(i int) int { return i + operand }
	case "*":
		return func(i int) int { return i * operand }
	}

	return func(i int) int { return i }
}

func parseThrowOp(operation []string) (ThrowOp, int) {
	divisor_slice := strings.Split(operation[0], " ")
	divisor_string := divisor_slice[len(divisor_slice)-1]
	divisor, _ := strconv.Atoi(divisor_string)

	true_monkey_slice := strings.Split(operation[1], " ")
	true_monkey_string := true_monkey_slice[len(true_monkey_slice)-1]
	true_monkey, _ := strconv.Atoi(true_monkey_string)

	false_monkey_slice := strings.Split(operation[2], " ")
	false_monkey_string := false_monkey_slice[len(false_monkey_slice)-1]
	false_monkey, _ := strconv.Atoi(false_monkey_string)

	return func(i int) int {
		if i%divisor == 0 {
			return true_monkey
		} else {
			return false_monkey
		}
	}, divisor
}

func main() {
	start := time.Now()

	f, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Print(err)
	}

	monkey_strings := strings.Split(string(f), "\n\n")

	monkeys := make([]Monkey, 0)
	monkey := Monkey{}

	for _, monkey_string := range monkey_strings {
		monkey_description := strings.Split(monkey_string, "\n")

		starting_items := make([]int, 0)
		starting_items_string := strings.Trim(monkey_description[1], " ")
		starting_items_string = strings.Split(starting_items_string, ": ")[1]
		starting_items_strings := strings.Split(starting_items_string, ", ")
		for _, item_string := range starting_items_strings {
			starting_item, _ := strconv.Atoi(item_string)
			starting_items = append(starting_items, starting_item)
		}
		monkey.items = starting_items

		operation_string := strings.Trim(monkey_description[2], " ")
		worry_op := parseWorryOp(operation_string)
		monkey.worry_op = worry_op

		throw_op, divisor := parseThrowOp(monkey_description[3:])
		monkey.throw_op = throw_op
		monkey.divisor = divisor

		monkeys = append(monkeys, monkey)
	}

	lcm := 1
	for i := 0; i < len(monkeys); i++ {
		lcm *= monkeys[i].divisor
	}

	for round := 0; round < 10000; round++ {
		for i := 0; i < len(monkeys); i++ {
			for _, item := range monkeys[i].items {
				item = monkeys[i].worry_op(item)
				item = item % lcm
				new_monkey_idx := monkeys[i].throw_op(item)

				monkeys[new_monkey_idx].items = append(monkeys[new_monkey_idx].items, item)
				monkeys[i].items = make([]int, 0)

				monkeys[i].inspect_count++
			}
		}
	}

	sort.Slice(monkeys, func(p, q int) bool {
		return monkeys[p].inspect_count > monkeys[q].inspect_count
	})

	monkey_business := monkeys[0].inspect_count * monkeys[1].inspect_count

	time_elapsed := time.Since(start)

	log.Printf(`
The level of monkey business is %d.
Solution generated in %s.`,
		monkey_business,
		time_elapsed,
	)
}
