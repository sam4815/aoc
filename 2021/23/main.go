package main

import (
	"io/ioutil"
	"log"
	"math"
	"strings"
	"time"
)

type Hallway []string
type Room []string
type Position struct {
	x, y int
}

type Burrow struct {
	hallway Hallway
	room_a  Room
	room_b  Room
	room_c  Room
	room_d  Room
	energy  int
}

func AbsInt(n int) int {
	if n >= 0 {
		return n
	} else {
		return 0 - n
	}
}

func (burrow Burrow) Clone() Burrow {
	clone := Burrow{
		hallway: make([]string, len(burrow.hallway)),
		room_a:  make([]string, len(burrow.room_a)),
		room_b:  make([]string, len(burrow.room_b)),
		room_c:  make([]string, len(burrow.room_c)),
		room_d:  make([]string, len(burrow.room_d)),
		energy:  burrow.energy,
	}
	copy(clone.hallway, burrow.hallway)
	copy(clone.room_a, burrow.room_a)
	copy(clone.room_b, burrow.room_b)
	copy(clone.room_c, burrow.room_c)
	copy(clone.room_d, burrow.room_d)

	return clone
}

func RoomPositionFromAmphipod(amphipod string) int {
	switch amphipod {
	case "A":
		return 2
	case "B":
		return 4
	case "C":
		return 6
	case "D":
		return 8
	default:
		return -1
	}
}

func (burrow Burrow) RoomFromPosition(position int) Room {
	switch position {
	case 2:
		return burrow.room_a
	case 4:
		return burrow.room_b
	case 6:
		return burrow.room_c
	case 8:
		return burrow.room_d
	default:
		return burrow.room_d
	}
}

func MoveCost(amphipod string) int {
	switch amphipod {
	case "A":
		return 1
	case "B":
		return 10
	case "C":
		return 100
	case "D":
		return 1000
	default:
		return -1
	}
}

func (burrow Burrow) IsMovable(amphipod string, amphipod_index int) bool {
	target_index := RoomPositionFromAmphipod(amphipod)
	room := burrow.RoomFromPosition(target_index)

	for _, cell := range room {
		if cell != amphipod && cell != "." {
			return false
		}
	}

	if amphipod_index > target_index {
		amphipod_index, target_index = target_index, amphipod_index
	}
	for i := amphipod_index + 1; i < target_index; i++ {
		if burrow.hallway[i] != "." {
			return false
		}
	}

	return true
}

func (room Room) Organized(amphipod string) bool {
	for i := 0; i < len(room); i++ {
		if room[i] != amphipod {
			return false
		}
	}

	return true
}

func (burrow Burrow) Organized() bool {
	return burrow.room_a.Organized("A") &&
		burrow.room_b.Organized("B") &&
		burrow.room_c.Organized("C") &&
		burrow.room_d.Organized("D")
}

func (room Room) MovableAmphipod(amphipod string) map[Position]string {
	for _, cell := range room {
		if cell == amphipod {
			continue
		}
		for i, cell := range room {
			if cell != "." {
				position := Position{x: RoomPositionFromAmphipod(amphipod), y: i}
				return map[Position]string{position: cell}
			}
		}
	}
	return make(map[Position]string)
}

func (burrow Burrow) MovableAmphipods() map[Position]string {
	all_moves := make(map[Position]string)

	for pos, amp := range burrow.room_a.MovableAmphipod("A") {
		all_moves[pos] = amp
	}
	for pos, amp := range burrow.room_b.MovableAmphipod("B") {
		all_moves[pos] = amp
	}
	for pos, amp := range burrow.room_c.MovableAmphipod("C") {
		all_moves[pos] = amp
	}
	for pos, amp := range burrow.room_d.MovableAmphipod("D") {
		all_moves[pos] = amp
	}

	return all_moves
}

func (burrow Burrow) OpenHallwayPositions(room_index int) []int {
	open_positions := make([]int, 0)

	for i := room_index + 1; i < 11; i++ {
		if burrow.hallway[i] == "." {
			if i != 2 && i != 4 && i != 6 && i != 8 {
				open_positions = append(open_positions, i)
			}
		} else {
			break
		}
	}

	for i := room_index - 1; i >= 0; i-- {
		if burrow.hallway[i] == "." {
			if i != 2 && i != 4 && i != 6 && i != 8 {
				open_positions = append(open_positions, i)
			}
		} else {
			break
		}
	}

	return open_positions
}

func (burrow Burrow) Moves() []Burrow {
	paths := make([]Burrow, 0)

	for x := 0; x < 11; x++ {
		if burrow.hallway[x] == "." {
			continue
		}
		amphipod := burrow.hallway[x]

		if burrow.IsMovable(amphipod, x) {
			path := burrow.Clone()
			room_index := RoomPositionFromAmphipod(amphipod)
			target_room := path.RoomFromPosition(room_index)

			for y := len(target_room) - 1; y >= 0; y-- {
				if target_room[y] == "." {
					steps_required := AbsInt(x-room_index) + 1 + y
					path.energy += steps_required * MoveCost(amphipod)
					target_room[y] = amphipod
					break
				}
			}

			path.hallway[x] = "."
			paths = append(paths, path)
		}
	}

	if len(paths) > 0 {
		return paths
	}

	for position, amphipod := range burrow.MovableAmphipods() {
		for _, open_pos := range burrow.OpenHallwayPositions(position.x) {
			path := burrow.Clone()
			path.hallway[open_pos] = amphipod
			path.RoomFromPosition(position.x)[position.y] = "."

			steps_required := AbsInt(open_pos-position.x) + 1 + position.y
			path.energy += steps_required * MoveCost(amphipod)

			paths = append(paths, path)
		}
	}

	return paths
}

func (burrow Burrow) String() string {
	return strings.Join(burrow.hallway, "") +
		strings.Join(burrow.room_a, "") +
		strings.Join(burrow.room_b, "") +
		strings.Join(burrow.room_c, "") +
		strings.Join(burrow.room_d, "")
}

func (burrow Burrow) BestPath() int {
	queue, curr_burrow := []Burrow{burrow}, burrow
	best_path := math.MaxInt
	visited := make(map[string]int)

	for len(queue) > 0 {
		// log.Print(len(queue))
		curr_burrow, queue = queue[0], queue[1:]
		// log.Print(curr_burrow.String())

		if curr_burrow.Organized() {
			// log.Print("WHOA")
			if curr_burrow.energy < best_path {
				log.Print(curr_burrow)
				best_path = curr_burrow.energy
			}
			continue
		}

		str := curr_burrow.String()
		if visited[str] != 0 && visited[str] < curr_burrow.energy {
			continue
		} else {
			visited[str] = curr_burrow.energy
		}

		queue = append(queue, curr_burrow.Moves()...)
	}

	return best_path
}

func CreateBurrow(burrow_strs []string) Burrow {
	burrow := Burrow{
		hallway: make([]string, 11),
		room_a:  make([]string, len(burrow_strs)-3),
		room_b:  make([]string, len(burrow_strs)-3),
		room_c:  make([]string, len(burrow_strs)-3),
		room_d:  make([]string, len(burrow_strs)-3),
	}

	for i := 0; i < 11; i++ {
		burrow.hallway[i] = string(burrow_strs[1][i+1])
	}
	for i := 2; i < len(burrow_strs)-1; i++ {
		burrow.room_a[i-2] = string(burrow_strs[i][3])
		burrow.room_b[i-2] = string(burrow_strs[i][5])
		burrow.room_c[i-2] = string(burrow_strs[i][7])
		burrow.room_d[i-2] = string(burrow_strs[i][9])
	}

	return burrow
}

func main() {
	start := time.Now()
	f, _ := ioutil.ReadFile("input.txt")
	burrow_strs := strings.Split(strings.TrimSpace(string(f)), "\n")

	partial_burrow := CreateBurrow(burrow_strs)
	burrow_strs = append(burrow_strs, burrow_strs[3:]...)
	burrow_strs[3] = "  #D#C#B#A#"
	burrow_strs[4] = "  #D#B#A#C#"
	full_burrow := CreateBurrow(burrow_strs)

	partial_least_energy := partial_burrow.BestPath()
	full_least_energy := full_burrow.BestPath()

	time_elapsed := time.Since(start)

	log.Printf(`
The least energy required to organize the amphipods in the partial burrow is %d.
The least energy required to organize the amphipods in the full burrow is %d.
Solution generated in %s.`,
		partial_least_energy,
		full_least_energy,
		time_elapsed,
	)
}
