package day_5

import (
	"github.com/voziv/aoc-2022/internal/util"
)

func runPartOne(inputFileName string) string {
	lines := util.GetFileContents(inputFileName)

	input := GetPuzzleInputFromLines(lines)

	//PrintCrates(crates)

	input.crates = RunCrateMover9000(input.moves, input.crates)

	return GetTopCrates(input.crates)
}

// RunCrateMover9000 Moves crates one at a time
func RunCrateMover9000(moves []Move, crates [][]string) [][]string {
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
