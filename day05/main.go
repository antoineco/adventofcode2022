package day05

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"

	"adventofcode2022/fs"
)

func Puzzle1(fs fs.FS, inputFile string) string {
	f, err := fs.Open(inputFile)
	if err != nil {
		panic(fmt.Errorf("opening input file: %w", err))
	}

	s := bufio.NewScanner(f)

	s.Scan()
	if err := s.Err(); err != nil {
		panic(fmt.Errorf("while scanning input file: %w", err))
	}

	numStacks := numStacks(s.Bytes())
	if numStacks == 0 {
		return ""
	}

	stacks := make([]*stack, numStacks)

	readCrates(s.Bytes(), stacks)

	for s.Scan() {
		if len(s.Bytes()) > 1 && s.Bytes()[1] == '1' {
			break
		}

		readCrates(s.Bytes(), stacks)
	}

	if err := s.Err(); err != nil {
		panic(fmt.Errorf("while scanning input file: %w", err))
	}

	for s.Scan() {
		if len(s.Bytes()) == 0 {
			continue
		}

		numCrates, srcStack, dstStack := readMove(s.Bytes())

		for i := 0; i < numCrates; i++ {
			stacks[dstStack].push(stacks[srcStack].pop())
		}
	}
	if err := s.Err(); err != nil {
		panic(fmt.Errorf("while scanning input file: %w", err))
	}

	var topCrates []byte

	for _, s := range stacks {
		if topCrate := s.pop(); topCrate != 0 {
			topCrates = append(topCrates, topCrate)
		}
	}

	return string(topCrates)
}

func Puzzle2(fs fs.FS, inputFile string) string {
	f, err := fs.Open(inputFile)
	if err != nil {
		panic(fmt.Errorf("opening input file: %w", err))
	}

	s := bufio.NewScanner(f)

	s.Scan()
	if err := s.Err(); err != nil {
		panic(fmt.Errorf("while scanning input file: %w", err))
	}

	numStacks := numStacks(s.Bytes())
	if numStacks == 0 {
		return ""
	}

	stacks := make([]*stack, numStacks)

	readCrates(s.Bytes(), stacks)

	for s.Scan() {
		if len(s.Bytes()) > 1 && s.Bytes()[1] == '1' {
			break
		}

		readCrates(s.Bytes(), stacks)
	}

	if err := s.Err(); err != nil {
		panic(fmt.Errorf("while scanning input file: %w", err))
	}

	for s.Scan() {
		if len(s.Bytes()) == 0 {
			continue
		}

		numCrates, srcStack, dstStack := readMove(s.Bytes())

		switch numCrates {
		case 1:
			stacks[dstStack].push(stacks[srcStack].pop())
		default:
			queue := make([]byte, numCrates)

			for i := 0; i < numCrates; i++ {
				queue[i] = stacks[srcStack].pop()
			}
			for i := numCrates - 1; i >= 0; i-- {
				stacks[dstStack].push(queue[i])
			}
		}
	}
	if err := s.Err(); err != nil {
		panic(fmt.Errorf("while scanning input file: %w", err))
	}

	var topCrates []byte

	for _, s := range stacks {
		if topCrate := s.pop(); topCrate != 0 {
			topCrates = append(topCrates, topCrate)
		}
	}

	return string(topCrates)
}

const (
	crateSize = 3
	crateSide = '['
)

func numStacks(input []byte) int {
	switch l := len(input); l {
	case 0:
		return 0
	case crateSize:
		return 1
	default:
		return 1 + (l-crateSize)/(crateSize+1)
	}
}

func readCrates(input []byte, stacks []*stack) {
	if stacks[0] == nil {
		stacks[0] = &stack{}
	}
	if input[0] == crateSide {
		stacks[0].append(input[1])
	}

	for i := crateSize + 1; i < len(input); i += crateSize + 1 {
		stackIdx := i / (crateSize + 1)

		if stacks[stackIdx] == nil {
			stacks[stackIdx] = &stack{}
		}

		if input[i] == crateSide {
			stacks[stackIdx].append(input[i+1])
		}
	}
}

func readMove(input []byte) (numCrates, srcStack, dstStack int) {
	var values [3]int

	for i, curValIdx := 0, 0; i < len(input); i++ {
		if input[i] == ' ' && '0' <= input[i+1] && input[i+1] <= '9' {
			i++
			var val []byte

			for ; i < len(input); i++ {
				if input[i] == ' ' {
					break
				}
				val = append(val, input[i])
			}

			v, err := strconv.Atoi(string(val))
			if err != nil {
				panic(fmt.Errorf("parsing %v to integer: %w", v, err))
			}

			values[curValIdx] = v
			curValIdx++
		}
	}

	return values[0], values[1] - 1, values[2] - 1
}

// Liberal stack implementation with "append" method for easy initial
// construction in reverse order while parsing the input file.
type stack struct {
	head *stackItem

	// only relevant during initial constuction using "append"
	tail *stackItem
}

type stackItem struct {
	val  byte
	next *stackItem
}

func (s *stack) push(val byte) {
	s.head = &stackItem{
		val:  val,
		next: s.head,
	}
}

func (s *stack) pop() byte {
	if s.head == nil {
		return 0
	}

	var val byte
	val, s.head = s.head.val, s.head.next
	return val
}

func (s *stack) String() string {
	if s.head == nil {
		return ""
	}

	var b strings.Builder

	for item := s.head; ; item = item.next {
		if item == nil {
			break
		}

		b.WriteByte(item.val)
	}

	return b.String()
}

// Only relevant during initial constuction.
func (s *stack) append(val byte) {
	tail := &stackItem{
		val: val,
	}

	if s.head == nil {
		s.tail = tail
		s.head = tail
		return
	}

	s.tail.next = tail
	s.tail = tail
}
