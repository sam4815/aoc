package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"time"
)

func findIndex(arr []int, target int) int {
	index := -1
	for i := 0; i < len(arr); i++ {
		if arr[i] == target {
			index = i
		}
	}

	return index
}

func mix(arr []int, ind []int) ([]int, []int) {
	result, indices := make([]int, len(arr)), make([]int, len(arr))
	arr_length := len(arr) - 1
	for i := 0; i < len(arr); i++ {
		result[i] = arr[i]
		indices[i] = ind[i]
	}

	for i := 0; i < len(arr); i++ {
		curr_index := findIndex(indices, i)
		curr_val := result[curr_index]

		target_index := (curr_index + curr_val + arr_length) % arr_length

		if target_index == 0 && curr_val != 0 {
			target_index = arr_length
		}

		for target_index < 0 {
			target_index += arr_length
		}

		indices = append(indices[:curr_index], indices[curr_index+1:]...)
		indices = append(indices[:target_index], append([]int{i}, indices[target_index:]...)...)

		result = append(result[:curr_index], result[curr_index+1:]...)
		result = append(result[:target_index], append([]int{curr_val}, result[target_index:]...)...)
	}

	return result, indices
}

func findGroveSum(arr []int) int {
	zero_index := findIndex(arr, 0)
	arr_length := len(arr)
	grove_coordinates := []int{(1000 + zero_index) % arr_length, (2000 + zero_index) % arr_length, (3000 + zero_index) % arr_length}

	return arr[grove_coordinates[0]] + arr[grove_coordinates[1]] + arr[grove_coordinates[2]]
}

func main() {
	start := time.Now()

	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)

	encrypted_file, indices := make([]int, 0), make([]int, 0)

	for scanner.Scan() {
		number, _ := strconv.Atoi(scanner.Text())
		encrypted_file = append(encrypted_file, number)
		indices = append(indices, len(indices))
	}

	mixed, _ := mix(encrypted_file, indices)
	partially_mixed_grove_sum := findGroveSum(mixed)

	for i := range encrypted_file {
		encrypted_file[i] = encrypted_file[i] * 811589153
	}

	for i := 0; i < 10; i++ {
		encrypted_file, indices = mix(encrypted_file, indices)
	}

	fully_mixed_grove_sum := findGroveSum(encrypted_file)

	time_elapsed := time.Since(start)

	log.Printf(`
The sum of the grove coordinates after one mix is %d.
The sum of the grove coordinates is %d.
Solution generated in %s.`,
		partially_mixed_grove_sum,
		fully_mixed_grove_sum,
		time_elapsed,
	)
}
