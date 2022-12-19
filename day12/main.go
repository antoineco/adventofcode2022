package day12

import (
	"bufio"
	"fmt"
	"math"

	"adventofcode2022/fs"
)

func Puzzle1(fs fs.FS, inputFile string) int {
	f, err := fs.Open(inputFile)
	if err != nil {
		panic(fmt.Errorf("opening input file: %w", err))
	}

	s := bufio.NewScanner(f)

	starts := make(set) // we expect a single start in part 1
	var end position

	var squares []byte
	columns := -1

	for s.Scan() {
		if columns < 0 {
			columns = len(s.Bytes())
		}

		for i, square := range s.Bytes() {
			switch square {
			case 'S':
				starts.add(position{x: i, y: len(squares) / columns})
				square = 'a'
			case 'E':
				end = position{x: i, y: len(squares) / columns}
				square = 'z'
			}

			squares = append(squares, square)
		}
	}

	if err := s.Err(); err != nil {
		panic(fmt.Errorf("while scanning input file: %w", err))
	}

	// NOTE: because each square is represented as a slice index in the
	// heightmap, a graph of integer indices would have been enough.
	// I used a graph of (x,y) tuples instead to facilitate troubleshooting.
	paths := newGraph()

	// The graph is constructed from 'E' to 'S' (backwards) so that the
	// same logic can be applied to both parts 1 and 2.
	// The Breadth-First Traversal is performed in that order, and
	// backtracks from 'S' to 'E' once 'S' is reached.
	for i := 0; i < len(squares); i++ {
		// left
		if i%columns > 0 {
			if squares[i] <= squares[i-1]+1 {
				paths.connect(position{x: i % columns, y: i / columns}, position{x: (i - 1) % columns, y: i / columns})
			}
		}
		// right
		if i%columns < columns-1 {
			if squares[i] <= squares[i+1]+1 {
				paths.connect(position{x: i % columns, y: i / columns}, position{x: (i + 1) % columns, y: i / columns})
			}
		}
		// up
		if i/columns > 0 {
			if squares[i] <= squares[i-columns]+1 {
				paths.connect(position{x: i % columns, y: i / columns}, position{x: i % columns, y: (i / columns) - 1})
			}
		}
		// down
		if i/columns < (len(squares)-1)/columns {
			if squares[i] <= squares[i+columns]+1 {
				paths.connect(position{x: i % columns, y: i / columns}, position{x: i % columns, y: (i / columns) + 1})
			}
		}
	}

	return shortestPath(paths, end, starts)
}

func Puzzle2(fs fs.FS, inputFile string) int {
	f, err := fs.Open(inputFile)
	if err != nil {
		panic(fmt.Errorf("opening input file: %w", err))
	}

	s := bufio.NewScanner(f)

	starts := make(set)
	var end position

	var squares []byte
	columns := -1

	for s.Scan() {
		if columns < 0 {
			columns = len(s.Bytes())
		}

		for i, square := range s.Bytes() {
			switch square {
			case 'S':
				square = 'a'
			case 'E':
				end = position{x: i, y: len(squares) / columns}
				square = 'z'
			}

			if square == 'a' {
				starts.add(position{x: i, y: len(squares) / columns})
			}

			squares = append(squares, square)
		}
	}

	if err := s.Err(); err != nil {
		panic(fmt.Errorf("while scanning input file: %w", err))
	}

	// NOTE: because each square is represented as a slice index in the
	// heightmap, a graph of integer indices would have been enough.
	// I used a graph of (x,y) tuples instead to facilitate troubleshooting.
	paths := newGraph()

	// The graph is constructed from 'E' to 'S' (backwards).
	// The Breadth-First Traversal is performed in that order, and
	// backtracks from 'a' to 'E' once the first 'a' is reached.
	for i := 0; i < len(squares); i++ {
		// left
		if i%columns > 0 {
			if squares[i] <= squares[i-1]+1 {
				paths.connect(position{x: i % columns, y: i / columns}, position{x: (i - 1) % columns, y: i / columns})
			}
		}
		// right
		if i%columns < columns-1 {
			if squares[i] <= squares[i+1]+1 {
				paths.connect(position{x: i % columns, y: i / columns}, position{x: (i + 1) % columns, y: i / columns})
			}
		}
		// up
		if i/columns > 0 {
			if squares[i] <= squares[i-columns]+1 {
				paths.connect(position{x: i % columns, y: i / columns}, position{x: i % columns, y: (i / columns) - 1})
			}
		}
		// down
		if i/columns < (len(squares)-1)/columns {
			if squares[i] <= squares[i+columns]+1 {
				paths.connect(position{x: i % columns, y: i / columns}, position{x: i % columns, y: (i / columns) + 1})
			}
		}
	}

	return shortestPath(paths, end, starts)
}

var nowhere = position{x: math.MaxInt, y: math.MaxInt}

// Breadth-First Traversal
func shortestPath(g *graph, end position, starts set) int {
	next := newQueue()
	visited := newGraph() // connects visited nodes to their previously visited parent

	next.add(end)
	visited.connect(end, nowhere)

	start := nowhere

traversal:
	for next.len() > 0 {
		square := next.get()

		for _, adjacent := range g.downEdges[square] {
			if !visited.hasTail(adjacent) {
				visited.connect(adjacent, square)
				next.add(adjacent)
			}

			if starts.has(adjacent) {
				start = adjacent
				break traversal
			}
		}
	}

	if !visited.hasTail(start) {
		return -1
	}

	steps := 0

	for nextStep := start; nextStep != end; {
		nextStep = visited.downEdges[nextStep][0]
		steps++
	}

	return steps
}

type position struct {
	x, y int
}

type graph struct {
	downEdges map[position][]position
}

func newGraph() *graph {
	return &graph{
		downEdges: make(map[position][]position),
	}
}

func (g *graph) connect(tail, head position) {
	g.downEdges[tail] = append(g.downEdges[tail], head)
}

func (g *graph) hasTail(tail position) bool {
	_, has := g.downEdges[tail]
	return has
}

type queue struct {
	index set
	items []position
}

func newQueue() *queue {
	return &queue{
		index: make(set),
	}
}

func (q *queue) has(v position) bool {
	return q.index.has(v)
}

func (q *queue) add(v position) {
	if q.index.has(v) {
		return
	}

	q.items = append(q.items, v)
	q.index.add(v)
}

func (q *queue) get() (v position) {
	v, q.items = q.items[0], q.items[1:]
	q.index.del(v)
	return v
}

func (q *queue) len() int {
	return q.index.len()
}

type set map[position]struct{}

func (s *set) has(v position) bool {
	_, exists := (*s)[v]
	return exists
}

func (s *set) add(v position) {
	if !s.has(v) {
		(*s)[v] = struct{}{}
	}
}

func (s *set) del(v position) {
	if s.has(v) {
		delete(*s, v)
	}
}

func (s *set) len() int {
	return len(*s)
}
