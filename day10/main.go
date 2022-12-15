package day10

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"

	"adventofcode2022/fs"
)

func Puzzle1(fs fs.FS, inputFile string) int {
	f, err := fs.Open(inputFile)
	if err != nil {
		panic(fmt.Errorf("opening input file: %w", err))
	}

	s := bufio.NewScanner(f)

	var signalStrength int

	const maxCycle = 220

	signalStrengthCycles := make(set, 6)
	for i := 20; i <= maxCycle; i += 40 {
		signalStrengthCycles.add(i)
	}

	registerVal := 1
	cycle := 0
	operations := make(queue, 0)

	for s.Scan() {
		cycle++

		if signalStrengthCycles.has(cycle) {
			signalStrength += cycle * registerVal
		}

		if cycle == maxCycle {
			break
		}

		instr, val := readInstruction(s.Bytes())

		operations.push(0) // delaying zero, each op takes at least one cycle
		if instr == addx {
			operations.push(val)
		}

		registerVal += operations.pop()
	}

	if err := s.Err(); err != nil {
		panic(fmt.Errorf("while scanning input file: %w", err))
	}

	if cycle < maxCycle {
		for i, remainingOperations := 0, len(operations); i < remainingOperations; i++ {
			cycle++

			if signalStrengthCycles.has(cycle) {
				signalStrength += cycle * registerVal
			}

			if cycle == maxCycle {
				break
			}

			registerVal += operations.pop()
		}
	}

	// simulate noop cycles if the input had less than maxCycle cycles
	if cycle < maxCycle {
		for i := 20; i <= maxCycle; i += 40 {
			signalStrengthCycles.add(i)
			if cycle < i {
				signalStrength += i * registerVal
			}
		}
	}

	return signalStrength
}

func Puzzle2(fs fs.FS, inputFile string) string {
	f, err := fs.Open(inputFile)
	if err != nil {
		panic(fmt.Errorf("opening input file: %w", err))
	}

	s := bufio.NewScanner(f)

	const maxCycle = 240
	const pixelsPerRow = 40

	var crt [maxCycle / pixelsPerRow][pixelsPerRow]byte

	registerVal := 1
	cycle := 0
	operations := make(queue, 0)

	currentRow := 0

	for s.Scan() {
		cycle++

		instr, val := readInstruction(s.Bytes())

		operations.push(0) // delaying zero, each op takes at least one cycle
		if instr == addx {
			operations.push(val)
		}

		currentColumn := (cycle - 1) % pixelsPerRow
		var drawPixel byte = '.'
		if registerVal-1 <= currentColumn && currentColumn <= registerVal+1 {
			drawPixel = '#'
		}
		crt[currentRow][currentColumn] = drawPixel

		if cycle == maxCycle {
			break
		}

		if cycle%pixelsPerRow == 0 {
			currentRow++
		}

		registerVal += operations.pop()
	}

	if err := s.Err(); err != nil {
		panic(fmt.Errorf("while scanning input file: %w", err))
	}

	if cycle < maxCycle {
		for i, remainingOperations := 0, len(operations); i < remainingOperations; i++ {
			cycle++

			currentColumn := (cycle - 1) % pixelsPerRow
			var drawPixel byte = '.'
			if registerVal-1 <= currentColumn && currentColumn <= registerVal+1 {
				drawPixel = '#'
			}
			crt[currentRow][currentColumn] = drawPixel

			if cycle == maxCycle {
				break
			}

			if cycle%pixelsPerRow == 0 {
				currentRow++
			}

			registerVal += operations.pop()
		}
	}

	var crtDisplay strings.Builder

	for row := range crt {
		for _, pixel := range crt[row] {
			crtDisplay.WriteByte(pixel)
		}
		if row < len(crt)-1 {
			crtDisplay.WriteByte('\n')
		}
	}

	return crtDisplay.String()
}

func readInstruction(input []byte) (instruction, int) {
	if input[0] == 'n' {
		return noop, 0
	}

	v, err := strconv.Atoi(string(input[len("addx")+1:]))
	if err != nil {
		panic(fmt.Errorf("parsing %v to integer: %w", input[len("addx")+1:], err))
	}

	return addx, v
}

type instruction uint8

const (
	noop instruction = iota
	addx
)

type queue []int

func (q *queue) push(v int) {
	*q = append(*q, v)
}

func (q *queue) pop() int {
	if len(*q) == 0 {
		return 0
	}

	v := (*q)[0]

	*q = (*q)[1:]

	return v
}

type set map[int]struct{}

func (s *set) has(v int) bool {
	_, exists := (*s)[v]
	return exists
}

func (s *set) add(v int) {
	if !s.has(v) {
		(*s)[v] = struct{}{}
	}
}
