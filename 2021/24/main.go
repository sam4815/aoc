package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
	"time"
)

type ALU struct {
	w, x, y, z int
	input      string
	invalid    bool
	ReadInput  func() string
}

type ModelNumberGenerator struct {
	Generate func() string
}

type ALUGenerator struct {
	Generate func() ALU
}

type Solution struct {
	z            int
	model_number string
}

func (alu *ALU) Init(input string) {
	alu.input = input
	n := -1
	alu.ReadInput = func() string {
		n = n + 1
		return string(alu.input[n])
	}
}

func (alu ALU) Value(variable string) int {
	switch variable {
	case "w":
		return alu.w
	case "x":
		return alu.x
	case "y":
		return alu.y
	case "z":
		return alu.z
	default:
		num, _ := strconv.Atoi(variable)
		return num
	}
}

func (alu *ALU) Store(register string, value int) {
	switch register {
	case "w":
		alu.w = value
	case "x":
		alu.x = value
	case "y":
		alu.y = value
	case "z":
		alu.z = value
	default:
		log.Fatal("Register not supported")
	}
}

func (alu *ALU) Perform(instruction string) {
	split_instruction := strings.Split(instruction, " ")
	op, register, variable := split_instruction[0], split_instruction[1], ""

	if op == "inp" {
		variable = alu.ReadInput()
	} else {
		variable = split_instruction[2]
	}

	a, b := alu.Value(register), alu.Value(variable)

	switch op {
	case "inp":
		alu.Store(register, b)
	case "add":
		alu.Store(register, a+b)
	case "mul":
		alu.Store(register, a*b)
	case "div":
		alu.Store(register, a/b)
	case "mod":
		alu.Store(register, a%b)
	case "eql":
		if a == b {
			alu.Store(register, 1)
		} else {
			alu.Store(register, 0)
		}
	}
}

func GreaterThan(a, b string) bool {
	a_num, _ := strconv.Atoi(a)
	b_num, _ := strconv.Atoi(b)

	return a_num > b_num
}

func LessThan(a, b string) bool {
	a_num, _ := strconv.Atoi(a)
	b_num, _ := strconv.Atoi(b)

	return a_num < b_num
}

func ForwardSolutions(instructions []string, input_solutions []Solution, compare func(string, string) bool) []Solution {
	output_solutions_map := make(map[int]Solution, 0)
	output_solutions := make([]Solution, 0)

	for _, input_solution := range input_solutions {
		for i := 9; i > 0; i-- {
			alu := ALU{z: input_solution.z}
			input := fmt.Sprint(i)
			alu.Init(input)

			for _, instruction := range instructions {
				alu.Perform(instruction)
			}

			if existing_solution, exists := output_solutions_map[alu.z]; exists {
				if compare(input_solution.model_number+input, existing_solution.model_number) {
					output_solutions_map[alu.z] = Solution{z: alu.z, model_number: input_solution.model_number + input}
				}
			} else {
				output_solutions_map[alu.z] = Solution{z: alu.z, model_number: input_solution.model_number + input}
			}
		}
	}

	for _, output_solution := range output_solutions_map {
		if output_solution.z < 10000000 {
			output_solutions = append(output_solutions, output_solution)
		}
	}

	return output_solutions
}

func main() {
	start := time.Now()
	f, _ := ioutil.ReadFile("input.txt")

	instructions := strings.Split(strings.TrimSpace(string(f)), "\n")
	instruction_sections := make([][]string, 0)
	for i := 0; i < len(instructions); i += 18 {
		instruction_sections = append(instruction_sections, instructions[i:i+18])
	}

	inputs := []Solution{{z: 0, model_number: ""}}

	for i := 0; i < len(instruction_sections); i++ {
		solutions := ForwardSolutions(instruction_sections[i], inputs, GreaterThan)
		inputs = make([]Solution, len(solutions))
		copy(inputs, solutions)
	}

	largest_solution := ""
	for _, input := range inputs {
		if input.z == 0 {
			largest_solution = input.model_number
			break
		}
	}

	inputs = []Solution{{z: 0, model_number: ""}}
	for i := 0; i < len(instruction_sections); i++ {
		solutions := ForwardSolutions(instruction_sections[i], inputs, LessThan)
		inputs = make([]Solution, len(solutions))
		copy(inputs, solutions)
	}

	smallest_solution := ""
	for _, input := range inputs {
		if input.z == 0 {
			smallest_solution = input.model_number
			break
		}
	}

	time_elapsed := time.Since(start)

	log.Printf(`
The largest possible model number is %s.
The smallest possible model number is %s.
Solution generated in %s.`,
		largest_solution,
		smallest_solution,
		time_elapsed,
	)
}
