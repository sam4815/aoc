package main

import (
	"bufio"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

type Observer interface {
	Update(*CPU)
}

type Screen struct {
	pixels [][]string
}

type CPU struct {
	cycle           int
	x_register      int
	signal_strength int
	observers       []Observer
}

func (screen *Screen) Update(cpu *CPU) {
	row := (cpu.cycle - 1) / 40
	column := (cpu.cycle - 1) % 40
	is_pixel_lit := math.Abs(float64(column-cpu.x_register)) <= 1.0

	if is_pixel_lit {
		screen.pixels[row][column] = "#"
	} else {
		screen.pixels[row][column] = "."
	}
}

func (cpu *CPU) Update() {
	if (cpu.cycle+20)%40 == 0 {
		cpu.signal_strength += cpu.cycle * cpu.x_register
	}
}

func (cpu *CPU) Attach(o Observer) {
	cpu.observers = append(cpu.observers, o)
}

func (cpu *CPU) IncrementCycle() {
	cpu.cycle++

	cpu.Update()
	for _, o := range cpu.observers {
		o.Update(cpu)
	}
}

func (cpu *CPU) UpdateX(x int) {
	cpu.x_register = cpu.x_register + x
}

func main() {
	start := time.Now()

	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)

	pixels := make([][]string, 6)
	for i := range pixels {
		pixels[i] = make([]string, 40)
	}

	screen := Screen{pixels: pixels}
	cpu := CPU{x_register: 1, cycle: 0, signal_strength: 0}

	cpu.Attach(&screen)

	for scanner.Scan() {
		instruction := strings.Split(scanner.Text(), " ")

		switch instruction[0] {
		case "noop":
			cpu.IncrementCycle()
			continue
		case "addx":
			cpu.IncrementCycle()
			cpu.IncrementCycle()

			x_operand, _ := strconv.Atoi(instruction[1])
			cpu.UpdateX(x_operand)
		}
	}

	time_elapsed := time.Since(start)

	log.Printf("The screen displays:")
	for i := range screen.pixels {
		log.Printf("%v", screen.pixels[i])
	}

	log.Printf(`
The sum of the signal strengths is %d.
Solution generated in %s.`,
		cpu.signal_strength,
		time_elapsed,
	)
}
