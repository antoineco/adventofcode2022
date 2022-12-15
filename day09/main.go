package day09

import (
	"bufio"
	"fmt"
	"math"
	"strconv"

	"adventofcode2022/fs"
)

func Puzzle1(fs fs.FS, inputFile string) int {
	f, err := fs.Open(inputFile)
	if err != nil {
		panic(fmt.Errorf("opening input file: %w", err))
	}

	s := bufio.NewScanner(f)

	var head, tail position

	visitedByTail := make(set)
	visitedByTail.add(fmt.Sprint(tail.x, ",", tail.y))

	for s.Scan() {
		d, steps := readMove(s.Bytes())
		switch d {
		case up:
			head.y += steps
		case down:
			head.y -= steps
		case left:
			head.x -= steps
		case right:
			head.x += steps
		}

		tail = follow(head, tail, visitedByTail)
	}

	if err := s.Err(); err != nil {
		panic(fmt.Errorf("while scanning input file: %w", err))
	}

	return len(visitedByTail)
}

func Puzzle2(fs fs.FS, inputFile string) int {
	f, err := fs.Open(inputFile)
	if err != nil {
		panic(fmt.Errorf("opening input file: %w", err))
	}

	s := bufio.NewScanner(f)

	var knots [10]position

	visitedByTail := make(set)
	visitedByTail.add(fmt.Sprint(knots[9].x, ",", knots[9].y))

	for s.Scan() {
		d, steps := readMove(s.Bytes())
		switch d {
		case up:
			knots[0].y += steps
		case down:
			knots[0].y -= steps
		case left:
			knots[0].x -= steps
		case right:
			knots[0].x += steps
		}

		knots = followAll(knots, visitedByTail)
	}

	if err := s.Err(); err != nil {
		panic(fmt.Errorf("while scanning input file: %w", err))
	}

	return len(visitedByTail)
}

func readMove(input []byte) (d direction, steps int) {
	switch input[0] {
	case 'U':
		d = up
	case 'D':
		d = down
	case 'L':
		d = left
	case 'R':
		d = right
	}

	var err error
	steps, err = strconv.Atoi(string(input[2:]))
	if err != nil {
		panic(fmt.Errorf("parsing %v to integer: %w", input[2:], err))
	}

	return d, steps
}

func follow(head, tail position, visitedByTail set) (newTail position) {
	for !isTouching(head, tail) {
		tail = moveTail(head, tail)

		visitedByTail.add(fmt.Sprint(tail.x, ",", tail.y))
	}

	return tail
}

func followAll(knots [10]position, visitedByTail set) [10]position {
	for i := 1; i < len(knots)-1; i++ {
		for !isTouching(knots[i-1], knots[i]) {
			knots[i] = moveTail(knots[i-1], knots[i])

			for j := i + 1; j < len(knots); j++ {
				for !isTouching(knots[j-1], knots[j]) {
					knots[j] = moveTail(knots[j-1], knots[j])
					if j == len(knots)-1 {
						visitedByTail.add(fmt.Sprint(knots[j].x, ",", knots[j].y))
					}
				}
			}
		}
	}

	return knots
}

func isTouching(head, tail position) bool {
	return math.Abs(float64(head.x-tail.x)) <= 1 && math.Abs(float64(head.y-tail.y)) <= 1
}

func moveTail(head, tail position) (newPosition position) {
	switch {
	// straight moves
	case tail.x == head.x && head.y-tail.y > 1:
		tail.y++
	case tail.x == head.x && head.y-tail.y < -1:
		tail.y--
	case tail.y == head.y && head.x-tail.x > 1:
		tail.x++
	case tail.y == head.y && head.x-tail.x < -1:
		tail.x--
	// diagonal moves
	case head.y > tail.y && head.x < tail.x:
		tail.x--
		tail.y++
	case head.y > tail.y && head.x > tail.x:
		tail.x++
		tail.y++
	case head.y < tail.y && head.x < tail.x:
		tail.x--
		tail.y--
	case head.y < tail.y && head.x > tail.x:
		tail.x++
		tail.y--
	}

	return tail
}

type position struct {
	x int
	y int
}

type direction uint8

const (
	up direction = iota
	down
	left
	right
)

type set map[string]struct{}

func (s *set) has(v string) bool {
	_, exists := (*s)[v]
	return exists
}

func (s *set) add(v string) {
	if !s.has(v) {
		(*s)[v] = struct{}{}
	}
}
