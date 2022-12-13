package main

import (
	"fmt"

	"adventofcode2022/day01"
	"adventofcode2022/day02"
	"adventofcode2022/day03"
	"adventofcode2022/day04"
	"adventofcode2022/day05"
	"adventofcode2022/day06"
	"adventofcode2022/day07"
	"adventofcode2022/day08"
	"adventofcode2022/fs"
)

func main() {
	var fs *fs.OSFS

	fmt.Println("Day 01 / Puzzle 1:", day01.Puzzle1(fs, "inputs/d01.txt"))
	fmt.Println("Day 01 / Puzzle 2:", day01.Puzzle2(fs, "inputs/d01.txt"))

	fmt.Println("Day 02 / Puzzle 1:", day02.Puzzle1(fs, "inputs/d02.txt"))
	fmt.Println("Day 02 / Puzzle 2:", day02.Puzzle2(fs, "inputs/d02.txt"))

	fmt.Println("Day 03 / Puzzle 1:", day03.Puzzle1(fs, "inputs/d03.txt"))
	fmt.Println("Day 03 / Puzzle 2:", day03.Puzzle2(fs, "inputs/d03.txt"))

	fmt.Println("Day 04 / Puzzle 1:", day04.Puzzle1(fs, "inputs/d04.txt"))
	fmt.Println("Day 04 / Puzzle 2:", day04.Puzzle2(fs, "inputs/d04.txt"))

	fmt.Println("Day 05 / Puzzle 1:", day05.Puzzle1(fs, "inputs/d05.txt"))
	fmt.Println("Day 05 / Puzzle 2:", day05.Puzzle2(fs, "inputs/d05.txt"))

	fmt.Println("Day 06 / Puzzle 1:", day06.Puzzle1(fs, "inputs/d06.txt"))
	fmt.Println("Day 06 / Puzzle 2:", day06.Puzzle2(fs, "inputs/d06.txt"))

	fmt.Println("Day 07 / Puzzle 1:", day07.Puzzle1(fs, "inputs/d07.txt"))
	fmt.Println("Day 07 / Puzzle 2:", day07.Puzzle2(fs, "inputs/d07.txt"))

	fmt.Println("Day 08 / Puzzle 1:", day08.Puzzle1(fs, "inputs/d08.txt"))
	fmt.Println("Day 08 / Puzzle 2:", day08.Puzzle2(fs, "inputs/d08.txt"))
}
