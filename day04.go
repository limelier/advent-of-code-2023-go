package main

import (
	"advent-of-code-2023/common"
	"strconv"
	"strings"
)

func main() {
	ch := common.ReadInput("inputs/day04/input.txt")

	totalPoints := 0
	cardWins := make(map[int]int)
	for line := range ch {
		sections := strings.Split(line, ": ")
		id, _ := strconv.Atoi(strings.Fields(sections[0])[1])
		parts := strings.Split(sections[1], " | ")

		winners := make(map[int]bool)
		for _, numStr := range strings.Fields(parts[0]) {
			num, _ := strconv.Atoi(numStr)
			winners[num] = true
		}

		points := 0
		wins := 0
		for _, numStr := range strings.Fields(parts[1]) {
			num, _ := strconv.Atoi(numStr)
			_, ok := winners[num]
			if ok {
				if points == 0 {
					points = 1
				} else {
					points *= 2
				}
				wins++
			}
		}

		totalPoints += points
		cardWins[id] = wins
	}

	println("Part 1:", totalPoints)

	totalCards := 0
	wonCopies := make(map[int]int)
	for i := 1; i <= len(cardWins); i++ {
		// get previously-won copies of this card
		copies, ok := wonCopies[i]
		if !ok {
			copies = 0
		}
		copies++             // add the original copy of every card
		totalCards += copies // add all copies of this card to the total

		// copy subsequent cards as needed
		wins := cardWins[i]
		for j := i + 1; j <= i+wins; j++ {
			wonCopies[j] += copies
		}
	}
	println("Part 2:", totalCards)
}
