package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func DivideByBits(nums [][]string, pos int) ([][]string, [][]string, int) {
	num_ones, num_zeros, winner := 0, 0, 0
	ones, zeros := make([][]string, 0), make([][]string, 0)

	for i := 0; i < len(nums); i++ {
		if nums[i][pos] == "1" {
			num_ones++
			ones = append(ones, nums[i])
		} else {
			num_zeros++
			zeros = append(zeros, nums[i])
		}
	}

	if num_ones > num_zeros {
		winner = 1
	} else if num_zeros > num_ones {
		winner = 0
	} else {
		winner = -1
	}

	return ones, zeros, winner
}

func main() {
	start := time.Now()

	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)

	num_ones, num_zeros := make([]int, 12), make([]int, 12)
	nums := make([][]string, 0)

	for scanner.Scan() {
		number := strings.Split(scanner.Text(), "")
		nums = append(nums, number)

		for pos := range number {
			if number[pos] == "1" {
				num_ones[pos]++
			} else {
				num_zeros[pos]++
			}
		}
	}

	epsilon_slice := make([]string, 12)
	gamma_slice := make([]string, 12)

	for i := 0; i < len(num_ones); i++ {
		if num_ones[i] > num_zeros[i] {
			epsilon_slice[i] = "1"
			gamma_slice[i] = "0"
		} else {
			epsilon_slice[i] = "0"
			gamma_slice[i] = "1"
		}
	}

	epsilon_binary := strings.Join(epsilon_slice, "")
	gamma_binary := strings.Join(gamma_slice, "")

	epsilon, err := strconv.ParseInt(epsilon_binary, 2, 64)
	if err != nil {
		log.Fatal(err)
	}

	gamma, err := strconv.ParseInt(gamma_binary, 2, 64)
	if err != nil {
		log.Fatal(err)
	}

	epsilon_gamma_product := epsilon * gamma

	oxygen_generator_possibilities := nums
	oxygen_position := 0
	co2_scrubber_possibilities := nums
	co2_position := 0

	for len(oxygen_generator_possibilities) != 1 {
		ones, zeros, winner := DivideByBits(oxygen_generator_possibilities, oxygen_position)
		log.Print(oxygen_position, len(oxygen_generator_possibilities), winner)
		oxygen_position++

		if winner == 0 {
			oxygen_generator_possibilities = zeros
		} else {
			oxygen_generator_possibilities = ones
		}
	}

	for len(co2_scrubber_possibilities) != 1 {
		ones, zeros, winner := DivideByBits(co2_scrubber_possibilities, co2_position)
		co2_position++

		if winner == 0 {
			co2_scrubber_possibilities = ones
		} else {
			co2_scrubber_possibilities = zeros
		}
	}

	oxygen_generator_binary := strings.Join(oxygen_generator_possibilities[0], "")
	co2_scrubber_binary := strings.Join(co2_scrubber_possibilities[0], "")

	oxygen_generator, err := strconv.ParseInt(oxygen_generator_binary, 2, 64)
	if err != nil {
		log.Fatal(err)
	}

	co2_scrubber, err := strconv.ParseInt(co2_scrubber_binary, 2, 64)
	if err != nil {
		log.Fatal(err)
	}

	oxygen_co2_product := oxygen_generator * co2_scrubber

	time_elapsed := time.Since(start)

	log.Printf(`
Multiplying the gamma rate by the epsilon rate gives %d.
Multiplying the oxygen generator by the CO2 scrubber gives %d.
Solution generated in %s.`,
		epsilon_gamma_product,
		oxygen_co2_product,
		time_elapsed,
	)
}
