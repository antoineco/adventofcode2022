package day11

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

	s := bufio.NewScanner(f)

	var monkeys []*monkey

	var currentMonkey int

	for s.Scan() {
		switch {
		case len(s.Bytes()) == 0:
			continue

		case isMonkey(s.Bytes()):
			currentMonkey = readMonkeyNr(s.Bytes())

			n := (currentMonkey + 1) - len(monkeys)
			for i := 0; i < n; i++ {
				monkeys = append(monkeys, (*monkey)(nil))
			}
			monkeys[currentMonkey] = &monkey{}

		case isStartingItems(s.Bytes()):
			monkeys[currentMonkey].items = readItems(s.Bytes())

		case isOperation(s.Bytes()):
			monkeys[currentMonkey].calculateWorry = readOperation(s.Bytes())

		case isTest(s.Bytes()):
			divVal := readTestDivValue(s.Bytes())
			s.Scan()
			monkeyTrue := readTestCondMonkey(s.Bytes())
			s.Scan()
			monkeyFalse := readTestCondMonkey(s.Bytes())

			monkeys[currentMonkey].decideRecipient = makeRecipientFn(divVal, monkeyTrue, monkeyFalse)
		}
	}

	if err := s.Err(); err != nil {
		panic(fmt.Errorf("while scanning input file: %w", err))
	}

	const numRounds = 20

	var maxInspectedLow, maxInspectedHigh int
	var maxInspectedHighMonkey int

	for r := 0; r < numRounds; r++ {
		for n, m := range monkeys {
			numItems := len(m.items)

			m.itemsInspected += numItems
			switch {
			case m.itemsInspected > maxInspectedHigh:
				if maxInspectedHighMonkey != n {
					maxInspectedLow, maxInspectedHigh, maxInspectedHighMonkey = maxInspectedHigh, m.itemsInspected, n
				} else {
					maxInspectedHigh, maxInspectedHighMonkey = m.itemsInspected, n
				}
			case m.itemsInspected > maxInspectedLow:
				maxInspectedLow = m.itemsInspected
			}

			for i := 0; i < numItems; i++ {
				itemWorryLvl := m.calculateWorry(m.inspectNext()) / 3
				monkeys[m.decideRecipient(itemWorryLvl)].receiveItem(itemWorryLvl)
			}
		}
	}

	return maxInspectedLow * maxInspectedHigh
}

// I needed to search for a hint for that part.
//
// Problem: Without trickin the numbers, worry levels overflow the integer capacity.
//
// Solution: The trick is to find a common multiple of all monkeys' divisors and perform a modulo to this value for each
// inspected item. This way, worry level values remain low, yet modulo operations inside recipientFuncs remain valid.
//
// Note: All divisors are prime numbers (in the provided example: 23, 19, 13, 17). The common multiple is the product of
// all divisors.
func Puzzle2(fs fs.FS, inputFile string) int {
	f, err := fs.Open(inputFile)
	if err != nil {
		panic(fmt.Errorf("opening input file: %w", err))
	}

	s := bufio.NewScanner(f)

	var monkeys []*monkey

	var currentMonkey int
	commonMultiple := 1

	for s.Scan() {
		switch {
		case len(s.Bytes()) == 0:
			continue

		case isMonkey(s.Bytes()):
			currentMonkey = readMonkeyNr(s.Bytes())

			n := (currentMonkey + 1) - len(monkeys)
			for i := 0; i < n; i++ {
				monkeys = append(monkeys, (*monkey)(nil))
			}
			monkeys[currentMonkey] = &monkey{}

		case isStartingItems(s.Bytes()):
			monkeys[currentMonkey].items = readItems(s.Bytes())

		case isOperation(s.Bytes()):
			monkeys[currentMonkey].calculateWorry = readOperation(s.Bytes())

		case isTest(s.Bytes()):
			divVal := readTestDivValue(s.Bytes())
			s.Scan()
			monkeyTrue := readTestCondMonkey(s.Bytes())
			s.Scan()
			monkeyFalse := readTestCondMonkey(s.Bytes())

			monkeys[currentMonkey].decideRecipient = makeRecipientFn(divVal, monkeyTrue, monkeyFalse)

			commonMultiple *= divVal
		}
	}

	if err := s.Err(); err != nil {
		panic(fmt.Errorf("while scanning input file: %w", err))
	}

	const numRounds = 10_000

	var maxInspectedLow, maxInspectedHigh int
	var maxInspectedHighMonkey int

	for r := 0; r < numRounds; r++ {
		for n, m := range monkeys {
			numItems := len(m.items)

			m.itemsInspected += numItems
			switch {
			case m.itemsInspected > maxInspectedHigh:
				if maxInspectedHighMonkey != n {
					maxInspectedLow, maxInspectedHigh, maxInspectedHighMonkey = maxInspectedHigh, m.itemsInspected, n
				} else {
					maxInspectedHigh, maxInspectedHighMonkey = m.itemsInspected, n
				}
			case m.itemsInspected > maxInspectedLow:
				maxInspectedLow = m.itemsInspected
			}

			for i := 0; i < numItems; i++ {
				itemWorryLvl := m.calculateWorry(m.inspectNext()) % commonMultiple
				monkeys[m.decideRecipient(itemWorryLvl)].receiveItem(itemWorryLvl)
			}
		}
	}

	return maxInspectedLow * maxInspectedHigh
}

type monkey struct {
	items           []int
	calculateWorry  worryCalcFunc
	decideRecipient recipientFunc

	itemsInspected int
}

func (m *monkey) inspectNext() (item int) {
	item, m.items = m.items[0], m.items[1:]
	return item
}

func (m *monkey) receiveItem(item int) {
	m.items = append(m.items, item)
}

func isMonkey(input []byte) bool {
	return input[0] == 'M'
}

func isStartingItems(input []byte) bool {
	return input[0] == ' ' && input[2] == 'S'
}

func isOperation(input []byte) bool {
	return input[0] == ' ' && input[2] == 'O'
}

func isTest(input []byte) bool {
	return input[0] == ' ' && input[2] == 'T'
}

func readMonkeyNr(input []byte) int {
	monkeyNumPos := len("Monkey") + 1 // "Monkey <monkey>:"

	n, err := strconv.Atoi(string(input[monkeyNumPos : len(input)-1]))
	if err != nil {
		panic(fmt.Errorf("parsing %v to integer: %w", input[monkeyNumPos:len(input)-1], err))
	}

	return n
}

func readItems(input []byte) []int {
	itemsPos := len("Starting items:") + 3 // "  Starting items: <items>"

	subs := bytes.Split(input[itemsPos:], []byte{',', ' '})

	items := make([]int, 0, len(subs))

	for _, sub := range subs {
		v, err := strconv.Atoi(string(sub))
		if err != nil {
			panic(fmt.Errorf("parsing %v to integer: %w", sub, err))
		}
		items = append(items, v)
	}

	return items
}

type worryCalcFunc func(old int) (new int)

func readOperation(input []byte) worryCalcFunc {
	opPos := len("Operation: new =") + 3 // "  Operation: new = <operation>"

	switch string(input[opPos:]) {
	case "old + old":
		return func(old int) int { return old + old }
	case "old * old":
		return func(old int) int { return old * old }
	}

	fields := bytes.Fields(input[opPos:])

	n := fields[2]
	y, err := strconv.Atoi(string(n))
	if err != nil {
		panic(fmt.Errorf("parsing %v to integer: %w", n, err))
	}

	switch opSign := fields[1][0]; opSign {
	case '+':
		return func(old int) int { return old + y }
	case '*':
		return func(old int) int { return old * y }
	default:
		panic(fmt.Errorf("unrecognized operation %q (%d)", string(opSign), opSign))
	}
}

func readTestDivValue(input []byte) int {
	divValPos := len("Test: divisible by") + 3 //"  Test: divisible by <val>"

	n, err := strconv.Atoi(string(input[divValPos:]))
	if err != nil {
		panic(fmt.Errorf("parsing %v to integer: %w", input[divValPos:], err))
	}

	return n
}

func readTestCondMonkey(input []byte) (monkey int) {
	const condCharPos = len("If") + 5 // "....If.<condition>:.throw to monkey <monkey>"

	condition := "true"
	if input[condCharPos] == 'f' {
		condition = "false"
	}

	preTxtLen := condCharPos + len(condition) + len(": throw to monkey ")

	n, err := strconv.Atoi(string(input[preTxtLen:]))
	if err != nil {
		panic(fmt.Errorf("parsing %v to integer: %w", input[preTxtLen:], err))
	}

	return n
}

type recipientFunc func(worryLvl int) (monkey int)

func makeRecipientFn(divVal, monkeyTrue, monkeyFalse int) recipientFunc {
	return func(worryLvl int) int {
		if worryLvl%divVal == 0 {
			return monkeyTrue
		}
		return monkeyFalse
	}
}
