package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
	"time"
)

type Snailfish struct {
	parent *Snailfish
	left   *Snailfish
	right  *Snailfish
	value  int
}

func parse(value string) (*Snailfish, int) {
	OPEN_SLICE, CLOSE_SLICE, COMMA := 91, 93, 44
	fish, i := Snailfish{}, 1

	for i < len(value) {
		if int(value[i]) == CLOSE_SLICE {
			break
		} else if int(value[i]) == COMMA {
			i += 1
		} else if int(value[i]) == OPEN_SLICE {
			sub_fish, idx := parse(value[i:])
			sub_fish.parent = &fish
			i += idx + 1
			if fish.left == nil {
				fish.left = sub_fish
			} else {
				fish.right = sub_fish
			}
		} else {
			curr_num, _ := strconv.Atoi(string(value[i]))
			i += 1
			sub_fish := Snailfish{value: curr_num}
			sub_fish.parent = &fish
			if fish.left == nil {
				fish.left = &sub_fish
			} else {
				fish.right = &sub_fish
			}
		}
	}

	return &fish, i
}

func (snailfish_a *Snailfish) Add(snailfish_b *Snailfish) *Snailfish {
	fish := Snailfish{left: snailfish_a, right: snailfish_b}
	snailfish_a.parent, snailfish_b.parent = &fish, &fish
	return &fish
}

func (snailfish *Snailfish) Depth() int {
	depth := 0
	curr_fish := snailfish
	for curr_fish.parent != nil {
		depth += 1
		curr_fish = curr_fish.parent
	}
	return depth
}

func (snailfish *Snailfish) AddToLeftFish(val int) {
	var left_fish *Snailfish
	curr_fish := snailfish
	for curr_fish.parent != nil {
		if curr_fish.parent.left != curr_fish {
			left_fish = curr_fish.parent.left
			break
		} else {
			curr_fish = curr_fish.parent
		}
	}

	if left_fish != nil {
		for left_fish.right != nil {
			left_fish = left_fish.right
		}
		left_fish.value += val
	}
}

func (snailfish *Snailfish) AddToRightFish(val int) {
	var right_fish *Snailfish
	curr_fish := snailfish
	for curr_fish.parent != nil {
		if curr_fish.parent.right != curr_fish {
			right_fish = curr_fish.parent.right
			break
		} else {
			curr_fish = curr_fish.parent
		}
	}

	if right_fish != nil {
		for right_fish.left != nil {
			right_fish = right_fish.left
		}
		right_fish.value += val
	}
}

func (snailfish *Snailfish) Explode() {
	snailfish.AddToLeftFish(snailfish.left.value)
	snailfish.AddToRightFish(snailfish.right.value)
	snailfish.left, snailfish.right = nil, nil
	snailfish.value = 0
}

func (snailfish *Snailfish) TryExplode() bool {
	fish_queue, curr_fish := []*Snailfish{snailfish}, snailfish

	for len(fish_queue) > 0 {
		curr_fish, fish_queue = fish_queue[0], fish_queue[1:]
		if curr_fish.Depth() == 4 && curr_fish.left != nil {
			curr_fish.Explode()
			return true
		}

		if curr_fish.left != nil {
			fish_queue = append([]*Snailfish{curr_fish.left, curr_fish.right}, fish_queue...)
		}
	}

	return false
}

func (snailfish *Snailfish) Split() {
	left_fish := Snailfish{value: snailfish.value / 2, parent: snailfish}
	right_fish := Snailfish{value: snailfish.value / 2, parent: snailfish}
	if snailfish.value%2 != 0 {
		right_fish.value += 1
	}

	snailfish.left, snailfish.right = &left_fish, &right_fish
	snailfish.value = 0
}

func (snailfish *Snailfish) TrySplit() bool {
	fish_queue, curr_fish := []*Snailfish{snailfish}, snailfish

	for len(fish_queue) > 0 {
		curr_fish, fish_queue = fish_queue[0], fish_queue[1:]
		if curr_fish.value >= 10 && curr_fish.left == nil {
			curr_fish.Split()
			return true
		}

		if curr_fish.left != nil {
			fish_queue = append([]*Snailfish{curr_fish.left, curr_fish.right}, fish_queue...)
		}
	}

	return false
}

func (snailfish *Snailfish) Reduce() *Snailfish {
	for true {
		if snailfish.TryExplode() {
			continue
		}
		if !snailfish.TrySplit() {
			break
		}
	}

	return snailfish
}

func (snailfish Snailfish) Stringify() string {
	if snailfish.left == nil {
		if snailfish.value >= 10 && snailfish.Depth() <= 4 {
			return fmt.Sprintf("\x1b[31m%d\033[0m", snailfish.value)
		} else {
			return fmt.Sprintf("%d", snailfish.value)
		}
	} else {
		if snailfish.Depth() >= 4 {
			return fmt.Sprintf("\x1b[41m[%s, %s]\033[0m", snailfish.left.Stringify(), snailfish.right.Stringify())
		} else {
			return fmt.Sprintf("[%s, %s]", snailfish.left.Stringify(), snailfish.right.Stringify())
		}
	}
}

func (snailfish Snailfish) Clone() *Snailfish {
	if snailfish.left == nil {
		return &Snailfish{value: snailfish.value}
	}

	left_snailfish, right_snailfish := snailfish.left.Clone(), snailfish.right.Clone()
	clone_fish := &Snailfish{left: left_snailfish, right: right_snailfish}
	left_snailfish.parent, right_snailfish.parent = clone_fish, clone_fish

	return clone_fish
}

func (snailfish Snailfish) Magnitude() int {
	if snailfish.left == nil {
		return snailfish.value
	}

	return 3*snailfish.left.Magnitude() + 2*snailfish.right.Magnitude()
}

func main() {
	start := time.Now()
	f, _ := ioutil.ReadFile("input.txt")
	snailfish := make([]*Snailfish, 0)

	for _, fish_str := range strings.Split(string(f), "\n") {
		if len(fish_str) > 0 {
			fish, _ := parse(fish_str)
			snailfish = append(snailfish, fish)
		}
	}

	snailfish_sum := snailfish[0].Clone()
	for _, fish := range snailfish[1:] {
		snailfish_sum = snailfish_sum.Add(fish.Clone()).Reduce()
	}

	total_magnitude := snailfish_sum.Magnitude()

	greatest_magnitude := 0

	for i, fish_a := range snailfish {
		for _, fish_b := range snailfish[i+1:] {
			forward_sum := fish_a.Clone().Add(fish_b.Clone()).Reduce()
			backward_sum := fish_b.Clone().Add(fish_a.Clone()).Reduce()

			if forward_sum.Magnitude() > greatest_magnitude {
				greatest_magnitude = forward_sum.Magnitude()
			}
			if backward_sum.Magnitude() > greatest_magnitude {
				greatest_magnitude = backward_sum.Magnitude()
			}
		}
	}

	time_elapsed := time.Since(start)

	log.Printf(`
The magnitude of the sum of all snailfish is %d.
The greatest possible magnitude of the sum of two snailfish is %d.
Solution generated in %s`,
		total_magnitude,
		greatest_magnitude,
		time_elapsed,
	)
}
