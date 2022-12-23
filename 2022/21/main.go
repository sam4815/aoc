package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type Equation struct {
	x_coefficient float64
	constant      float64
}

type Monkey struct {
	name          string
	operator      func(x, y Equation) Equation
	value         Equation
	x_monkey_name string
	y_monkey_name string
}

func add(x, y Equation) Equation {
	result := Equation{}
	result.x_coefficient = x.x_coefficient + y.x_coefficient
	result.constant = x.constant + y.constant

	return result
}

func sub(x, y Equation) Equation {
	result := Equation{}
	result.x_coefficient = x.x_coefficient - y.x_coefficient
	result.constant = x.constant - y.constant

	return result
}

func mul(x, y Equation) Equation {
	result := Equation{}

	if x.x_coefficient == 0 {
		result.x_coefficient = x.constant * y.x_coefficient
		result.constant = x.constant * y.constant
	} else {
		result.x_coefficient = y.constant * x.x_coefficient
		result.constant = y.constant * x.constant
	}

	return result
}

func div(x, y Equation) Equation {
	result := Equation{}

	if y.x_coefficient != 0 {
		log.Fatal("Cannot divide by linear equation")
	}

	result.x_coefficient = x.x_coefficient / y.constant
	result.constant = x.constant / y.constant

	return result
}

func equals(x, y Equation) Equation {
	lhs, rhs := x, y
	if y.x_coefficient != 0 {
		rhs, lhs = y, x
	}

	if lhs.constant != 0 {
		rhs = add(rhs, Equation{constant: -lhs.constant})
		lhs = add(lhs, Equation{constant: -lhs.constant})
	}

	if lhs.x_coefficient != 1 {
		rhs = mul(rhs, Equation{constant: 1 / lhs.x_coefficient})
		lhs = mul(lhs, Equation{constant: 1 / lhs.x_coefficient})
	}

	return rhs
}

func stringToOperator(str string) func(x, y Equation) Equation {
	switch str {
	case "+":
		return add
	case "-":
		return sub
	case "*":
		return mul
	case "/":
		return div
	}
	return add
}

func parseMonkey(monkey_string string) Monkey {
	monkey_parts := strings.Split(monkey_string, ": ")
	monkey_name, equation := monkey_parts[0], monkey_parts[1]

	monkey := Monkey{name: monkey_name}

	if len(equation) <= 5 {
		constant, _ := strconv.ParseFloat(equation, 64)
		monkey.value = Equation{constant: constant}
		return monkey
	}

	equation_parts := strings.Split(equation, " ")
	x_name, operand, y_name := equation_parts[0], equation_parts[1], equation_parts[2]

	monkey.operator = stringToOperator(operand)
	monkey.x_monkey_name = x_name
	monkey.y_monkey_name = y_name

	return monkey
}

func (monkey *Monkey) Value(monkey_map map[string]Monkey) Equation {
	if monkey.value != (Equation{}) {
		return monkey.value
	}

	x_monkey := monkey_map[monkey.x_monkey_name]
	y_monkey := monkey_map[monkey.y_monkey_name]

	return monkey.operator(x_monkey.Value(monkey_map), y_monkey.Value(monkey_map))
}

func main() {
	start := time.Now()

	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)

	monkeys := make(map[string]Monkey)

	for scanner.Scan() {
		monkey := parseMonkey(scanner.Text())
		monkeys[monkey.name] = monkey
	}

	root_monkey := monkeys["root"]
	root_monkey_value := root_monkey.Value(monkeys)

	if root, ok := monkeys["root"]; ok {
		root.operator = equals
		monkeys["root"] = root
	}
	if humn, ok := monkeys["humn"]; ok {
		humn.value = Equation{x_coefficient: 1}
		monkeys["humn"] = humn
	}

	modified_root_monkey := monkeys["root"]
	human_value := modified_root_monkey.Value(monkeys)

	time_elapsed := time.Since(start)

	log.Printf(`
The root monkey will yell %f.
The human has to yell %f.
Solution generated in %s.`,
		root_monkey_value.constant,
		human_value.constant,
		time_elapsed,
	)
}
