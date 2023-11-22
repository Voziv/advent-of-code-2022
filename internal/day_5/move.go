package day_5

import (
	"github.com/voziv/aoc-2022/internal/util"
	"strconv"
)

type Move struct {
	from     int
	to       int
	quantity int
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
