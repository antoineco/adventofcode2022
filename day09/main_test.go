package day09_test

import (
	"testing"

	"adventofcode2022/day09"
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
				'R', ' ', '4', '\n',
				'U', ' ', '4', '\n',
				'L', ' ', '3', '\n',
				'D', ' ', '1', '\n',
				'R', ' ', '4', '\n',
				'D', ' ', '1', '\n',
				'L', ' ', '5', '\n',
				'R', ' ', '2', '\n',
			},
			expect: 13,
		},
		{
			desc: "Single move",
			input: []byte{
				'R', ' ', '1', '\n',
			},
			expect: 1,
		},
		{
			desc: "Backtrack moves",
			input: []byte{
				'R', ' ', '2', '\n',
				'L', ' ', '2', '\n',
				'R', ' ', '1', '\n',
				'U', ' ', '1', '\n',
				'D', ' ', '2', '\n',
			},
			expect: 2,
		},
		{
			desc: "No diagonals",
			input: []byte{
				'R', ' ', '2', '\n',
				'L', ' ', '1', '\n',
				'U', ' ', '3', '\n',
			},
			expect: 4,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			fs := fs.NewMemFS()
			if err := fs.CreateFile("input", tc.input); err != nil {
				t.Fatal("Error creating in-memory test file:", err)
			}

			out := day09.Puzzle1(fs, "input")

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
			desc: "Provided example 1",
			input: []byte{
				'R', ' ', '4', '\n',
				'U', ' ', '4', '\n',
				'L', ' ', '3', '\n',
				'D', ' ', '1', '\n',
				'R', ' ', '4', '\n',
				'D', ' ', '1', '\n',
				'L', ' ', '5', '\n',
				'R', ' ', '2', '\n',
			},
			expect: 1,
		},
		{
			desc: "Provided example 2",
			input: []byte{
				'R', ' ', '5', '\n',
				'U', ' ', '8', '\n',
				'L', ' ', '8', '\n',
				'D', ' ', '3', '\n',
				'R', ' ', '1', '7', '\n',
				'D', ' ', '1', '0', '\n',
				'L', ' ', '2', '5', '\n',
				'U', ' ', '2', '0', '\n',
			},
			expect: 36,
		},
		{
			desc: "Large single move",
			input: []byte{
				'R', ' ', '1', '0', '\n',
			},
			expect: 2,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			fs := fs.NewMemFS()
			if err := fs.CreateFile("input", tc.input); err != nil {
				t.Fatal("Error creating in-memory test file:", err)
			}

			out := day09.Puzzle2(fs, "input")

			if out != tc.expect {
				t.Errorf("Expected %d, got %d", tc.expect, out)
			}
		})
	}
}
