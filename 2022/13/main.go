package main

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

func appendStringifiedNumber(str *[]string, dst *[]interface{}) {
	if len(*str) == 0 {
		return
	}

	curr_num_str := strings.Join(*str, "")
	curr_num, _ := strconv.Atoi(curr_num_str)
	*dst = append(*dst, curr_num)
	*str = make([]string, 0)
}

func parse(value string) ([]interface{}, int) {
	OPEN_SLICE := 91
	CLOSE_SLICE := 93
	COMMA := 44

	result := make([]interface{}, 0)
	curr_item := make([]string, 0)
	i := 1

	for i < len(value) {
		if int(value[i]) == CLOSE_SLICE {
			appendStringifiedNumber(&curr_item, &result)
			break
		} else if int(value[i]) == COMMA {
			appendStringifiedNumber(&curr_item, &result)
			i += 1
		} else if int(value[i]) == OPEN_SLICE {
			item, idx := parse(value[i:])
			result = append(result, item)
			i += idx + 1
		} else {
			curr_item = append(curr_item, value[i:i+1])
			i += 1
		}
	}

	return result, i
}

func isNumber(value interface{}) bool {
	_, ok := value.(int)
	return ok
}

func isSlice(value interface{}) bool {
	_, ok := value.([]interface{})
	return ok
}

func compare(a interface{}, b interface{}) int {
	if isNumber(a) && isNumber(b) {
		if a.(int) < b.(int) {
			return 1
		} else if a.(int) > b.(int) {
			return -1
		} else {
			return 0
		}
	}

	if isSlice(a) && isSlice(b) {
		smaller_len := len(a.([]interface{}))
		if len(b.([]interface{})) < smaller_len {
			smaller_len = len(b.([]interface{}))
		}

		for i := 0; i < smaller_len; i++ {
			comparison := compare(a.([]interface{})[i], b.([]interface{})[i])
			if comparison == 1 {
				return 1
			} else if comparison == -1 {
				return -1
			}
		}

		if len(a.([]interface{})) < len(b.([]interface{})) {
			return 1
		} else if len(a.([]interface{})) > len(b.([]interface{})) {
			return -1
		}

		return 0
	}

	if isNumber(a) && isSlice(b) {
		return compare([]interface{}{a}, b)
	}

	if isSlice(a) && isNumber(b) {
		return compare(a, []interface{}{b})
	}

	return 0
}

func main() {
	start := time.Now()

	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)

	divider_one, _ := parse("[[2]]")
	divider_two, _ := parse("[[6]]")
	packets := [][]interface{}{divider_one, divider_two}

	pairs := make([][][]interface{}, 0)
	pair := make([][]interface{}, 0)

	for scanner.Scan() {
		if scanner.Text() == "" {
			continue
		}

		value, _ := parse(scanner.Text())

		pair = append(pair, value)
		packets = append(packets, value)

		if len(pair) == 2 {
			pairs = append(pairs, pair)
			pair = make([][]interface{}, 0)
		}
	}

	indices_sum := 0
	for i, pair := range pairs {
		if compare(pair[0], pair[1]) == 1 {
			indices_sum += i + 1
		}
	}

	sort.Slice(packets, func(p, q int) bool {
		return compare(packets[p], packets[q]) > 0
	})

	divider_one_index, divider_two_index := 0, 0

	for i := range packets {
		if compare(divider_one, packets[i]) == 0 {
			divider_one_index = i + 1
		} else if compare(divider_two, packets[i]) == 0 {
			divider_two_index = i + 1
		}
	}

	decoder_key := divider_one_index * divider_two_index

	time_elapsed := time.Since(start)

	log.Printf(`
The sum of the indices of the pairs in the right order is %d.
The decoder key is %d.
Solution generated in %s.`,
		indices_sum,
		decoder_key,
		time_elapsed,
	)
}
