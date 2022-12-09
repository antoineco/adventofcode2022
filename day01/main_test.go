package day01_test

import (
	"testing"

	"adventofcode2022/day01"
	"adventofcode2022/fs"
)

func TestPuzzle1(t *testing.T) {
	testCases := []struct {
		desc   string
		input  []byte
		expect int
	}{
		{
			desc: "Provided example",
			input: []byte{
				'1', '0', '0', '0', '\n',
				'2', '0', '0', '0', '\n',
				'3', '0', '0', '0', '\n',
				'\n',
				'4', '0', '0', '0', '\n',
				'\n',
				'5', '0', '0', '0', '\n',
				'6', '0', '0', '0', '\n',
				'\n',
				'7', '0', '0', '0', '\n',
				'8', '0', '0', '0', '\n',
				'9', '0', '0', '0', '\n',
				'\n',
				'1', '0', '0', '0', '0', '\n',
			},
			expect: 24_000,
		},
		{
			desc: "Multiple elves carry an equal amount of calories",
			input: []byte{
				'1', '\n',
				'2', '\n',
				'3', '\n',
				'\n',
				'1', '\n',
				'2', '\n',
				'3', '\n',
				'\n',
				'1', '\n',
				'2', '\n',
			},
			expect: 6,
		},
		{
			desc: "First elf carries the most calories",
			input: []byte{
				'1', '\n',
				'2', '\n',
				'3', '\n',
				'\n',
				'1', '\n',
				'2', '\n',
				'\n',
				'1', '\n',
			},
			expect: 6,
		},
		{
			desc: "Last elf carries the most calories",
			input: []byte{
				'1', '\n',
				'2', '\n',
				'\n',
				'1', '\n',
				'\n',
				'1', '\n',
				'2', '\n',
				'3', '\n',
			},
			expect: 6,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			fs := fs.NewMemFS()
			if err := fs.CreateFile("input", tc.input); err != nil {
				t.Fatal("Error creating in-memory test file:", err)
			}

			out := day01.Puzzle1(fs, "input")

			if out != tc.expect {
				t.Errorf("Expected %d, got %d", tc.expect, out)
			}
		})
	}
}

func TestPuzzle2(t *testing.T) {
	testCases := []struct {
		desc   string
		input  []byte
		expect int
	}{
		{
			desc: "Provided example",
			input: []byte{
				'1', '0', '0', '0', '\n',
				'2', '0', '0', '0', '\n',
				'3', '0', '0', '0', '\n',
				'\n',
				'4', '0', '0', '0', '\n',
				'\n',
				'5', '0', '0', '0', '\n',
				'6', '0', '0', '0', '\n',
				'\n',
				'7', '0', '0', '0', '\n',
				'8', '0', '0', '0', '\n',
				'9', '0', '0', '0', '\n',
				'\n',
				'1', '0', '0', '0', '0', '\n',
			},
			expect: 45_000,
		},
		{
			desc: "Multiple elves carry an equal amount of calories",
			input: []byte{
				'4', '\n',
				'\n',
				'5', '\n',
				'\n',
				'3', '\n',
				'\n',
				'4', '\n',
				'\n',
				'2', '\n',
				'\n',
				'1', '\n',
			},
			expect: 13,
		},
		{
			desc: "Last elf in top 3",
			input: []byte{
				'4', '\n',
				'\n',
				'1', '\n',
				'\n',
				'3', '\n',
				'\n',
				'6', '\n',
				'\n',
				'2', '\n',
				'\n',
				'5', '\n',
			},
			expect: 15,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			fs := fs.NewMemFS()
			if err := fs.CreateFile("input", tc.input); err != nil {
				t.Fatal("Error creating in-memory test file:", err)
			}

			out := day01.Puzzle2(fs, "input")

			if out != tc.expect {
				t.Errorf("Expected %d, got %d", tc.expect, out)
			}
		})
	}
}
