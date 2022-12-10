package day02_test

import (
	"testing"

	"adventofcode2022/day02"
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
				'A', ' ', 'Y', '\n',
				'B', ' ', 'X', '\n',
				'C', ' ', 'Z', '\n',
			},
			expect: 15,
		},
		{
			desc: "Only losses with various moves",
			input: []byte{
				'A', ' ', 'Z', '\n',
				'B', ' ', 'X', '\n',
				'C', ' ', 'Y', '\n',
			},
			expect: (3 + 0) + (1 + 0) + (2 + 0),
		},
		{
			desc: "Only draws with various moves",
			input: []byte{
				'A', ' ', 'X', '\n',
				'B', ' ', 'Y', '\n',
				'C', ' ', 'Z', '\n',
			},
			expect: (1 + 3) + (2 + 3) + (3 + 3),
		},
		{
			desc: "Only wins with various moves",
			input: []byte{
				'A', ' ', 'Y', '\n',
				'B', ' ', 'Z', '\n',
				'C', ' ', 'X', '\n',
			},
			expect: (2 + 6) + (3 + 6) + (1 + 6),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			fs := fs.NewMemFS()
			if err := fs.CreateFile("input", tc.input); err != nil {
				t.Fatal("Error creating in-memory test file:", err)
			}

			out := day02.Puzzle1(fs, "input")

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
				'A', ' ', 'Y', '\n',
				'B', ' ', 'X', '\n',
				'C', ' ', 'Z', '\n',
			},
			expect: 12,
		},
		{
			desc: "Should lose all",
			input: []byte{
				'A', ' ', 'X', '\n',
				'B', ' ', 'X', '\n',
				'C', ' ', 'X', '\n',
			},
			expect: (3 + 0) + (1 + 0) + (2 + 0),
		},
		{
			desc: "Should draw all",
			input: []byte{
				'A', ' ', 'Y', '\n',
				'B', ' ', 'Y', '\n',
				'C', ' ', 'Y', '\n',
			},
			expect: (1 + 3) + (2 + 3) + (3 + 3),
		},
		{
			desc: "Should win all",
			input: []byte{
				'A', ' ', 'Z', '\n',
				'B', ' ', 'Z', '\n',
				'C', ' ', 'Z', '\n',
			},
			expect: (2 + 6) + (3 + 6) + (1 + 6),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			fs := fs.NewMemFS()
			if err := fs.CreateFile("input", tc.input); err != nil {
				t.Fatal("Error creating in-memory test file:", err)
			}

			out := day02.Puzzle2(fs, "input")

			if out != tc.expect {
				t.Errorf("Expected %d, got %d", tc.expect, out)
			}
		})
	}
}
