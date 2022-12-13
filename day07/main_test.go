package day07_test

import (
	"testing"

	"adventofcode2022/day07"
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
				'$', ' ', 'c', 'd', ' ', '/', '\n',
				'$', ' ', 'l', 's', '\n',
				'd', 'i', 'r', ' ', 'a', '\n',
				'1', '4', '8', '4', '8', '5', '1', '4', ' ', 'b', '.', 't', 'x', 't', '\n',
				'8', '5', '0', '4', '1', '5', '6', ' ', 'c', '.', 'd', 'a', 't', '\n',
				'd', 'i', 'r', ' ', 'd', '\n',
				'$', ' ', 'c', 'd', ' ', 'a', '\n',
				'$', ' ', 'l', 's', '\n',
				'd', 'i', 'r', ' ', 'e', '\n',
				'2', '9', '1', '1', '6', ' ', 'f', '\n',
				'2', '5', '5', '7', ' ', 'g', '\n',
				'6', '2', '5', '9', '6', ' ', 'h', '.', 'l', 's', 't', '\n',
				'$', ' ', 'c', 'd', ' ', 'e', '\n',
				'$', ' ', 'l', 's', '\n',
				'5', '8', '4', ' ', 'i', '\n',
				'$', ' ', 'c', 'd', ' ', '.', '.', '\n',
				'$', ' ', 'c', 'd', ' ', '.', '.', '\n',
				'$', ' ', 'c', 'd', ' ', 'd', '\n',
				'$', ' ', 'l', 's', '\n',
				'4', '0', '6', '0', '1', '7', '4', ' ', 'j', '\n',
				'8', '0', '3', '3', '0', '2', '0', ' ', 'd', '.', 'l', 'o', 'g', '\n',
				'5', '6', '2', '6', '1', '5', '2', ' ', 'd', '.', 'e', 'x', 't', '\n',
				'7', '2', '1', '4', '2', '9', '6', ' ', 'k', '\n',
			},
			expect: 95437,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			fs := fs.NewMemFS()
			if err := fs.CreateFile("input", tc.input); err != nil {
				t.Fatal("Error creating in-memory test file:", err)
			}

			out := day07.Puzzle1(fs, "input")

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
				'$', ' ', 'c', 'd', ' ', '/', '\n',
				'$', ' ', 'l', 's', '\n',
				'd', 'i', 'r', ' ', 'a', '\n',
				'1', '4', '8', '4', '8', '5', '1', '4', ' ', 'b', '.', 't', 'x', 't', '\n',
				'8', '5', '0', '4', '1', '5', '6', ' ', 'c', '.', 'd', 'a', 't', '\n',
				'd', 'i', 'r', ' ', 'd', '\n',
				'$', ' ', 'c', 'd', ' ', 'a', '\n',
				'$', ' ', 'l', 's', '\n',
				'd', 'i', 'r', ' ', 'e', '\n',
				'2', '9', '1', '1', '6', ' ', 'f', '\n',
				'2', '5', '5', '7', ' ', 'g', '\n',
				'6', '2', '5', '9', '6', ' ', 'h', '.', 'l', 's', 't', '\n',
				'$', ' ', 'c', 'd', ' ', 'e', '\n',
				'$', ' ', 'l', 's', '\n',
				'5', '8', '4', ' ', 'i', '\n',
				'$', ' ', 'c', 'd', ' ', '.', '.', '\n',
				'$', ' ', 'c', 'd', ' ', '.', '.', '\n',
				'$', ' ', 'c', 'd', ' ', 'd', '\n',
				'$', ' ', 'l', 's', '\n',
				'4', '0', '6', '0', '1', '7', '4', ' ', 'j', '\n',
				'8', '0', '3', '3', '0', '2', '0', ' ', 'd', '.', 'l', 'o', 'g', '\n',
				'5', '6', '2', '6', '1', '5', '2', ' ', 'd', '.', 'e', 'x', 't', '\n',
				'7', '2', '1', '4', '2', '9', '6', ' ', 'k', '\n',
			},
			expect: 24933642,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			fs := fs.NewMemFS()
			if err := fs.CreateFile("input", tc.input); err != nil {
				t.Fatal("Error creating in-memory test file:", err)
			}

			out := day07.Puzzle2(fs, "input")

			if out != tc.expect {
				t.Errorf("Expected %d, got %d", tc.expect, out)
			}
		})
	}
}
