package day_5

import (
	"github.com/voziv/aoc-2022/internal/util"
)

func runPartTwo(inputFileName string) string {
	lines := util.GetFileContents(inputFileName)

	input := GetPuzzleInputFromLines(lines)

	//PrintCrates(crates)

	input.crates = RunCrateMover9001(input.moves, input.crates)

	return GetTopCrates(input.crates)
}

// RunCrateMover9001 Moves crates in stacks, so if a moving 3, it moves all 3 as a stack thus preserving the order
func RunCrateMover9001(moves []Move, crates [][]string) [][]string {
	for _, move := range moves {
		//fmt.Printf("Moving %d from %d to %d\n", move.quantity, move.from, move.to)
		cratesToMove := crates[move.from][len(crates[move.from])-move.quantity:]
		crates[move.from] = crates[move.from][:len(crates[move.from])-move.quantity]
		crates[move.to] = append(crates[move.to], cratesToMove...)
	}

	return crates
}
