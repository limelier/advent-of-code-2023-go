package main

import (
	"advent-of-code-2023/common"
	"strconv"
	"strings"
	"unicode"
)

type PartPos struct {
	r int
	c int
}

func main() {
	matrix := make([][]rune, 0)
	ch := common.ReadInput("inputs/day03/input.txt")
	for line := range ch {
		matrix = append(matrix, []rune(line))
	}

	rows := len(matrix)
	cols := len(matrix[0])

	var partNumberSum, gearRatioSum int
	unmatchedGearNumbers := make(map[PartPos]int)
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			ch := matrix[r][c]
			if !unicode.IsDigit(ch) {
				continue
			}

			// consume number
			cPastNumEnd := c + 1
			numSB := strings.Builder{}
			numSB.WriteRune(ch)
			for cPastNumEnd < cols && unicode.IsDigit(matrix[r][cPastNumEnd]) {
				numSB.WriteRune(matrix[r][cPastNumEnd])
				cPastNumEnd++
			}
			number, _ := strconv.Atoi(numSB.String())

			// scan for symbol in bounding box
			hasSymbol := false
		scanLoop:
			for rScan := max(r-1, 0); rScan <= min(r+1, rows-1); rScan++ {
				for cScan := max(c-1, 0); cScan <= min(cPastNumEnd, cols-1); cScan++ {
					chScan := matrix[rScan][cScan]
					if chScan == '.' || unicode.IsDigit(chScan) {
						continue
					}

					// part 2 segment: handle possible gear numbers
					if chScan == '*' {
						key := PartPos{rScan, cScan}
						existing, exists := unmatchedGearNumbers[key]
						if exists {
							gearRatioSum += number * existing
							delete(unmatchedGearNumbers, key)
						} else {
							unmatchedGearNumbers[key] = number
						}
					}

					hasSymbol = true
					break scanLoop
				}
			}

			c = cPastNumEnd
			if hasSymbol {
				partNumberSum += number
			}
		}
	}

	println("Part 1:", partNumberSum)
	println("Part 2:", gearRatioSum)
}
