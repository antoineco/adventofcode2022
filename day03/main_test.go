package day03_test

import (
	"testing"

	"adventofcode2022/day03"
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
				'v', 'J', 'r', 'w', 'p', 'W', 't', 'w', 'J', 'g', 'W', 'r', 'h', 'c', 's', 'F', 'M',
				'M', 'f', 'F', 'F', 'h', 'F', 'p', '\n',
				'j', 'q', 'H', 'R', 'N', 'q', 'R', 'j', 'q', 'z', 'j', 'G', 'D', 'L', 'G', 'L', 'r',
				's', 'F', 'M', 'f', 'F', 'Z', 'S', 'r', 'L', 'r', 'F', 'Z', 's', 'S', 'L', '\n',
				'P', 'm', 'm', 'd', 'z', 'q', 'P', 'r', 'V', 'v', 'P', 'w', 'w', 'T', 'W', 'B', 'w',
				'g', '\n',
				'w', 'M', 'q', 'v', 'L', 'M', 'Z', 'H', 'h', 'H', 'M', 'v', 'w', 'L', 'H', 'j', 'b',
				'v', 'c', 'j', 'n', 'n', 'S', 'B', 'n', 'v', 'T', 'Q', 'F', 'n', '\n',
				't', 't', 'g', 'J', 't', 'R', 'G', 'J', 'Q', 'c', 't', 'T', 'Z', 't', 'Z', 'T', '\n',
				'C', 'r', 'Z', 's', 'J', 's', 'P', 'P', 'Z', 's', 'G', 'z', 'w', 'w', 's', 'L', 'w',
				'L', 'm', 'p', 'w', 'M', 'D', 'w', '\n',
			},
			expect: 157,
		},
		{
			desc: "Multiple times the same item type v",
			input: []byte{
				'v', 'J', 'r', 'M', 'v', 'F', 'F', 'h', '\n',
				'v', 'J', 'r', 'M', 'v', 'F', 'F', 'h', '\n',
				'v', 'J', 'r', 'M', 'v', 'F', 'F', 'h', '\n',
			},
			expect: 22 + 22 + 22,
		},
		{
			desc: "Odd number of items",
			input: []byte{
				'v', 'J', 'r', 'M', 'v', 'x', 'F', 'F', 'h', '\n',
			},
			expect: 22,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			fs := fs.NewMemFS()
			if err := fs.CreateFile("input", tc.input); err != nil {
				t.Fatal("Error creating in-memory test file:", err)
			}

			out := day03.Puzzle1(fs, "input")

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
				'v', 'J', 'r', 'w', 'p', 'W', 't', 'w', 'J', 'g', 'W', 'r', 'h', 'c', 's', 'F', 'M',
				'M', 'f', 'F', 'F', 'h', 'F', 'p', '\n',
				'j', 'q', 'H', 'R', 'N', 'q', 'R', 'j', 'q', 'z', 'j', 'G', 'D', 'L', 'G', 'L', 'r',
				's', 'F', 'M', 'f', 'F', 'Z', 'S', 'r', 'L', 'r', 'F', 'Z', 's', 'S', 'L', '\n',
				'P', 'm', 'm', 'd', 'z', 'q', 'P', 'r', 'V', 'v', 'P', 'w', 'w', 'T', 'W', 'B', 'w',
				'g', '\n',
				'w', 'M', 'q', 'v', 'L', 'M', 'Z', 'H', 'h', 'H', 'M', 'v', 'w', 'L', 'H', 'j', 'b',
				'v', 'c', 'j', 'n', 'n', 'S', 'B', 'n', 'v', 'T', 'Q', 'F', 'n', '\n',
				't', 't', 'g', 'J', 't', 'R', 'G', 'J', 'Q', 'c', 't', 'T', 'Z', 't', 'Z', 'T', '\n',
				'C', 'r', 'Z', 's', 'J', 's', 'P', 'P', 'Z', 's', 'G', 'z', 'w', 'w', 's', 'L', 'w',
				'L', 'm', 'p', 'w', 'M', 'D', 'w', '\n',
			},
			expect: 70,
		},
		{
			desc: "Reduced sample",
			input: []byte{
				'a', 'b', 'x', '\n',
				'x', 'c', 'd', '\n',
				'e', 'f', 'x', '\n',

				'a', 'b', 'y', '\n',
				'y', 'c', 'd', '\n',
				'e', 'f', 'y', '\n',

				'a', 'b', 'z', '\n',
				'z', 'c', 'd', '\n',
				'e', 'f', 'z', '\n',
			},
			expect: 24 + 25 + 26,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			fs := fs.NewMemFS()
			if err := fs.CreateFile("input", tc.input); err != nil {
				t.Fatal("Error creating in-memory test file:", err)
			}

			out := day03.Puzzle2(fs, "input")

			if out != tc.expect {
				t.Errorf("Expected %d, got %d", tc.expect, out)
			}
		})
	}
}
