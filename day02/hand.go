package day02

import (
	"log"
	"strconv"
	"strings"
)

// Hand is a set of colored cubes drawn from a Bag during a Game
type Hand struct {
	Red   int
	Green int
	Blue  int
}

// NewHand parses part of a line of the puzzle input into a valid Hand in a Game
func NewHand(s string) *Hand {
	hand := Hand{}
	components := strings.Split(s, ",")

	for _, component := range components {
		component = strings.TrimSpace(component)
		parts := strings.Split(component, " ")
		number, err := strconv.Atoi(parts[0])
		if err != nil {
			log.Panicf("Illegal number in hand string: '%s' <- '%s' <- '%s'", parts[0], component, s)
		}

		switch parts[1] {
		case "red":
			hand.Red = number
		case "green":
			hand.Green = number
		case "blue":
			hand.Blue = number
		default:
			log.Panicf("Illegal color in hand string: '%s'", parts[1])
		}
	}

	return &hand
}
