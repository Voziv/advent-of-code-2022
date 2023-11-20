package rock_paper_scissors

import (
	"errors"
	"fmt"
)

type Move int

const (
	Rock     Move = 1
	Paper         = 2
	Scissors      = 3
)

func NewMove(move string) (Move, error) {
	switch move {
	case "A":
		return Rock, nil
	case "X":
		return Rock, nil
	case "B":
		return Paper, nil
	case "Y":
		return Paper, nil
	case "C":
		return Scissors, nil
	case "Z":
		return Scissors, nil
	}

	return 0, errors.New(fmt.Sprintf("'%s' is an invalid move", move))
}

func (m Move) String() string {
	switch m {
	case Rock:
		return "Rock"
	case Paper:
		return "Paper"
	case Scissors:
		return "Scissors"
	}
	return fmt.Sprintf("Invalid %d", m)
}

func getMoveToForceWin(move Move) Move {
	if move == Rock {
		return Paper
	} else if move == Paper {
		return Scissors
	} else {
		return Rock
	}
}

func getMoveToForceDraw(move Move) Move {
	return move
}

func getMoveToForceLoss(move Move) Move {
	if move == Rock {
		return Scissors
	} else if move == Paper {
		return Rock
	} else {
		return Paper
	}
}

func calculateOutcome(ourMove Move, theirMove Move) int {
	if (ourMove == Rock && theirMove == Scissors) || (ourMove == Paper && theirMove == Rock) || (ourMove == Scissors && theirMove == Paper) {
		return Win
	} else if ourMove == theirMove {
		return Draw
	} else {
		return Loss
	}
}
