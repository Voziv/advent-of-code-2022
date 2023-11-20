package rock_paper_scissors

// X = Force Loss
// Y = Force Draw
// Z = Force Win

// Scoring Points (What you played + outcome)
// A = Rock = 1
// B = Paper = 2
// C = Scissors = 3
// Loss = 0
// Draw = 3
// Win = 6

const (
	Loss = 0
	Draw = 3
	Win  = 6
)

type Round struct {
	line          string
	expectedScore int
	roundNumber   int
	ourMove       Move
	theirMove     Move
	forceOutcome  string
}

func NewRound(line string, expectedScore int, roundNumber int, theirMove Move, ourMove Move, forceOutcome string) *Round {
	return &Round{
		line:          line,
		expectedScore: expectedScore,
		roundNumber:   roundNumber,
		ourMove:       ourMove,
		theirMove:     theirMove,
		forceOutcome:  forceOutcome,
	}
}

func (r *Round) Score() int {
	return int(r.ourMove) + calculateOutcome(r.ourMove, r.theirMove)
}

// ScoreWithSecretStrategy
// X = Force Loss
// Y = Force Draw
// Z = Force Win
func (r *Round) ScoreWithSecretStrategy() int {
	move := r.ourMove
	if r.forceOutcome == "X" { // Force Loss
		move = getMoveToForceLoss(r.theirMove)
	} else if r.forceOutcome == "Y" { // Force Draw
		move = getMoveToForceDraw(r.theirMove)
	} else if r.forceOutcome == "Z" { // Force Win
		move = getMoveToForceWin(r.theirMove)
	}
	return int(move) + calculateOutcome(move, r.theirMove)
}

func (r *Round) RoundNumber() int {
	return r.roundNumber
}

func (r *Round) OurMove() Move {
	return r.ourMove
}

func (r *Round) TheirMove() Move {
	return r.theirMove
}
