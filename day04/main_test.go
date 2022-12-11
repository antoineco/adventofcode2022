package day04_test

import (
	"testing"

	"adventofcode2022/day04"
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
				'2', '-', '4', ',', '6', '-', '8', '\n',
				'2', '-', '3', ',', '4', '-', '5', '\n',
				'5', '-', '7', ',', '7', '-', '9', '\n',
				'2', '-', '8', ',', '3', '-', '7', '\n',
				'6', '-', '6', ',', '4', '-', '6', '\n',
				'2', '-', '6', ',', '4', '-', '8', '\n',
			},
			expect: 2,
		},
		{
			desc: "Strictly equal range",
			input: []byte{
				'2', '-', '2', ',', '2', '-', '2', '\n',
				'2', '-', '3', ',', '4', '-', '5', '\n',
			},
			expect: 1,
		},
		{
			desc: "Swapped overlapping ranges",
			input: []byte{
				'2', '-', '4', ',', '2', '-', '3', '\n',
				'2', '-', '3', ',', '2', '-', '4', '\n',
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

			out := day04.Puzzle1(fs, "input")

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
				'2', '-', '4', ',', '6', '-', '8', '\n',
				'2', '-', '3', ',', '4', '-', '5', '\n',
				'5', '-', '7', ',', '7', '-', '9', '\n',
				'2', '-', '8', ',', '3', '-', '7', '\n',
				'6', '-', '6', ',', '4', '-', '6', '\n',
				'2', '-', '6', ',', '4', '-', '8', '\n',
			},
			expect: 4,
		},
		{
			desc: "Strictly equal range",
			input: []byte{
				'2', '-', '2', ',', '2', '-', '2', '\n',
				'2', '-', '3', ',', '4', '-', '5', '\n',
			},
			expect: 1,
		},
		{
			desc: "Swapped overlapping ranges",
			input: []byte{
				'2', '-', '4', ',', '2', '-', '3', '\n',
				'2', '-', '3', ',', '2', '-', '4', '\n',
			},
			expect: 2,
		},
		{
			desc: "Contiguous range",
			input: []byte{
				'2', '-', '4', ',', '5', '-', '9', '\n',
				'1', '-', '3', ',', '3', '-', '4', '\n',
			},
			expect: 1,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			fs := fs.NewMemFS()
			if err := fs.CreateFile("input", tc.input); err != nil {
				t.Fatal("Error creating in-memory test file:", err)
			}

			out := day04.Puzzle2(fs, "input")

			if out != tc.expect {
				t.Errorf("Expected %d, got %d", tc.expect, out)
			}
		})
	}
}
