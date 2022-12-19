package day12_test

import (
	"testing"

	"adventofcode2022/day12"
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
				'S', 'a', 'b', 'q', 'p', 'o', 'n', 'm', '\n',
				'a', 'b', 'c', 'r', 'y', 'x', 'x', 'l', '\n',
				'a', 'c', 'c', 's', 'z', 'E', 'x', 'k', '\n',
				'a', 'c', 'c', 't', 'u', 'v', 'w', 'j', '\n',
				'a', 'b', 'd', 'e', 'f', 'g', 'h', 'i', '\n',
			},
			expect: 31,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			fs := fs.NewMemFS()
			if err := fs.CreateFile("input", tc.input); err != nil {
				t.Fatal("Error creating in-memory test file:", err)
			}

			out := day12.Puzzle1(fs, "input")

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
				'S', 'a', 'b', 'q', 'p', 'o', 'n', 'm', '\n',
				'a', 'b', 'c', 'r', 'y', 'x', 'x', 'l', '\n',
				'a', 'c', 'c', 's', 'z', 'E', 'x', 'k', '\n',
				'a', 'c', 'c', 't', 'u', 'v', 'w', 'j', '\n',
				'a', 'b', 'd', 'e', 'f', 'g', 'h', 'i', '\n',
			},
			expect: 29,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			fs := fs.NewMemFS()
			if err := fs.CreateFile("input", tc.input); err != nil {
				t.Fatal("Error creating in-memory test file:", err)
			}

			out := day12.Puzzle2(fs, "input")

			if out != tc.expect {
				t.Errorf("Expected %d, got %d", tc.expect, out)
			}
		})
	}
}
