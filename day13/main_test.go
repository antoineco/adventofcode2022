package day13_test

import (
	"testing"

	"adventofcode2022/day13"
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
				'[', '1', ',', '1', ',', '3', ',', '1', ',', '1', ']', '\n',
				'[', '1', ',', '1', ',', '5', ',', '1', ',', '1', ']', '\n',
				'\n',
				'[', '[', '1', ']', ',', '[', '2', ',', '3', ',', '4', ']', ']', '\n',
				'[', '[', '1', ']', ',', '4', ']', '\n',
				'\n',
				'[', '9', ']', '\n',
				'[', '[', '8', ',', '7', ',', '6', ']', ']', '\n',
				'\n',
				'[', '[', '4', ',', '4', ']', ',', '4', ',', '4', ']', '\n',
				'[', '[', '4', ',', '4', ']', ',', '4', ',', '4', ',', '4', ']', '\n',
				'\n',
				'[', '7', ',', '7', ',', '7', ',', '7', ']', '\n',
				'[', '7', ',', '7', ',', '7', ']', '\n',
				'\n',
				'[', ']', '\n',
				'[', '3', ']', '\n',
				'\n',
				'[', '[', '[', ']', ']', ']', '\n',
				'[', '[', ']', ']', '\n',
				'\n',
				'[', '1', ',', '[', '2', ',', '[', '3', ',', '[', '4', ',', '[', '5', ',', '6', ',', '7', ']', ']', ']', ']', ',', '8', ',', '9', ']', '\n',
				'[', '1', ',', '[', '2', ',', '[', '3', ',', '[', '4', ',', '[', '5', ',', '6', ',', '0', ']', ']', ']', ']', ',', '8', ',', '9', ']', '\n',
			},
			expect: 13,
		},
		/*
			My initial solution succeeded on the example, but
			yielded an incorrect result on the following case.

			My mistake was to compare "[]int vs int" using
			  compare(left[0], right)
			instead of
			  compare(left, []int{right})

			Therefore, it was performing the following comparisons:

			- Comparing [[[9,3,0,8],[5]]] vs [9,[[6,3,9],5]]
			   - Comparing [[9,3,0,8],[5]] vs 9
			     - Comparing [9,3,0,8] vs 9
			       - Comparing 9 vs 9
			   - Comparing 3 vs [[6,3,9],5]
			     - Comparing 3 vs [6,3,9]
			       - Comparing 3 vs 6
			         -> in order

			whereas the correct comparison is:

			- Comparing [[[9,3,0,8],[5]]] vs [9,[[6,3,9],5]]
			  - Comparing [[9,3,0,8],[5]] vs 9
			    - Comparing [[9,3,0,8],[5]] vs [9]
			      - Comparing [9,3,0,8] vs 9
			        - Comparing [9,3,0,8] vs [9]
			          - Comparing 9 vs 9
			            -> not in order (right is out of items)
		*/
		{
			desc: "Nested list vs integer",
			input: []byte{
				'[', '[', '[', '9', ',', '3', ',', '0', ',', '8', ']', ',', '[', '5', ']', ']', ']', '\n',
				'[', '9', ',', '[', '[', '6', ',', '3', ',', '9', ']', ',', '5', ']', ']', '\n',
			},
			expect: 0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			fs := fs.NewMemFS()
			if err := fs.CreateFile("input", tc.input); err != nil {
				t.Fatal("Error creating in-memory test file:", err)
			}

			out := day13.Puzzle1(fs, "input")

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
				'[', '1', ',', '1', ',', '3', ',', '1', ',', '1', ']', '\n',
				'[', '1', ',', '1', ',', '5', ',', '1', ',', '1', ']', '\n',
				'\n',
				'[', '[', '1', ']', ',', '[', '2', ',', '3', ',', '4', ']', ']', '\n',
				'[', '[', '1', ']', ',', '4', ']', '\n',
				'\n',
				'[', '9', ']', '\n',
				'[', '[', '8', ',', '7', ',', '6', ']', ']', '\n',
				'\n',
				'[', '[', '4', ',', '4', ']', ',', '4', ',', '4', ']', '\n',
				'[', '[', '4', ',', '4', ']', ',', '4', ',', '4', ',', '4', ']', '\n',
				'\n',
				'[', '7', ',', '7', ',', '7', ',', '7', ']', '\n',
				'[', '7', ',', '7', ',', '7', ']', '\n',
				'\n',
				'[', ']', '\n',
				'[', '3', ']', '\n',
				'\n',
				'[', '[', '[', ']', ']', ']', '\n',
				'[', '[', ']', ']', '\n',
				'\n',
				'[', '1', ',', '[', '2', ',', '[', '3', ',', '[', '4', ',', '[', '5', ',', '6', ',', '7', ']', ']', ']', ']', ',', '8', ',', '9', ']', '\n',
				'[', '1', ',', '[', '2', ',', '[', '3', ',', '[', '4', ',', '[', '5', ',', '6', ',', '0', ']', ']', ']', ']', ',', '8', ',', '9', ']', '\n',
			},
			expect: 140,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			fs := fs.NewMemFS()
			if err := fs.CreateFile("input", tc.input); err != nil {
				t.Fatal("Error creating in-memory test file:", err)
			}

			out := day13.Puzzle2(fs, "input")

			if out != tc.expect {
				t.Errorf("Expected %d, got %d", tc.expect, out)
			}
		})
	}
}
