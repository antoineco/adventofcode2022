package main

import (
	"fmt"

	"adventofcode2022/day01"
	"adventofcode2022/fs"
)

func main() {
	var fs *fs.OSFS

	fmt.Println("Day 01 / Puzzle 1:", day01.Puzzle1(fs, "inputs/d01.txt"))
	fmt.Println("Day 01 / Puzzle 2:", day01.Puzzle2(fs, "inputs/d01.txt"))
}
