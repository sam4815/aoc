package main

import (
	"bufio"
	"log"
	"os"
	"time"
)

func main() {
	start := time.Now()

	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)

	packet_marker_index, message_marker_index := 0, 0
	letter_map := map[byte]bool{}
	found_packet := false

	for scanner.Scan() {
		buffer := scanner.Text()
		for pos := range buffer {
			for i := pos; i >= pos-4; i-- {
				if i < 0 {
					continue
				}

				letter_map[buffer[i]] = true

				if i == pos-3 && !found_packet && len(letter_map) == 4 {
					found_packet = true
					packet_marker_index = pos + 1
				}
			}

			if len(letter_map) == 14 {
				message_marker_index = pos + 1
				break
			}

			letter_map = map[byte]bool{}
		}
	}

	time_elapsed := time.Since(start)

	log.Printf(`
The first start-of-packet marker ends after character %d.
The first start-of-message marker ends after character %d.
Solution generated in %s.`,
		packet_marker_index,
		message_marker_index,
		time_elapsed,
	)
}
