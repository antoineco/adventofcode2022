package day04

import (
	"bufio"
	"bytes"
	"fmt"
	"strconv"

	"adventofcode2022/fs"
)

func Puzzle1(fs fs.FS, inputFile string) int {
	f, err := fs.Open(inputFile)
	if err != nil {
		panic(fmt.Errorf("opening input file: %w", err))
	}

	fullOverlaps := 0

	s := bufio.NewScanner(f)

	for s.Scan() {
		elf1Range, elf2Range := ranges(s.Bytes())

		elf1Low, elf1High := rangeBoundaries(elf1Range)
		elf2Low, elf2High := rangeBoundaries(elf2Range)

		if (elf1Low <= elf2Low && elf2Low <= elf1High && elf2High <= elf1High) ||
			(elf2Low <= elf1Low && elf1Low <= elf2High && elf1High <= elf2High) {

			fullOverlaps++
		}
	}

	return fullOverlaps
}

func Puzzle2(fs fs.FS, inputFile string) int {
	f, err := fs.Open(inputFile)
	if err != nil {
		panic(fmt.Errorf("opening input file: %w", err))
	}

	overlappingPairs := 0

	s := bufio.NewScanner(f)

	for s.Scan() {
		elf1Range, elf2Range := ranges(s.Bytes())

		elf1Low, elf1High := rangeBoundaries(elf1Range)
		elf2Low, elf2High := rangeBoundaries(elf2Range)

		if (elf1Low <= elf2Low && elf1High >= elf2Low) ||
			(elf2Low <= elf1Low && elf2High >= elf1Low) {

			overlappingPairs++
		}
	}

	return overlappingPairs
}

func ranges(input []byte) (elf1, elf2 []byte) {
	ranges := bytes.Split(input, []byte{','})

	if l := len(ranges); l != 2 {
		panic(fmt.Errorf("read line with %d range(s) instead of 2", l))
	}

	return ranges[0], ranges[1]
}

func rangeBoundaries(input []byte) (low, high int) {
	boundaries := bytes.Split(input, []byte{'-'})

	if l := len(boundaries); l != 2 {
		panic(fmt.Errorf("read range with %d boundaries(s) instead of 2", l))
	}

	var err error

	low, err = strconv.Atoi(string(boundaries[0]))
	if err != nil {
		panic(fmt.Errorf("parsing %v to integer: %w", low, err))
	}

	high, err = strconv.Atoi(string(boundaries[1]))
	if err != nil {
		panic(fmt.Errorf("parsing %v to integer: %w", low, err))
	}
	if low > high {
		panic(fmt.Errorf("low boundary %d is greater than high boundary %d", low, high))
	}

	return low, high
}
