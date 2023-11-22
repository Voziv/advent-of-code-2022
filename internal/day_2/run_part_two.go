package day_2

import (
	"github.com/voziv/aoc-2022/internal/util"
	"strconv"
	"strings"
)

func runPartTwo(inputFileName string) string {
	moves := util.GetFileContents(inputFileName)

	var rounds []*Round

	for index, move := range moves {

		if move != "" {
			moveArray := strings.Split(move, " ")
			theirMove, err := NewMove(moveArray[0])
			if err != nil {
				panic(err)
			}

			ourMove, err := NewMove(moveArray[1])
			if err != nil {
				panic(err)
			}

			rounds = append(rounds, NewRound(move, index, theirMove, ourMove, moveArray[1]))
		}
	}

	ourScoreWithStrategy := 0
	for _, round := range rounds {
		ourScoreWithStrategy += round.ScoreWithSecretStrategy()
	}

	return strconv.Itoa(ourScoreWithStrategy)
}
