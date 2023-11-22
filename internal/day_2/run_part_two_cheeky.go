package day_2

import (
	"github.com/voziv/aoc-2022/internal/util"
	"strconv"
)

var CheekyPreCalc = map[string]int{
	"A X": 3 + 0, // Rock and must lose. Play scissors
	"B X": 1 + 0, // Paper and must lose. Play Rock
	"C X": 2 + 0, // Scissors and must lose. Play paper
	"A Y": 1 + 3, // Rock and must draw. Play rock
	"B Y": 2 + 3, // Paper and must draw. Play paper
	"C Y": 3 + 3, // Scissors and must draw. Play scissors
	"A Z": 2 + 6, // Rock and must win. Play paper
	"B Z": 3 + 6, // Paper and must win. Play scissors
	"C Z": 1 + 6, // Scissors and must win. Play paper
}

func runPartTwoCheeky(inputFileName string) string {
	moves := util.GetFileContents(inputFileName)

	preCalcTotal := 0

	for _, move := range moves {
		if move != "" {
			preCalcTotal += CheekyPreCalc[move]
		}
	}

	return strconv.Itoa(preCalcTotal)
}
