package day_5

import (
	"fmt"
	"github.com/voziv/aoc-2022/internal/util"
	"strconv"
	"strings"
)

func runPartOne(inputFileName string) string {
	lines := util.GetFileContents(inputFileName)
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

	//PrintCrates(crates)

	crates = ProcessMoves(moves, crates)

	return GetTopCrates(crates)
}

func ProcessMoves(moves []Move, crates [][]string) [][]string {
	for _, move := range moves {
		//fmt.Printf("Moving %d from %d to %d\n", move.quantity, move.from, move.to)
		for i := 0; i < move.quantity; i++ {
			crate := crates[move.from][len(crates[move.from])-1]
			crates[move.from] = crates[move.from][:len(crates[move.from])-1]
			crates[move.to] = append(crates[move.to], crate)
		}
	}

	return crates
}

func GetTopCrates(crates [][]string) string {
	result := ""
	for i := 0; i < len(crates); i++ {
		result += crates[i][len(crates[i])-1]
	}
	return result
}

func PrintCrates(crates [][]string) {
	fmt.Println("Crates")
	highest := 0
	for i := 0; i < len(crates); i++ {
		if len(crates[i]) > highest {
			highest = len(crates[i])
		}
	}

	for j := highest; j >= 0; j-- {
		for i := 0; i < len(crates); i++ {
			if len(crates[i]) > j {
				fmt.Printf("(%s) ", crates[i][j])
			} else {
				fmt.Printf("    ")
			}
		}
		fmt.Println()
	}

	for i := 0; i < len(crates); i++ {
		fmt.Printf(" %-2d ", i)
	}
	fmt.Println()
	fmt.Println("End of Crates")
}

func NewMoveFromInstruction(tokens []string) Move {
	quantity, err := strconv.Atoi(tokens[1])
	util.CheckErr(err)

	from, err := strconv.Atoi(tokens[3])
	util.CheckErr(err)

	to, err := strconv.Atoi(tokens[5])
	util.CheckErr(err)

	return Move{
		from - 1,
		to - 1,
		quantity,
	}
}

type Move struct {
	from     int
	to       int
	quantity int
}
