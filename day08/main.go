package day08

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

	s := bufio.NewScanner(f)

	var treePatch [][]int

	for s.Scan() {
		if treePatch == nil {
			treePatch = make([][]int, len(s.Bytes()))
		}

		for i, digit := range s.Bytes() {
			treeHeight, err := strconv.Atoi(string([]byte{digit}))
			if err != nil {
				panic(fmt.Errorf("parsing %v to integer: %w", []byte{digit}, err))
			}
			treePatch[i] = append(treePatch[i], treeHeight)
		}
	}

	if err := s.Err(); err != nil {
		panic(fmt.Errorf("while scanning input file: %w", err))
	}

	visibleTrees := make([][]int, len(treePatch))
	for i := range visibleTrees {
		visibleTrees[i] = make([]int, len(treePatch[0]))
	}

	var numVisibleTrees int

	/*
		Example

		[[3 2 6 3 3] [0 5 5 3 5] [3 5 3 5 3] [7 1 3 4 9] [3 2 2 9 0]]
	*/

	columns, rows := len(treePatch), len(treePatch[0])

	/*
		1) Traverse each columns from top to bottom and back

		.
		30373
		25512
		65332
		33549
		35390
		.
	*/
	for x := 0; x < columns; x++ {
		// NOTE: possible time optimization?
		// During first traversal, store data (what?) to potentially avoid the reverse traversal?

		highestFromEdge := -1
		for y := 0; y < rows; y++ {
			if treePatch[x][y] > highestFromEdge {
				highestFromEdge = treePatch[x][y]
				if visibleTrees[x][y] == 0 {
					numVisibleTrees++
				}
				visibleTrees[x][y]++
			}
		}

		highestFromEdge = -1
		for y := rows - 1; y >= 0; y-- {
			if treePatch[x][y] > highestFromEdge {
				highestFromEdge = treePatch[x][y]
				if visibleTrees[x][y] == 0 {
					numVisibleTrees++
				}
				visibleTrees[x][y]++
			}
		}
	}

	/*
		2) Traverse each row from left to right and back

		.30373.
		 25512
		 65332
		 33549
		 35390
	*/
	for y := 0; y < rows; y++ {
		highestFromEdge := -1
		for x := 0; x < columns; x++ {
			if treePatch[x][y] > highestFromEdge {
				highestFromEdge = treePatch[x][y]
				if visibleTrees[x][y] == 0 {
					numVisibleTrees++
				}
				visibleTrees[x][y]++
			}
		}

		highestFromEdge = -1
		for x := columns - 1; x >= 0; x-- {
			if treePatch[x][y] > highestFromEdge {
				highestFromEdge = treePatch[x][y]
				if visibleTrees[x][y] == 0 {
					numVisibleTrees++
				}
				visibleTrees[x][y]++
			}
		}
	}

	return numVisibleTrees
}

func Puzzle2(fs fs.FS, inputFile string) int {
	f, err := fs.Open(inputFile)
	if err != nil {
		panic(fmt.Errorf("opening input file: %w", err))
	}

	s := bufio.NewScanner(f)

	var treePatch [][]int

	for s.Scan() {
		if treePatch == nil {
			treePatch = make([][]int, len(s.Bytes()))
		}

		for i, digit := range s.Bytes() {
			treeHeight, err := strconv.Atoi(string([]byte{digit}))
			if err != nil {
				panic(fmt.Errorf("parsing %v to integer: %w", []byte{digit}, err))
			}
			treePatch[i] = append(treePatch[i], treeHeight)
		}
	}

	if err := s.Err(); err != nil {
		panic(fmt.Errorf("while scanning input file: %w", err))
	}

	var highestScenicScore int

	columns, rows := len(treePatch), len(treePatch[0])

	for x := 1; x < columns-1; x++ {
		for y := 1; y < rows-1; y++ {
			var currentScenicScore int
			currentTreeHeight := treePatch[x][y]

			// left
			visibleTreesCurrentDirection := 0
			for i := x - 1; i >= 0; i-- {
				if treePatch[i][y] >= currentTreeHeight {
					i = -1
				}
				visibleTreesCurrentDirection++
			}
			currentScenicScore = visibleTreesCurrentDirection

			// right
			visibleTreesCurrentDirection = 0
			for i := x + 1; i < columns; i++ {
				if treePatch[i][y] >= currentTreeHeight {
					i = columns
				}
				visibleTreesCurrentDirection++
			}
			currentScenicScore *= visibleTreesCurrentDirection

			// up
			visibleTreesCurrentDirection = 0
			for j := y - 1; j >= 0; j-- {
				if treePatch[x][j] >= currentTreeHeight {
					j = -1
				}
				visibleTreesCurrentDirection++
			}
			currentScenicScore *= visibleTreesCurrentDirection

			// down
			visibleTreesCurrentDirection = 0
			for j := y + 1; j < rows; j++ {
				if treePatch[x][j] >= currentTreeHeight {
					j = rows
				}
				visibleTreesCurrentDirection++
			}
			currentScenicScore *= visibleTreesCurrentDirection

			if currentScenicScore > highestScenicScore {
				highestScenicScore = currentScenicScore
			}
		}
	}

	return highestScenicScore
}
