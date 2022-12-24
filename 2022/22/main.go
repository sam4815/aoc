package main

import (
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type Facing int
type CubeFace int

const (
	Right Facing = 0
	Down         = 1
	Left         = 2
	Up           = 3
)

const (
	A CubeFace = 0
	B          = 1
	C          = 2
	D          = 3
	E          = 4
	F          = 5
)

type Player struct {
	position []int
	facing   Facing
}

type Board struct {
	width     int
	height    int
	tiles     [][]string
	faces     map[CubeFace][]int
	face_size int
	player    Player
}

func (player Player) GetFacingSymbol() string {
	switch player.facing {
	case Right:
		return ">"
	case Down:
		return "v"
	case Up:
		return "^"
	case Left:
		return "<"
	}
	return "?"
}

func (board *Board) RotatePlayer(direction string) {
	switch direction {
	case "L":
		board.player.facing = (board.player.facing - 1 + 4) % 4
		board.MarkPlayerOnBoard()
	case "R":
		board.player.facing = (board.player.facing + 1 + 4) % 4
		board.MarkPlayerOnBoard()
	}
}

func (board Board) GetPassword() int {
	return 1000*(board.player.position[0]+1) + 4*(board.player.position[1]+1) + int(board.player.facing)
}

func (board Board) GetNextPosition(position []int) []int {
	switch board.player.facing {
	case Right:
		return []int{position[0], position[1] + 1}
	case Down:
		return []int{position[0] + 1, position[1]}
	case Left:
		return []int{position[0], position[1] - 1}
	case Up:
		return []int{position[0] - 1, position[1]}
	}
	return position
}

func (board *Board) BuildTiles(map_str string) {
	board.tiles = make([][]string, 0)

	for _, row := range strings.Split(map_str, "\n") {
		tiles := make([]string, len(row))
		for j, tile := range strings.Split(row, "") {
			tiles[j] = tile
		}
		board.tiles = append(board.tiles, tiles)

		if len(tiles) > board.width {
			board.width = len(tiles)
		}
	}

	board.height = len(board.tiles)
	// Pad rows with empty space to keep row length consistent
	for index, row := range board.tiles {
		diff := board.width - len(row)
		if diff == 0 {
			continue
		}

		fill := make([]string, diff)
		for i := 0; i < diff; i++ {
			fill[i] = " "
		}
		board.tiles[index] = append(board.tiles[index], fill...)
	}

	board.player.position = board.Get2DPosition(board.player.position)
	board.MarkPlayerOnBoard()
}

func lcd(x, y int) int {
	a, b := x, y
	if y > x {
		a, b = y, x
	}

	if b == 0 {
		return x
	}

	return lcd(b, a%b)
}

func (board *Board) FindFaces() {
	board.face_size = lcd(board.height, board.width)
	faces := make([][]int, 0)

	for i := 0; i < board.height; i += board.face_size {
		for j := 0; j < board.width; j += board.face_size {
			if board.tiles[i][j] != " " {
				faces = append(faces, []int{i, j})
			}
		}
	}

	board.faces = map[CubeFace][]int{
		A: faces[0],
		B: faces[1],
		C: faces[2],
		D: faces[3],
		E: faces[4],
		F: faces[5],
	}
}

func (board *Board) FindFace(position []int) CubeFace {
	for face, face_start := range board.faces {
		if position[0] >= face_start[0] && position[0] < face_start[0]+board.face_size &&
			position[1] >= face_start[1] && position[1] < face_start[1]+board.face_size {
			return face
		}
	}
	return -1
}

func (board Board) FindAdjacentPlayerPosition() Player {
	curr_facing := board.player.facing
	curr_face := board.FindFace(board.player.position)
	face_span := board.face_size - 1
	y_distance := board.player.position[0] % board.face_size
	x_distance := board.player.position[1] % board.face_size

	if curr_face == A && curr_facing == Up {
		new_pos := []int{board.faces[F][0] + x_distance, board.faces[F][1]}
		return Player{position: new_pos, facing: Right}
	}

	if curr_face == A && curr_facing == Left {
		new_pos := []int{board.faces[D][0] + (face_span - y_distance), board.faces[D][1]}
		return Player{position: new_pos, facing: Right}
	}

	if curr_face == B && curr_facing == Up {
		new_pos := []int{board.faces[F][0] + face_span, board.faces[F][1] + x_distance}
		return Player{position: new_pos, facing: Up}
	}

	if curr_face == B && curr_facing == Right {
		new_pos := []int{board.faces[E][0] + (face_span - y_distance), board.faces[E][1] + face_span}
		return Player{position: new_pos, facing: Left}
	}

	if curr_face == B && curr_facing == Down {
		new_pos := []int{board.faces[C][0] + x_distance, board.faces[C][1] + face_span}
		return Player{position: new_pos, facing: Left}
	}

	if curr_face == C && curr_facing == Left {
		new_pos := []int{board.faces[D][0], board.faces[D][1] + y_distance}
		return Player{position: new_pos, facing: Down}
	}

	if curr_face == C && curr_facing == Right {
		new_pos := []int{board.faces[B][0] + face_span, board.faces[B][1] + y_distance}
		return Player{position: new_pos, facing: Up}
	}

	if curr_face == D && curr_facing == Up {
		new_pos := []int{board.faces[C][0] + x_distance, board.faces[C][1]}
		return Player{position: new_pos, facing: Right}
	}

	if curr_face == D && curr_facing == Left {
		new_pos := []int{board.faces[A][0] + (face_span - y_distance), board.faces[A][1]}
		return Player{position: new_pos, facing: Right}
	}

	if curr_face == E && curr_facing == Right {
		new_pos := []int{board.faces[B][0] + (face_span - y_distance), board.faces[B][1] + face_span}
		return Player{position: new_pos, facing: Left}
	}

	if curr_face == E && curr_facing == Down {
		new_pos := []int{board.faces[F][0] + x_distance, board.faces[F][1] + face_span}
		return Player{position: new_pos, facing: Left}
	}

	if curr_face == F && curr_facing == Right {
		new_pos := []int{board.faces[E][0] + face_span, board.faces[E][1] + y_distance}
		return Player{position: new_pos, facing: Up}
	}

	if curr_face == F && curr_facing == Down {
		new_pos := []int{board.faces[B][0], board.faces[B][1] + x_distance}
		return Player{position: new_pos, facing: Down}
	}

	if curr_face == F && curr_facing == Left {
		new_pos := []int{board.faces[A][0], board.faces[A][1] + y_distance}
		return Player{position: new_pos, facing: Down}
	}

	log.Fatal("Unsupported edge cross")

	return board.player
}

func (board *Board) Get2DPosition(position []int) []int {
	adjusted_position := []int{(position[0] + board.height) % board.height, (position[1] + board.width) % board.width}
	tile := board.tiles[adjusted_position[0]][adjusted_position[1]]
	if tile != " " {
		return adjusted_position
	}

	return board.Get2DPosition(board.GetNextPosition(adjusted_position))
}

func (board Board) GetNext3DPosition() Player {
	next_position := board.GetNextPosition(board.player.position)
	if next_position[0] >= 0 && next_position[0] < board.height && next_position[1] >= 0 && next_position[1] < board.width {
		if board.tiles[next_position[0]][next_position[1]] != " " {
			return Player{position: next_position, facing: board.player.facing}
		}
	}

	return board.FindAdjacentPlayerPosition()
}

func (board *Board) CanMoveToPosition(position []int) bool {
	if board.tiles[position[0]][position[1]] == "#" {
		return false
	}

	return true
}

func (board *Board) MarkPlayerOnBoard() {
	board.tiles[board.player.position[0]][board.player.position[1]] = board.player.GetFacingSymbol()
}

func (board *Board) MovePlayer(num_moves int, dimensions int) {
	for i := 0; i < num_moves; i++ {
		moved_player := Player{facing: board.player.facing}

		if dimensions == 2 {
			moved_player.position = board.Get2DPosition(board.GetNextPosition(board.player.position))
		} else if dimensions == 3 {
			moved_player = board.GetNext3DPosition()
		}

		if !board.CanMoveToPosition(moved_player.position) {
			break
		}

		board.player.position[0], board.player.position[1] = moved_player.position[0], moved_player.position[1]
		board.player.facing = moved_player.facing
		board.MarkPlayerOnBoard()
	}
}

func main() {
	start := time.Now()

	f, err := os.ReadFile("input.txt")
	if err != nil {
		log.Print(err)
	}

	file_str := strings.Split(string(f), "\n\n")
	map_str, directions := file_str[0], file_str[1]

	player := Player{facing: Right, position: []int{0, 0}}
	board_2d := Board{player: player}
	board_3d := Board{player: player}

	board_2d.BuildTiles(map_str)

	board_3d.BuildTiles(map_str)
	board_3d.FindFaces()

	num_steps_str := ""

	for _, char := range strings.Split(directions, "") {
		if char == "L" || char == "R" {
			num_steps, _ := strconv.Atoi(num_steps_str)
			num_steps_str = ""

			board_2d.MovePlayer(num_steps, 2)
			board_2d.RotatePlayer(char)

			board_3d.MovePlayer(num_steps, 3)
			board_3d.RotatePlayer(char)
			continue
		}

		num_steps_str += char
	}

	if len(num_steps_str) > 0 {
		num_steps, _ := strconv.Atoi(strings.TrimSpace(num_steps_str))
		board_2d.MovePlayer(num_steps, 2)
		board_3d.MovePlayer(num_steps, 3)
	}

	password_2d := board_2d.GetPassword()
	password_3d := board_3d.GetPassword()

	time_elapsed := time.Since(start)

	log.Printf(`
The final 2D password is %d.
The final 3D password is %d.
Solution generated in %s.`,
		password_2d,
		password_3d,
		time_elapsed,
	)
}
