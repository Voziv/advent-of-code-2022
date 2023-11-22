package day_5

import (
	"strings"
)

type PuzzleInput struct {
	crates [][]string
	moves  []Move
}

func GetPuzzleInputFromLines(lines []string) PuzzleInput {
	var moves []Move
	var crates [][]string

	for _, line := range lines {
		if strings.Index(line, "[") != -1 {
			column := 0
			for i := 0; i < len(line); i += 4 {
				if len(crates) <= column {
					crates = append(crates, []string{})
				}

				if line[i:i+1] == "[" {
					crate := line[i+1 : i+2]
					crates[column] = append([]string{crate}, crates[column]...)
				}
				column++
			}
		} else if strings.Index(line, "move") != -1 {
			tokens := strings.Split(line, " ")
			moves = append(moves, NewMoveFromInstruction(tokens))
		}
	}

	return PuzzleInput{
		crates: crates,
		moves:  moves,
	}
}
