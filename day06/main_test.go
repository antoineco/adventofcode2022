package day06_test

import (
	"testing"

	"adventofcode2022/day06"
	"adventofcode2022/fs"
)

func TestPuzzle1(t *testing.T) {
	testCases := []struct {
		desc   string
		input  []byte
		expect int
	}{
		{
			desc: "Provided example 1",
			input: []byte{
				'm', 'j', 'q', 'j', 'p', 'q', 'm', 'g', 'b', 'l', 'j', 's', 'p', 'h', 'd', 'z', 't',
				'n', 'v', 'j', 'f', 'q', 'w', 'r', 'c', 'g', 's', 'm', 'l', 'b', '\n',
			},
			expect: 7,
		},
		{
			desc: "Provided example 2",
			input: []byte{
				'b', 'v', 'w', 'b', 'j', 'p', 'l', 'b', 'g', 'v', 'b', 'h', 's', 'r', 'l', 'p', 'g',
				'd', 'm', 'j', 'q', 'w', 'f', 't', 'v', 'n', 'c', 'z', '\n',
			},
			expect: 5,
		},
		{
			desc: "Provided example 3",
			input: []byte{
				'n', 'p', 'p', 'd', 'v', 'j', 't', 'h', 'q', 'l', 'd', 'p', 'w', 'n', 'c', 'q', 's',
				'z', 'v', 'f', 't', 'b', 'r', 'm', 'j', 'l', 'h', 'g', '\n',
			},
			expect: 6,
		},
		{
			desc: "Provided example 4",
			input: []byte{
				'n', 'z', 'n', 'r', 'n', 'f', 'r', 'f', 'n', 't', 'j', 'f', 'm', 'v', 'f', 'w', 'm',
				'z', 'd', 'f', 'j', 'l', 'v', 't', 'q', 'n', 'b', 'h', 'c', 'p', 'r', 's', 'g', '\n',
			},
			expect: 10,
		},
		{
			desc: "Provided example 5",
			input: []byte{
				'z', 'c', 'f', 'z', 'f', 'w', 'z', 'z', 'q', 'f', 'r', 'l', 'j', 'w', 'z', 'l', 'r',
				'f', 'n', 'p', 'q', 'd', 'b', 'h', 't', 'm', 's', 'c', 'g', 'v', 'j', 'w', '\n',
			},
			expect: 11,
		},
		{
			desc: "Begins with packet marker",
			input: []byte{
				'w', 'x', 'y', 'z', 'a', 'a', 'a', 'a', '\n',
			},
			expect: 4,
		},
		{
			desc: "One lead character before packet marker",
			input: []byte{
				'y', 'w', 'x', 'y', 'z', 'a', 'a', 'a', 'a', '\n',
			},
			expect: 5,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			fs := fs.NewMemFS()
			if err := fs.CreateFile("input", tc.input); err != nil {
				t.Fatal("Error creating in-memory test file:", err)
			}

			out := day06.Puzzle1(fs, "input")

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
				'm', 'j', 'q', 'j', 'p', 'q', 'm', 'g', 'b', 'l', 'j', 's', 'p', 'h', 'd', 'z', 't',
				'n', 'v', 'j', 'f', 'q', 'w', 'r', 'c', 'g', 's', 'm', 'l', 'b', '\n',
			},
			expect: 19,
		},
		{
			desc: "Provided example 2",
			input: []byte{
				'b', 'v', 'w', 'b', 'j', 'p', 'l', 'b', 'g', 'v', 'b', 'h', 's', 'r', 'l', 'p', 'g',
				'd', 'm', 'j', 'q', 'w', 'f', 't', 'v', 'n', 'c', 'z', '\n',
			},
			expect: 23,
		},
		{
			desc: "Provided example 3",
			input: []byte{
				'n', 'p', 'p', 'd', 'v', 'j', 't', 'h', 'q', 'l', 'd', 'p', 'w', 'n', 'c', 'q', 's',
				'z', 'v', 'f', 't', 'b', 'r', 'm', 'j', 'l', 'h', 'g', '\n',
			},
			expect: 23,
		},
		{
			desc: "Provided example 4",
			input: []byte{
				'n', 'z', 'n', 'r', 'n', 'f', 'r', 'f', 'n', 't', 'j', 'f', 'm', 'v', 'f', 'w', 'm',
				'z', 'd', 'f', 'j', 'l', 'v', 't', 'q', 'n', 'b', 'h', 'c', 'p', 'r', 's', 'g', '\n',
			},
			expect: 29,
		},
		{
			desc: "Provided example 5",
			input: []byte{
				'z', 'c', 'f', 'z', 'f', 'w', 'z', 'z', 'q', 'f', 'r', 'l', 'j', 'w', 'z', 'l', 'r',
				'f', 'n', 'p', 'q', 'd', 'b', 'h', 't', 'm', 's', 'c', 'g', 'v', 'j', 'w', '\n',
			},
			expect: 26,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			fs := fs.NewMemFS()
			if err := fs.CreateFile("input", tc.input); err != nil {
				t.Fatal("Error creating in-memory test file:", err)
			}

			out := day06.Puzzle2(fs, "input")

			if out != tc.expect {
				t.Errorf("Expected %d, got %d", tc.expect, out)
			}
		})
	}
}
