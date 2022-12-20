package day13

import (
	"bufio"
	"encoding/json"
	"fmt"
	"math"

	"adventofcode2022/fs"
)

func Puzzle1(fs fs.FS, inputFile string) int {
	f, err := fs.Open(inputFile)
	if err != nil {
		panic(fmt.Errorf("opening input file: %w", err))
	}

	s := bufio.NewScanner(f)

	packetPairs := make([][2][]any, 1)

	currentLine := -1

	for s.Scan() {
		currentLine++

		if len(s.Bytes()) == 0 {
			packetPairs = append(packetPairs, [2][]any{})
			continue
		}

		var packet []any
		if err := json.Unmarshal(s.Bytes(), &packet); err != nil {
			panic(fmt.Errorf("deserializing input: %w", err))
		}

		packetPairs[len(packetPairs)-1][currentLine%3] = packet
	}

	if err := s.Err(); err != nil {
		panic(fmt.Errorf("while scanning input file: %w", err))
	}

	sumOrdered := 0

	for i, pair := range packetPairs {
		if isPairInOrder(pair[0], pair[1]) {
			sumOrdered += i + 1
		}
	}

	return sumOrdered
}

func Puzzle2(fs fs.FS, inputFile string) int {
	f, err := fs.Open(inputFile)
	if err != nil {
		panic(fmt.Errorf("opening input file: %w", err))
	}

	s := bufio.NewScanner(f)

	dividerPacket1 := []any{[]any{float64(2)}}
	dividerPacket2 := []any{[]any{float64(6)}}
	divider1Idx, divider2Idx := 1, 2

	for s.Scan() {
		if len(s.Bytes()) == 0 {
			continue
		}

		var packet []any
		if err := json.Unmarshal(s.Bytes(), &packet); err != nil {
			panic(fmt.Errorf("deserializing input: %w", err))
		}

		if isPairInOrder(packet, dividerPacket1) {
			divider1Idx++
		}
		if isPairInOrder(packet, dividerPacket2) {
			divider2Idx++
		}
	}

	if err := s.Err(); err != nil {
		panic(fmt.Errorf("while scanning input file: %w", err))
	}

	return divider1Idx * divider2Idx
}

func isPairInOrder(left, right any) bool {
	mustCompareNext, inOrder := valuesInOrder(left, right)
	if mustCompareNext {
		panic("Comparison ended on an equality")
	}
	return inOrder
}

func valuesInOrder(left, right any) (mustCompareNext, isInOrder bool) {
	switch left := left.(type) {
	case float64:
		switch right := right.(type) {
		case float64:
			return left == right, left < right
		case []any:
			return valuesInOrder([]any{left}, right)
		}

	case []any:
		switch right := right.(type) {
		case float64:
			return valuesInOrder(left, []any{right})
		case []any:
			maxCompare := int(math.Min(float64(len(left)), float64(len(right))))

			for i := 0; i < maxCompare; i++ {
				mustCompareNext, isInOrder = valuesInOrder(left[i], right[i])
				if mustCompareNext {
					continue
				}
				return false, isInOrder
			}

			// No decision could be taken after comparing values
			// one by one, proceed to comparing packet lengths.
			switch {
			case len(left) < len(right):
				return false, true
			case len(left) > len(right):
				return false, false
			default:
				return true, false
			}
		}
	}

	panic(fmt.Errorf("Unsupported types: %T/%T", left, right))
}
