package main

import (
	"advent-of-code-2023/common"
	"fmt"
	"log"
	"regexp"
)

// TextToDigit converts a digit string ("one", "1") to a number
func TextToDigit(text string) int {
	switch text {
	case "one", "1":
		return 1
	case "two", "2":
		return 2
	case "three", "3":
		return 3
	case "four", "4":
		return 4
	case "five", "5":
		return 5
	case "six", "6":
		return 6
	case "seven", "7":
		return 7
	case "eight", "8":
		return 8
	case "nine", "9":
		return 9
	case "zero", "0":
		return 0
	default:
		// log.Panicf doesn't work here, thanks go!
		log.Printf("TextToDigit called with non-digit %s", text)
		panic(nil)
	}
}

// FindFirstAndLastDigit returns the number created by the first and last digits (according to the regexp) found in line
func FindFirstAndLastDigit(regexp *regexp.Regexp, line string) int {
	// can't simply find all matches; digits may overlap ('eightwo' -> 82)

	first := TextToDigit(regexp.FindString(line))

	last := ""
	lastIdx := len(line) - 1
	for last == "" {
		last = regexp.FindString(line[lastIdx:])
		lastIdx--
	}
	lastNum := TextToDigit(last)

	fmt.Printf("%s -> %d %d -> %d\n", line, first, lastNum, first*10+lastNum)
	return first*10 + lastNum
}

func main() {
	ch := common.ReadInput("inputs/day01/input.txt")
	var part1Sum, part2Sum int

	for line := range ch {
		part1Sum += FindFirstAndLastDigit(regexp.MustCompile(`\d`), line)
		part2Regexp := regexp.MustCompile(`one|two|three|four|five|six|seven|eight|nine|zero|\d`)
		part2Sum += FindFirstAndLastDigit(part2Regexp, line)
	}

	println(part1Sum)
	println(part2Sum)
}
