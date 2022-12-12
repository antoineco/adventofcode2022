package day06

import (
	"bufio"
	"fmt"

	"adventofcode2022/fs"
)

func Puzzle1(fs fs.FS, inputFile string) int {
	f, err := fs.Open(inputFile)
	if err != nil {
		panic(fmt.Errorf("opening input file: %w", err))
	}

	s := bufio.NewScanner(f)

	s.Scan()

	if err := s.Err(); err != nil {
		panic(fmt.Errorf("while scanning input file: %w", err))
	}

	input := s.Bytes()

	const winSize = 4

	if len(input) < winSize {
		return -1
	}

	// Track number of characters in current 4-char sliding window.
	// Sum is always exactly 4.
	// Decrement leftmost value as the window slides, remove if 0.
	charWindow := make(map[byte]int)

	for i := 0; i < winSize-1; i++ {
		charWindow[input[i]]++
	}
	if len(charWindow) == winSize-1 && // 3 distinct values in first 3 chars
		charWindow[input[winSize-1]] == 0 { // 4th char not already present

		return winSize
	}

	charWindow[input[winSize-1]]++

	for i := winSize; i < len(input); i++ {
		if charWindow[input[i-winSize]] == 1 {
			delete(charWindow, input[i-winSize])
		} else {
			charWindow[input[i-winSize]]--
		}

		if len(charWindow) == winSize-1 && charWindow[input[i]] == 0 { // same as initial case
			return i + 1
		}

		charWindow[input[i]]++
	}

	return -1
}

func Puzzle2(fs fs.FS, inputFile string) int {
	f, err := fs.Open(inputFile)
	if err != nil {
		panic(fmt.Errorf("opening input file: %w", err))
	}

	s := bufio.NewScanner(f)

	s.Scan()

	if err := s.Err(); err != nil {
		panic(fmt.Errorf("while scanning input file: %w", err))
	}

	input := s.Bytes()

	const winSize = 14

	if len(input) < winSize {
		return -1
	}

	// Track number of characters in current 14-char sliding window.
	// Sum is always exactly 14.
	// Decrement leftmost value as the window slides, remove if 0.
	charWindow := make(map[byte]int)

	for i := 0; i < winSize-1; i++ {
		charWindow[input[i]]++
	}
	if len(charWindow) == winSize-1 && // 13 distinct values in first 13 chars
		charWindow[input[winSize-1]] == 0 { // 14th char not already present

		return winSize
	}

	charWindow[input[winSize-1]]++

	for i := winSize; i < len(input); i++ {
		if charWindow[input[i-winSize]] == 1 {
			delete(charWindow, input[i-winSize])
		} else {
			charWindow[input[i-winSize]]--
		}

		if len(charWindow) == winSize-1 && charWindow[input[i]] == 0 { // same as initial case
			return i + 1
		}

		charWindow[input[i]]++
	}

	return -1
}
