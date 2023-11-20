package rock_paper_scissors

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var CHEEKY_PRE_CALC = map[string]int{
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

func Play(inputFileName string) {
	moves := ParseInput(inputFileName)

	var rounds []*Round

	preCalcTotal := 0

	for index, move := range moves {

		if move != "" {
			preCalcTotal += CHEEKY_PRE_CALC[move]

			moveArray := strings.Split(move, " ")
			theirMove, err := NewMove(moveArray[0])
			if err != nil {
				panic(err)
			}

			ourMove, err := NewMove(moveArray[1])
			if err != nil {
				panic(err)
			}

			rounds = append(rounds, NewRound(move, CHEEKY_PRE_CALC[move], index, theirMove, ourMove, moveArray[1]))
		}
	}

	ourScore := 0
	ourScoreWithStrategy := 0
	for index, round := range rounds {
		ourScore += round.Score()
		ourScoreWithStrategy += round.ScoreWithSecretStrategy()
		if round.ScoreWithSecretStrategy() != round.expectedScore {
			fmt.Printf("Round %d - %s, Our score %d. With secret strategy our score we scored %d\n", index, round.line, round.Score(), round.ScoreWithSecretStrategy())
		}
	}

	fmt.Printf("Precalc Total: %d\n", preCalcTotal)

	fmt.Printf("Tournament finished. Our score %d. With secret strategy our score we scored %d\n", ourScore, ourScoreWithStrategy)
}

func ParseInput(inputFile string) []string {
	file, err := os.Open(inputFile)
	if err != nil {
		panic(err)
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)

	scanner := bufio.NewScanner(file)

	// All the problem wants is the highest total amount, as well as the top three totals summed.
	// We don't actually need to know which elf has what.
	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	// TODO lrobert: Not really sure what this error is supposed to be catching. Examples have been very unclear.
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return lines
}
