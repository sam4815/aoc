package main

import (
	"bufio"
	"log"
	"math"
	"os"
	"time"
)

type Snafu string
type Number int

func snafuCharToInt(char string) int {
	switch char {
	case "2":
		return 2
	case "1":
		return 1
	case "0":
		return 0
	case "-":
		return -1
	case "=":
		return -2
	default:
		log.Fatal("Unsupported char")
		return 0
	}
}

func intToSnafuChar(i int) string {
	switch i {
	case 2:
		return "2"
	case 1:
		return "1"
	case 0:
		return "0"
	case -1:
		return "-"
	case -2:
		return "="
	default:
		log.Fatal("Unsupported int")
		return ""
	}
}

func powInt(x, y int) int {
	return int(math.Pow(float64(x), float64(y)))
}

func (snafu Snafu) Decode() int {
	length := len(snafu)
	sum := 0

	for i, char := range snafu {
		order := length - i - 1
		sum += powInt(5, order) * snafuCharToInt(string(char))
	}

	return sum
}

func (num Number) Encode() string {
	length := 20
	five_powers := make([]int, length)

	for i := range five_powers {
		order := length - i - 1
		power := powInt(5, order)
		for int(num) >= power {
			num -= Number(power)
			five_powers[i] += 1
		}
	}

	for i := len(five_powers) - 1; i >= 0; i-- {
		if five_powers[i] > 2 {
			five_powers[i-1] += 1
			five_powers[i] -= 5
		}
	}

	snafu := ""

	for i := range five_powers {
		if len(snafu) > 0 || five_powers[i] != 0 {
			snafu += intToSnafuChar(five_powers[i])
		}
	}

	return snafu
}

func main() {
	start := time.Now()

	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)

	snafus := make([]Snafu, 0)

	for scanner.Scan() {
		snafu := scanner.Text()
		snafus = append(snafus, Snafu(snafu))
	}

	snafu_sum := 0
	for _, snafu := range snafus {
		snafu_sum += snafu.Decode()
	}

	required_snafu := Number(snafu_sum).Encode()

	time_elapsed := time.Since(start)

	log.Printf(`
The required SNAFU number is %s.
Solution generated in %s.`,
		required_snafu,
		time_elapsed,
	)
}
