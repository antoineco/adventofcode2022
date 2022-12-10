package day03

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

	sumPriorities := 0

	s := bufio.NewScanner(f)

	for s.Scan() {
		items := s.Bytes()

		foundTypes := make(set, len(items)/2)

		for _, t := range items[:len(items)/2] {
			foundTypes.add(t)
		}

		for _, t := range items[len(items)/2:] {
			if foundTypes.has(t) {
				sumPriorities += itemPriority(t)
				break
			}
		}
	}

	return sumPriorities
}

func Puzzle2(fs fs.FS, inputFile string) int {
	f, err := fs.Open(inputFile)
	if err != nil {
		panic(fmt.Errorf("opening input file: %w", err))
	}

	const groupSize = 3
	groupItems := make([][]byte, groupSize)

	sumPriorities := 0

	s := bufio.NewScanner(f)

	lineNr := 0

	for s.Scan() {
		lineNr++

		idx := lineNr%groupSize - 1
		if idx == -1 {
			idx = groupSize - 1
		}

		groupItems[idx] = s.Bytes()

		if idx == groupSize-1 {
			sumPriorities += itemPriority(findBadge(groupItems...))
		}
	}

	return sumPriorities
}

func itemPriority(typ byte) int {
	if 'a' <= typ && typ <= 'z' {
		return int(typ%'a') + 1
	}
	if 'A' <= typ && typ <= 'Z' {
		return int(typ%'A') + 27
	}
	return -1
}

func findBadge(itemLists ...[]byte) byte {
	foundTypes := make([]set, len(itemLists))

	for i := range foundTypes {
		if foundTypes[i] == nil {
			foundTypes[i] = make(set, len(itemLists[i]))
		}
		for _, item := range itemLists[i] {
			foundTypes[i].add(item)
		}
	}

	var foundBadge bool

	for typ := range foundTypes[0] {
		foundBadge = true
		for _, ft := range foundTypes[1:] {
			if !ft.has(typ) {
				foundBadge = false
				break
			}
		}
		if foundBadge {
			return typ
		}
	}

	return 0
}

type set map[byte]struct{}

func (s *set) has(v byte) bool {
	_, exists := (*s)[v]
	return exists
}

func (s *set) add(v byte) {
	if !s.has(v) {
		(*s)[v] = struct{}{}
	}
}
