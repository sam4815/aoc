package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
	"time"
)

type Die struct {
	Roll func() int
}

type Player struct {
	score    int
	position int
}

type Game struct {
	turn                     int
	player_one               Player
	player_two               Player
	die                      Die
	num_outcomes_represented int
}

func (die *Die) Init() {
	n := 0
	die.Roll = func() int {
		n = (n % 100) + 1
		return n
	}
}

func (game Game) Clone() Game {
	return Game{
		turn:                     game.turn,
		player_one:               Player{score: game.player_one.score, position: game.player_one.position},
		player_two:               Player{score: game.player_two.score, position: game.player_two.position},
		num_outcomes_represented: game.num_outcomes_represented,
	}
}

func (player *Player) Move(steps int) {
	player.position = ((player.position + steps - 1) % 10) + 1
	player.score += player.position
}

func (game *Game) PlayRound() {
	rolls := game.die.Roll() + game.die.Roll() + game.die.Roll()

	if game.turn%2 == 0 {
		game.player_one.Move(rolls)
	} else {
		game.player_two.Move(rolls)
	}

	game.turn += 1
}

func (game Game) GetOutcomes() []Game {
	outcomes := make([]Game, 0)
	permutation_frequency := map[int]int{
		3: 1,
		4: 3,
		5: 6,
		6: 7,
		7: 6,
		8: 3,
		9: 1,
	}

	for i := 3; i <= 9; i++ {
		outcome := game.Clone()
		if outcome.turn%2 == 0 {
			outcome.player_one.Move(i)
		} else {
			outcome.player_two.Move(i)
		}
		outcome.turn += 1
		outcome.num_outcomes_represented *= permutation_frequency[i]
		outcomes = append(outcomes, outcome)
	}

	return outcomes
}

func (game Game) HasWinner(target int) bool {
	return game.player_one.score >= target || game.player_two.score >= target
}

func (game Game) FinalScore() int {
	losing_score := game.player_one.score
	if game.player_two.score < losing_score {
		losing_score = game.player_two.score
	}

	return losing_score * game.turn * 3
}

func (game Game) Stringify() string {
	return fmt.Sprintf(
		"Round %d. Player 1 score: %d, position: %d. Player 2 score: %d, position: %d.",
		game.turn,
		game.player_one.score,
		game.player_one.position,
		game.player_two.score,
		game.player_two.position,
	)
}

func (game Game) Permutations(target int) (int, int) {
	queue, curr_game := []Game{game}, game
	one_wins, two_wins := 0, 0

	for len(queue) > 0 {
		curr_game, queue = queue[0], queue[1:]

		if curr_game.HasWinner(target) {
			if curr_game.player_one.score > curr_game.player_two.score {
				one_wins += curr_game.num_outcomes_represented
			} else {
				two_wins += curr_game.num_outcomes_represented
			}
			continue
		}

		queue = append(queue, curr_game.GetOutcomes()...)
	}

	return one_wins, two_wins
}

func main() {
	start := time.Now()
	f, _ := ioutil.ReadFile("input.txt")
	player_strings := strings.Split(strings.TrimSpace(string(f)), "\n")
	player_one_pos_str := string(player_strings[0][len(player_strings[0])-1])
	player_two_pos_str := player_strings[1][len(player_strings[1])-1]
	player_one_pos, _ := strconv.Atoi(string(player_one_pos_str))
	player_two_pos, _ := strconv.Atoi(string(player_two_pos_str))

	deterministic_game := Game{
		die:        Die{},
		player_one: Player{position: player_one_pos},
		player_two: Player{position: player_two_pos},
	}
	deterministic_game.die.Init()

	for !deterministic_game.HasWinner(1000) {
		deterministic_game.PlayRound()
	}
	final_score := deterministic_game.FinalScore()

	game := Game{
		player_one:               Player{position: player_one_pos},
		player_two:               Player{position: player_two_pos},
		num_outcomes_represented: 1,
	}

	most_wins, player_two_wins := game.Permutations(21)
	if player_two_wins > most_wins {
		most_wins = player_two_wins
	}

	time_elapsed := time.Since(start)

	log.Printf(`
The final score after the game is won is %d.
The player with the greatest number of wins wins in %d universes.
Solution generated in %s.`,
		final_score,
		most_wins,
		time_elapsed,
	)
}
