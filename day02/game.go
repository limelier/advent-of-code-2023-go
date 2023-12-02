package day02

import (
	"log"
	"strconv"
	"strings"
)

// Game represents a succession of hands drawn from a Bag
type Game struct {
	// Id uniquely identifies the game in the context of the day's puzzle
	Id    int
	hands []*Hand
}

// NewGame parses a line of the puzzle input into a Game
func NewGame(line string) *Game {
	halves := strings.Split(line, ":")

	header := strings.Split(halves[0], " ")
	id, err := strconv.Atoi(header[1])
	if err != nil {
		log.Panicf("Illegal non-numeric ID in game line: %s", header[1])
	}

	body := strings.Split(halves[1], "; ")
	hands := make([]*Hand, len(body))

	for i, s := range body {
		hands[i] = NewHand(s)
	}

	return &Game{id, hands}
}

// IsPossibleWith checks if the game could have been played with a certain Bag
func (g *Game) IsPossibleWith(bag *Bag) bool {
	for _, hand := range g.hands {
		if !bag.CanDraw(hand) {
			return false
		}
	}
	return true
}

// MinimumBag finds the Bag containing the minimum amount of cubes required to play the game
func (g *Game) MinimumBag() *Bag {
	bag := Bag{}
	for _, hand := range g.hands {
		bag.ExpandToContain(hand)
	}

	return &bag
}
