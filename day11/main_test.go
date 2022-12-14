package day11_test

import (
	"testing"

	"adventofcode2022/day11"
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
				'M', 'o', 'n', 'k', 'e', 'y', ' ', '0', ':', '\n',
				' ', ' ', 'S', 't', 'a', 'r', 't', 'i', 'n', 'g', ' ', 'i', 't', 'e', 'm', 's', ':', ' ', '7', '9', ',', ' ', '9', '8', '\n',
				' ', ' ', 'O', 'p', 'e', 'r', 'a', 't', 'i', 'o', 'n', ':', ' ', 'n', 'e', 'w', ' ', '=', ' ', 'o', 'l', 'd', ' ', '*', ' ', '1', '9', '\n',
				' ', ' ', 'T', 'e', 's', 't', ':', ' ', 'd', 'i', 'v', 'i', 's', 'i', 'b', 'l', 'e', ' ', 'b', 'y', ' ', '2', '3', '\n',
				' ', ' ', ' ', ' ', 'I', 'f', ' ', 't', 'r', 'u', 'e', ':', ' ', 't', 'h', 'r', 'o', 'w', ' ', 't', 'o', ' ', 'm', 'o', 'n', 'k', 'e', 'y', ' ', '2', '\n',
				' ', ' ', ' ', ' ', 'I', 'f', ' ', 'f', 'a', 'l', 's', 'e', ':', ' ', 't', 'h', 'r', 'o', 'w', ' ', 't', 'o', ' ', 'm', 'o', 'n', 'k', 'e', 'y', ' ', '3', '\n',
				'\n',
				'M', 'o', 'n', 'k', 'e', 'y', ' ', '1', ':', '\n',
				' ', ' ', 'S', 't', 'a', 'r', 't', 'i', 'n', 'g', ' ', 'i', 't', 'e', 'm', 's', ':', ' ', '5', '4', ',', ' ', '6', '5', ',', ' ', '7', '5', ',', ' ', '7', '4', '\n',
				' ', ' ', 'O', 'p', 'e', 'r', 'a', 't', 'i', 'o', 'n', ':', ' ', 'n', 'e', 'w', ' ', '=', ' ', 'o', 'l', 'd', ' ', '+', ' ', '6', '\n',
				' ', ' ', 'T', 'e', 's', 't', ':', ' ', 'd', 'i', 'v', 'i', 's', 'i', 'b', 'l', 'e', ' ', 'b', 'y', ' ', '1', '9', '\n',
				' ', ' ', ' ', ' ', 'I', 'f', ' ', 't', 'r', 'u', 'e', ':', ' ', 't', 'h', 'r', 'o', 'w', ' ', 't', 'o', ' ', 'm', 'o', 'n', 'k', 'e', 'y', ' ', '2', '\n',
				' ', ' ', ' ', ' ', 'I', 'f', ' ', 'f', 'a', 'l', 's', 'e', ':', ' ', 't', 'h', 'r', 'o', 'w', ' ', 't', 'o', ' ', 'm', 'o', 'n', 'k', 'e', 'y', ' ', '0', '\n',
				'\n',
				'M', 'o', 'n', 'k', 'e', 'y', ' ', '2', ':', '\n',
				' ', ' ', 'S', 't', 'a', 'r', 't', 'i', 'n', 'g', ' ', 'i', 't', 'e', 'm', 's', ':', ' ', '7', '9', ',', ' ', '6', '0', ',', ' ', '9', '7', '\n',
				' ', ' ', 'O', 'p', 'e', 'r', 'a', 't', 'i', 'o', 'n', ':', ' ', 'n', 'e', 'w', ' ', '=', ' ', 'o', 'l', 'd', ' ', '*', ' ', 'o', 'l', 'd', '\n',
				' ', ' ', 'T', 'e', 's', 't', ':', ' ', 'd', 'i', 'v', 'i', 's', 'i', 'b', 'l', 'e', ' ', 'b', 'y', ' ', '1', '3', '\n',
				' ', ' ', ' ', ' ', 'I', 'f', ' ', 't', 'r', 'u', 'e', ':', ' ', 't', 'h', 'r', 'o', 'w', ' ', 't', 'o', ' ', 'm', 'o', 'n', 'k', 'e', 'y', ' ', '1', '\n',
				' ', ' ', ' ', ' ', 'I', 'f', ' ', 'f', 'a', 'l', 's', 'e', ':', ' ', 't', 'h', 'r', 'o', 'w', ' ', 't', 'o', ' ', 'm', 'o', 'n', 'k', 'e', 'y', ' ', '3', '\n',
				'\n',
				'M', 'o', 'n', 'k', 'e', 'y', ' ', '3', ':', '\n',
				' ', ' ', 'S', 't', 'a', 'r', 't', 'i', 'n', 'g', ' ', 'i', 't', 'e', 'm', 's', ':', ' ', '7', '4', '\n',
				' ', ' ', 'O', 'p', 'e', 'r', 'a', 't', 'i', 'o', 'n', ':', ' ', 'n', 'e', 'w', ' ', '=', ' ', 'o', 'l', 'd', ' ', '+', ' ', '3', '\n',
				' ', ' ', 'T', 'e', 's', 't', ':', ' ', 'd', 'i', 'v', 'i', 's', 'i', 'b', 'l', 'e', ' ', 'b', 'y', ' ', '1', '7', '\n',
				' ', ' ', ' ', ' ', 'I', 'f', ' ', 't', 'r', 'u', 'e', ':', ' ', 't', 'h', 'r', 'o', 'w', ' ', 't', 'o', ' ', 'm', 'o', 'n', 'k', 'e', 'y', ' ', '0', '\n',
				' ', ' ', ' ', ' ', 'I', 'f', ' ', 'f', 'a', 'l', 's', 'e', ':', ' ', 't', 'h', 'r', 'o', 'w', ' ', 't', 'o', ' ', 'm', 'o', 'n', 'k', 'e', 'y', ' ', '1', '\n',
			},
			expect: 10605,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			fs := fs.NewMemFS()
			if err := fs.CreateFile("input", tc.input); err != nil {
				t.Fatal("Error creating in-memory test file:", err)
			}

			out := day11.Puzzle1(fs, "input")

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
				'M', 'o', 'n', 'k', 'e', 'y', ' ', '0', ':', '\n',
				' ', ' ', 'S', 't', 'a', 'r', 't', 'i', 'n', 'g', ' ', 'i', 't', 'e', 'm', 's', ':', ' ', '7', '9', ',', ' ', '9', '8', '\n',
				' ', ' ', 'O', 'p', 'e', 'r', 'a', 't', 'i', 'o', 'n', ':', ' ', 'n', 'e', 'w', ' ', '=', ' ', 'o', 'l', 'd', ' ', '*', ' ', '1', '9', '\n',
				' ', ' ', 'T', 'e', 's', 't', ':', ' ', 'd', 'i', 'v', 'i', 's', 'i', 'b', 'l', 'e', ' ', 'b', 'y', ' ', '2', '3', '\n',
				' ', ' ', ' ', ' ', 'I', 'f', ' ', 't', 'r', 'u', 'e', ':', ' ', 't', 'h', 'r', 'o', 'w', ' ', 't', 'o', ' ', 'm', 'o', 'n', 'k', 'e', 'y', ' ', '2', '\n',
				' ', ' ', ' ', ' ', 'I', 'f', ' ', 'f', 'a', 'l', 's', 'e', ':', ' ', 't', 'h', 'r', 'o', 'w', ' ', 't', 'o', ' ', 'm', 'o', 'n', 'k', 'e', 'y', ' ', '3', '\n',
				'\n',
				'M', 'o', 'n', 'k', 'e', 'y', ' ', '1', ':', '\n',
				' ', ' ', 'S', 't', 'a', 'r', 't', 'i', 'n', 'g', ' ', 'i', 't', 'e', 'm', 's', ':', ' ', '5', '4', ',', ' ', '6', '5', ',', ' ', '7', '5', ',', ' ', '7', '4', '\n',
				' ', ' ', 'O', 'p', 'e', 'r', 'a', 't', 'i', 'o', 'n', ':', ' ', 'n', 'e', 'w', ' ', '=', ' ', 'o', 'l', 'd', ' ', '+', ' ', '6', '\n',
				' ', ' ', 'T', 'e', 's', 't', ':', ' ', 'd', 'i', 'v', 'i', 's', 'i', 'b', 'l', 'e', ' ', 'b', 'y', ' ', '1', '9', '\n',
				' ', ' ', ' ', ' ', 'I', 'f', ' ', 't', 'r', 'u', 'e', ':', ' ', 't', 'h', 'r', 'o', 'w', ' ', 't', 'o', ' ', 'm', 'o', 'n', 'k', 'e', 'y', ' ', '2', '\n',
				' ', ' ', ' ', ' ', 'I', 'f', ' ', 'f', 'a', 'l', 's', 'e', ':', ' ', 't', 'h', 'r', 'o', 'w', ' ', 't', 'o', ' ', 'm', 'o', 'n', 'k', 'e', 'y', ' ', '0', '\n',
				'\n',
				'M', 'o', 'n', 'k', 'e', 'y', ' ', '2', ':', '\n',
				' ', ' ', 'S', 't', 'a', 'r', 't', 'i', 'n', 'g', ' ', 'i', 't', 'e', 'm', 's', ':', ' ', '7', '9', ',', ' ', '6', '0', ',', ' ', '9', '7', '\n',
				' ', ' ', 'O', 'p', 'e', 'r', 'a', 't', 'i', 'o', 'n', ':', ' ', 'n', 'e', 'w', ' ', '=', ' ', 'o', 'l', 'd', ' ', '*', ' ', 'o', 'l', 'd', '\n',
				' ', ' ', 'T', 'e', 's', 't', ':', ' ', 'd', 'i', 'v', 'i', 's', 'i', 'b', 'l', 'e', ' ', 'b', 'y', ' ', '1', '3', '\n',
				' ', ' ', ' ', ' ', 'I', 'f', ' ', 't', 'r', 'u', 'e', ':', ' ', 't', 'h', 'r', 'o', 'w', ' ', 't', 'o', ' ', 'm', 'o', 'n', 'k', 'e', 'y', ' ', '1', '\n',
				' ', ' ', ' ', ' ', 'I', 'f', ' ', 'f', 'a', 'l', 's', 'e', ':', ' ', 't', 'h', 'r', 'o', 'w', ' ', 't', 'o', ' ', 'm', 'o', 'n', 'k', 'e', 'y', ' ', '3', '\n',
				'\n',
				'M', 'o', 'n', 'k', 'e', 'y', ' ', '3', ':', '\n',
				' ', ' ', 'S', 't', 'a', 'r', 't', 'i', 'n', 'g', ' ', 'i', 't', 'e', 'm', 's', ':', ' ', '7', '4', '\n',
				' ', ' ', 'O', 'p', 'e', 'r', 'a', 't', 'i', 'o', 'n', ':', ' ', 'n', 'e', 'w', ' ', '=', ' ', 'o', 'l', 'd', ' ', '+', ' ', '3', '\n',
				' ', ' ', 'T', 'e', 's', 't', ':', ' ', 'd', 'i', 'v', 'i', 's', 'i', 'b', 'l', 'e', ' ', 'b', 'y', ' ', '1', '7', '\n',
				' ', ' ', ' ', ' ', 'I', 'f', ' ', 't', 'r', 'u', 'e', ':', ' ', 't', 'h', 'r', 'o', 'w', ' ', 't', 'o', ' ', 'm', 'o', 'n', 'k', 'e', 'y', ' ', '0', '\n',
				' ', ' ', ' ', ' ', 'I', 'f', ' ', 'f', 'a', 'l', 's', 'e', ':', ' ', 't', 'h', 'r', 'o', 'w', ' ', 't', 'o', ' ', 'm', 'o', 'n', 'k', 'e', 'y', ' ', '1', '\n',
			},
			expect: 2_713_310_158,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			fs := fs.NewMemFS()
			if err := fs.CreateFile("input", tc.input); err != nil {
				t.Fatal("Error creating in-memory test file:", err)
			}

			out := day11.Puzzle2(fs, "input")

			if out != tc.expect {
				t.Errorf("Expected %d, got %d", tc.expect, out)
			}
		})
	}
}
