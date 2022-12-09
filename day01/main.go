package day01

import (
	"bufio"
	"fmt"
	"strconv"

	"adventofcode2022/fs"
)

func Puzzle1(fs fs.FS, inputFile string) int {
	f, err := fs.Open(inputFile)
	if err != nil {
		panic(fmt.Errorf("opening input file: %w", err))
	}

	maxCals := -1
	curCals := 0

	s := bufio.NewScanner(f)

	for s.Scan() {
		line := s.Text()
		if line == "" {
			curCals = 0
			continue
		}

		cals, err := strconv.Atoi(line)
		if err != nil {
			panic(fmt.Errorf("parsing value to integer: %w", err))
		}

		curCals += cals

		if curCals > maxCals {
			maxCals = curCals
		}
	}

	if err := s.Err(); err != nil {
		panic(fmt.Errorf("while scanning input file: %w", err))
	}

	if maxCals < 0 {
		panic(fmt.Errorf("negative calories count"))
	}

	return maxCals
}

func Puzzle2(fs fs.FS, inputFile string) int {
	f, err := fs.Open(inputFile)
	if err != nil {
		panic(fmt.Errorf("opening input file: %w", err))
	}

	maxCalsMin := -1
	maxCalsMid := -1
	maxCalsMax := -1

	curCals := 0

	s := bufio.NewScanner(f)

	for s.Scan() {
		line := s.Text()
		if line == "" {
			switch {
			case curCals > maxCalsMax:
				maxCalsMin, maxCalsMid, maxCalsMax = maxCalsMid, maxCalsMax, curCals
			case curCals > maxCalsMid:
				maxCalsMin, maxCalsMid = maxCalsMid, curCals
			case curCals > maxCalsMin:
				maxCalsMin = curCals
			}

			curCals = 0
			continue
		}

		cals, err := strconv.Atoi(line)
		if err != nil {
			panic(fmt.Errorf("parsing value to integer: %w", err))
		}

		curCals += cals
	}

	if err := s.Err(); err != nil {
		panic(fmt.Errorf("while scanning input file: %w", err))
	}

	// account for last elf (EOF)
	switch {
	case curCals > maxCalsMax:
		maxCalsMin, maxCalsMid, maxCalsMax = maxCalsMid, maxCalsMax, curCals
	case curCals > maxCalsMid:
		maxCalsMin, maxCalsMid = maxCalsMid, curCals
	case curCals > maxCalsMin:
		maxCalsMin = curCals
	}

	if maxCalsMin < 0 || maxCalsMid < 0 || maxCalsMax < 0 {
		panic(fmt.Errorf("negative calories count"))
	}

	return maxCalsMin + maxCalsMid + maxCalsMax
}
