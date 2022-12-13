package day08_test

import (
	"testing"

	"adventofcode2022/day08"
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
				'3', '0', '3', '7', '3', '\n',
				'2', '5', '5', '1', '2', '\n',
				'6', '5', '3', '3', '2', '\n',
				'3', '3', '5', '4', '9', '\n',
				'3', '5', '3', '9', '0', '\n',
			},
			expect: 21,
		},
		{
			desc: "Single row",
			input: []byte{
				'3', '0', '3', '7', '3', '\n',
			},
			expect: 5,
		},
		{
			desc: "Single column",
			input: []byte{
				'3', '\n',
				'2', '\n',
				'6', '\n',
				'3', '\n',
				'3', '\n',
			},
			expect: 5,
		},
		{
			desc: "No inner tree 1",
			input: []byte{
				'3', '3', '\n',
				'2', '2', '\n',
				'6', '2', '\n',
				'3', '9', '\n',
				'3', '0', '\n',
			},
			expect: 10,
		},
		{
			desc: "No inner tree 2",
			input: []byte{
				'3', '0', '3', '7', '3', '\n',
				'3', '5', '3', '9', '0', '\n',
			},
			expect: 10,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			fs := fs.NewMemFS()
			if err := fs.CreateFile("input", tc.input); err != nil {
				t.Fatal("Error creating in-memory test file:", err)
			}

			out := day08.Puzzle1(fs, "input")

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
				'3', '0', '3', '7', '3', '\n',
				'2', '5', '5', '1', '2', '\n',
				'6', '5', '3', '3', '2', '\n',
				'3', '3', '5', '4', '9', '\n',
				'3', '5', '3', '9', '0', '\n',
			},
			expect: 8,
		},
		{
			desc: "Extended sample",
			input: []byte{
				'0', '0', '2', '0', '0', '\n',
				'0', '0', '3', '0', '0', '\n',
				'0', '0', '1', '0', '0', '\n',
				'1', '2', '9', '2', '1', '\n',
				'0', '0', '1', '0', '0', '\n',
				'0', '0', '3', '0', '0', '\n',
				'0', '0', '2', '0', '0', '\n',
			},
			expect: 36,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			fs := fs.NewMemFS()
			if err := fs.CreateFile("input", tc.input); err != nil {
				t.Fatal("Error creating in-memory test file:", err)
			}

			out := day08.Puzzle2(fs, "input")

			if out != tc.expect {
				t.Errorf("Expected %d, got %d", tc.expect, out)
			}
		})
	}
}
