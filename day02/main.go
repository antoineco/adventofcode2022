package day02

import (
	"bufio"
	"bytes"
	"fmt"

	"adventofcode2022/fs"
)

func Puzzle1(fs fs.FS, inputFile string) int {
	f, err := fs.Open(inputFile)
	if err != nil {
		panic(fmt.Errorf("opening input file: %w", err))
	}

	totalScore := 0

	s := bufio.NewScanner(f)

	for s.Scan() {
		codes := bytes.Fields(s.Bytes())
		if l := len(codes); l != 2 {
			panic(fmt.Errorf("read line with %d column(s) instead of 2", l))
		}
		if l := len(codes[0]); l != 1 {
			panic(fmt.Errorf("left field has %d byte(s) instead of 1", l))
		}
		if l := len(codes[1]); l != 1 {
			panic(fmt.Errorf("right field has %d byte(s) instead of 1", l))
		}

		opponentMove, selfMove := readMove(codes[0][0]), readMove(codes[1][0])

		totalScore += playMoves(opponentMove, selfMove).score() + selfMove.score()
	}

	return totalScore
}

func Puzzle2(fs fs.FS, inputFile string) int {
	f, err := fs.Open(inputFile)
	if err != nil {
		panic(fmt.Errorf("opening input file: %w", err))
	}

	totalScore := 0

	s := bufio.NewScanner(f)

	for s.Scan() {
		codes := bytes.Fields(s.Bytes())
		if l := len(codes); l != 2 {
			panic(fmt.Errorf("read line with %d column(s) instead of 2", l))
		}
		if l := len(codes[0]); l != 1 {
			panic(fmt.Errorf("left field has %d byte(s) instead of 1", l))
		}
		if l := len(codes[1]); l != 1 {
			panic(fmt.Errorf("right field has %d byte(s) instead of 1", l))
		}

		opponentMove := readMove(codes[0][0])
		selfMove := needMove(opponentMove, readOutcome(codes[1][0]))

		totalScore += playMoves(opponentMove, selfMove).score() + selfMove.score()
	}

	return totalScore
}

type move int8

const (
	moveUnknown move = iota - 1
	moveRock
	movePaper
	moveScissors
)

func (m move) score() int {
	switch m {
	case moveRock:
		return 1
	case movePaper:
		return 2
	case moveScissors:
		return 3
	default:
		return -1
	}
}

func readMove(code byte) move {
	switch code {
	case 'A', 'X':
		return moveRock
	case 'B', 'Y':
		return movePaper
	case 'C', 'Z':
		return moveScissors
	default:
		return moveUnknown
	}
}

type outcome int8

const (
	outcomeUnknown outcome = iota - 1
	outcomeLoss
	outcomeDraw
	outcomeWin
)

func (o outcome) score() int {
	switch o {
	case outcomeLoss:
		return 0
	case outcomeDraw:
		return 3
	case outcomeWin:
		return 6
	default:
		return -1
	}
}

func readOutcome(code byte) outcome {
	switch code {
	case 'X':
		return outcomeLoss
	case 'Y':
		return outcomeDraw
	case 'Z':
		return outcomeWin
	default:
		return outcomeUnknown
	}
}

// Although it doesn't result in the shortest possible code, the game is simple
// enough to enumerate all possible moves in the two functions below.

func playMoves(opponent, self move) outcome {
	switch {
	case opponent == moveUnknown || self == moveUnknown:
		return outcomeUnknown
	case opponent == self:
		return outcomeDraw
	case opponent == moveRock && self == movePaper:
		return outcomeWin
	case opponent == movePaper && self == moveScissors:
		return outcomeWin
	case opponent == moveScissors && self == moveRock:
		return outcomeWin
	default:
		return outcomeLoss
	}
}

func needMove(opponent move, expect outcome) move {
	switch {
	case opponent == moveUnknown || expect == outcomeUnknown:
		return moveUnknown
	case expect == outcomeDraw:
		return opponent
	case opponent == moveRock && expect == outcomeLoss:
		return moveScissors
	case opponent == moveRock && expect == outcomeWin:
		return movePaper
	case opponent == movePaper && expect == outcomeLoss:
		return moveRock
	case opponent == movePaper && expect == outcomeWin:
		return moveScissors
	case opponent == moveScissors && expect == outcomeLoss:
		return movePaper
	case opponent == moveScissors && expect == outcomeWin:
		return moveRock
	default:
		return moveUnknown
	}
}
