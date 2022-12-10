package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

type Weapon int8
type Result int8
type WeaponKey struct {
	X, Y Weapon
}
type WeaponResultKey struct {
	X Weapon
	Y Result
}

const (
	Rock     Weapon = 1
	Paper           = 2
	Scissors        = 3
)
const (
	Lose Result = 0
	Draw        = 3
	Win         = 6
)

func main() {
	outcomes := map[WeaponKey]Result{
		{Rock, Scissors}:     Win,
		{Rock, Rock}:         Draw,
		{Rock, Paper}:        Lose,
		{Paper, Rock}:        Win,
		{Paper, Paper}:       Draw,
		{Paper, Scissors}:    Lose,
		{Scissors, Paper}:    Win,
		{Scissors, Scissors}: Draw,
		{Scissors, Rock}:     Lose,
	}

	required_weapons := map[WeaponResultKey]Weapon{
		{Rock, Win}:      Paper,
		{Rock, Draw}:     Rock,
		{Rock, Lose}:     Scissors,
		{Paper, Win}:     Scissors,
		{Paper, Draw}:    Paper,
		{Paper, Lose}:    Rock,
		{Scissors, Win}:  Rock,
		{Scissors, Draw}: Scissors,
		{Scissors, Lose}: Paper,
	}

	weapon_map := map[string]Weapon{
		"A": Rock,
		"B": Paper,
		"C": Scissors,
		"X": Rock,
		"Y": Paper,
		"Z": Scissors,
	}

	outcome_map := map[string]Result{
		"X": Lose,
		"Y": Draw,
		"Z": Win,
	}

	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)

	correct_score, naive_score := 0, 0

	for scanner.Scan() {
		round := strings.Split(scanner.Text(), " ")

		opponent_weapon, my_weapon := weapon_map[round[0]], weapon_map[round[1]]
		outcome := outcomes[WeaponKey{my_weapon, opponent_weapon}]

		naive_score += int(outcome) + int(my_weapon)

		desired_outcome := outcome_map[round[1]]
		required_weapon := required_weapons[WeaponResultKey{opponent_weapon, desired_outcome}]

		correct_score += int(desired_outcome) + int(required_weapon)
	}

	log.Printf(`
Following my interpretation of the strategy guide, my total score is %d.
Following the strategy guide properly, my total score is %d.`,
		naive_score,
		correct_score,
	)
}
