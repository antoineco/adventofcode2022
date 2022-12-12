package day05_test

import (
	"testing"

	"adventofcode2022/day05"
	"adventofcode2022/fs"
)

func TestPuzzle1(t *testing.T) {
	testCases := []struct {
		desc   string
		input  []byte
		expect string
	}{
		{
			desc: "Provided example",
			input: []byte{
				' ', ' ', ' ', ' ', '[', 'D', ']', ' ', ' ', ' ', ' ', '\n',
				'[', 'N', ']', ' ', '[', 'C', ']', ' ', ' ', ' ', ' ', '\n',
				'[', 'Z', ']', ' ', '[', 'M', ']', ' ', '[', 'P', ']', '\n',
				' ', '1', ' ', ' ', ' ', '2', ' ', ' ', ' ', '3', '\n',
				'\n',
				'm', 'o', 'v', 'e', ' ', '1', ' ', 'f', 'r', 'o', 'm', ' ', '2', ' ', 't', 'o', ' ', '1', '\n',
				'm', 'o', 'v', 'e', ' ', '3', ' ', 'f', 'r', 'o', 'm', ' ', '1', ' ', 't', 'o', ' ', '3', '\n',
				'm', 'o', 'v', 'e', ' ', '2', ' ', 'f', 'r', 'o', 'm', ' ', '2', ' ', 't', 'o', ' ', '1', '\n',
				'm', 'o', 'v', 'e', ' ', '1', ' ', 'f', 'r', 'o', 'm', ' ', '1', ' ', 't', 'o', ' ', '2', '\n',
			},
			expect: "CMZ",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			fs := fs.NewMemFS()
			if err := fs.CreateFile("input", tc.input); err != nil {
				t.Fatal("Error creating in-memory test file:", err)
			}

			out := day05.Puzzle1(fs, "input")

			if out != tc.expect {
				t.Errorf("Expected %q, got %q", tc.expect, out)
			}
		})
	}
}

func TestPuzzle2(t *testing.T) {
	testCases := []struct {
		desc   string
		input  []byte
		expect string
	}{
		{
			desc: "Provided example",
			input: []byte{
				' ', ' ', ' ', ' ', '[', 'D', ']', ' ', ' ', ' ', ' ', '\n',
				'[', 'N', ']', ' ', '[', 'C', ']', ' ', ' ', ' ', ' ', '\n',
				'[', 'Z', ']', ' ', '[', 'M', ']', ' ', '[', 'P', ']', '\n',
				' ', '1', ' ', ' ', ' ', '2', ' ', ' ', ' ', '3', '\n',
				'\n',
				'm', 'o', 'v', 'e', ' ', '1', ' ', 'f', 'r', 'o', 'm', ' ', '2', ' ', 't', 'o', ' ', '1', '\n',
				'm', 'o', 'v', 'e', ' ', '3', ' ', 'f', 'r', 'o', 'm', ' ', '1', ' ', 't', 'o', ' ', '3', '\n',
				'm', 'o', 'v', 'e', ' ', '2', ' ', 'f', 'r', 'o', 'm', ' ', '2', ' ', 't', 'o', ' ', '1', '\n',
				'm', 'o', 'v', 'e', ' ', '1', ' ', 'f', 'r', 'o', 'm', ' ', '1', ' ', 't', 'o', ' ', '2', '\n',
			},
			expect: "MCD",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			fs := fs.NewMemFS()
			if err := fs.CreateFile("input", tc.input); err != nil {
				t.Fatal("Error creating in-memory test file:", err)
			}

			out := day05.Puzzle2(fs, "input")

			if out != tc.expect {
				t.Errorf("Expected %q, got %q", tc.expect, out)
			}
		})
	}
}
