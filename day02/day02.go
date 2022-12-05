package day02

import (
	"strings"
)

// Opponent
// A = Rock
// B = Paper
// C = Scissors
const oROCK = "A"
const oPAPER = "B"
const oSCISSORS = "C"

// Me
// X = Rock
// Y = Paper
// Z = Scissors
const mROCK = "X"
const mPAPER = "Y"
const mSCISSORS = "Z"

// Game
const WIN = "W"
const DRAW = "D"
const LOSE = "L"

var POINTS = map[string]int{
	mROCK:     1,
	mPAPER:    2,
	mSCISSORS: 3,
	WIN:       6,
	DRAW:      3,
	LOSE:      0,
}

func Part1(input string) int {
	totalScore := 0
	rounds := strings.Split(input, "\n")
	for _, r := range rounds {
		games := strings.Split(r, " ")
		if len(games) != 2 {
			continue
		}
		opponent := games[0]
		me := games[1]

		totalScore += scoreOfRPS(opponent, me)
	}

	return totalScore
}

func scoreOfRPS(opponent, me string) int {
	return POINTS[me] + POINTS[resultOfRPS(opponent, me)]
}

func resultOfRPS(opponent, me string) (result string) {
	result = DRAW
	if me == mROCK {
		if opponent == oSCISSORS {
			result = WIN
		}
		if opponent == oPAPER {
			result = LOSE
		}
		return
	}
	if me == mPAPER {
		if opponent == oROCK {
			result = WIN
		}
		if opponent == oSCISSORS {
			result = LOSE
		}
		return
	}
	if me == mSCISSORS {
		if opponent == oPAPER {
			result = WIN
		}
		if opponent == oROCK {
			result = LOSE
		}
		return
	}
	return
}

func Part2(input string) int {
	totalScore := 0
	rounds := strings.Split(input, "\n")
	for _, r := range rounds {
		games := strings.Split(r, " ")
		if len(games) != 2 {
			continue
		}
		opponent := games[0]

		totalScore += scoreOfRPS(opponent, whatNeedToDo(opponent, games[1]))
	}

	return totalScore
}

func whatNeedToDo(opponent, me string) string {
	if me == mROCK { // Need to loose
		if opponent == oSCISSORS {
			return mPAPER
		}
		if opponent == oROCK {
			return mSCISSORS
		}
		return me
	}
	if me == mPAPER { // Need to draw
		if opponent == oSCISSORS {
			return mSCISSORS
		}
		if opponent == oROCK {
			return mROCK
		}
		return me
	}
	if me == mSCISSORS { // Need to win
		if opponent == oSCISSORS {
			return mROCK
		}
		if opponent == oROCK {
			return mPAPER
		}
		return me
	}
	return me
}
