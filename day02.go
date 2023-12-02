package main

import (
	"advent-of-code-2023/common"
	"advent-of-code-2023/day02"
)

type Game = day02.Game
type Bag = day02.Bag

func main() {
	ch := common.ReadInput("inputs/day02/input.txt")
	games := make([]*Game, 0)
	for line := range ch {
		games = append(games, day02.NewGame(line))
	}

	// part 1
	bag := Bag{Red: 12, Green: 13, Blue: 14}
	idSum := 0
	for _, game := range games {
		if game.IsPossibleWith(&bag) {
			idSum += game.Id
		}
	}
	println("Part 1:", idSum)

	// part 2
	powerSum := 0
	for _, game := range games {
		powerSum += game.MinimumBag().Power()
	}
	println("Part 2:", powerSum)
}
