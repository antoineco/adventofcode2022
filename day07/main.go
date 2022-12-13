package day07

import (
	"bufio"
	"bytes"
	"fmt"
	"strconv"
	"strings"

	"adventofcode2022/fs"
)

func Puzzle1(fs fs.FS, inputFile string) int {
	f, err := fs.Open(inputFile)
	if err != nil {
		panic(fmt.Errorf("opening input file: %w", err))
	}

	s := bufio.NewScanner(f)

	tree := newFsTree()

	currentDir := tree.rootNode

	for s.Scan() {
		if len(s.Bytes()) == 0 {
			continue
		}

		switch {
		case isCdCommand(s.Bytes()):
			switch dirName := string(s.Bytes()[5:]); dirName {
			case rootDir:
				currentDir = tree.rootNode

			case prevDir:
				if currentDir != tree.rootNode {
					currentDir = tree.upEdges[currentDir]
				}

			default:
				childDirNode, exists := tree.downEdges[currentDir][dirName]
				if !exists {
					childDirNode = &fsTreeNode{}
					tree.connect(dirName, childDirNode, currentDir)
				}

				currentDir = childDirNode
			}

		case isLsCommand(s.Bytes()):
			// no-op, following line scans read listed nodes until
			// the next command

		default:
			nodeSize, nodeName := readSizeAndName(s.Bytes())

			if _, exists := tree.downEdges[currentDir][nodeName]; exists {
				continue
			}

			childNode := &fsTreeNode{}
			if nodeSize > 0 {
				childNode.typ = fsTreeTypeFile
				childNode.size = nodeSize
			}

			tree.connect(nodeName, childNode, currentDir)
		}
	}

	if err := s.Err(); err != nil {
		panic(fmt.Errorf("while scanning input file: %w", err))
	}

	//fmt.Print(tree)

	var subDirsSizes []int
	rootNodeSize := sumSizeAtFsNode(tree.rootNode, tree.downEdges, &subDirsSizes)

	var sumMaxDirSize int
	const maxTotalDirSize = 100_000

	if rootNodeSize <= maxTotalDirSize {
		sumMaxDirSize += rootNodeSize
	}
	for _, dirSize := range subDirsSizes {
		if dirSize <= maxTotalDirSize {
			sumMaxDirSize += dirSize
		}
	}

	return sumMaxDirSize
}

func Puzzle2(fs fs.FS, inputFile string) int {
	f, err := fs.Open(inputFile)
	if err != nil {
		panic(fmt.Errorf("opening input file: %w", err))
	}

	s := bufio.NewScanner(f)

	tree := newFsTree()

	currentDir := tree.rootNode

	for s.Scan() {
		if len(s.Bytes()) == 0 {
			continue
		}

		switch {
		case isCdCommand(s.Bytes()):
			switch dirName := string(s.Bytes()[5:]); dirName {
			case rootDir:
				currentDir = tree.rootNode

			case prevDir:
				if currentDir != tree.rootNode {
					currentDir = tree.upEdges[currentDir]
				}

			default:
				childDirNode, exists := tree.downEdges[currentDir][dirName]
				if !exists {
					childDirNode = &fsTreeNode{}
					tree.connect(dirName, childDirNode, currentDir)
				}

				currentDir = childDirNode
			}

		case isLsCommand(s.Bytes()):
			// no-op, following line scans read listed nodes until
			// the next command

		default:
			nodeSize, nodeName := readSizeAndName(s.Bytes())

			if _, exists := tree.downEdges[currentDir][nodeName]; exists {
				continue
			}

			childNode := &fsTreeNode{}
			if nodeSize > 0 {
				childNode.typ = fsTreeTypeFile
				childNode.size = nodeSize
			}

			tree.connect(nodeName, childNode, currentDir)
		}
	}

	if err := s.Err(); err != nil {
		panic(fmt.Errorf("while scanning input file: %w", err))
	}

	var subDirsSizes []int
	rootNodeSize := sumSizeAtFsNode(tree.rootNode, tree.downEdges, &subDirsSizes)

	const (
		totalDiskSpace = 70_000_000
		minUnusedSpace = 30_000_000
	)

	freeDiskSpace := totalDiskSpace - rootNodeSize
	if freeDiskSpace > minUnusedSpace {
		return -1
	}

	minDiskSpaceToFree := minUnusedSpace - freeDiskSpace

	if rootNodeSize < minDiskSpaceToFree {
		return -1
	}

	diskSizeToFree := -1

	if rootNodeSize < diskSizeToFree {
		diskSizeToFree = rootNodeSize
	}
	for _, dirSize := range subDirsSizes {
		if minDiskSpaceToFree <= dirSize && (dirSize < diskSizeToFree || diskSizeToFree < 0) {
			diskSizeToFree = dirSize
		}
	}

	return diskSizeToFree
}

const (
	rootDir = "/"
	prevDir = ".."
)

func isCdCommand(input []byte) bool {
	return input[0] == '$' && input[2] == 'c'
}

func isLsCommand(input []byte) bool {
	return input[0] == '$' && input[2] == 'l'
}

func readSizeAndName(input []byte) (size int, name string) {
	fields := bytes.Fields(input)
	if l := len(fields); l != 2 {
		panic(fmt.Errorf("read line with %d field(s) instead of 2", l))
	}

	if fields[0][0] != 'd' {
		var val []byte
		for _, digit := range fields[0] {
			val = append(val, digit)
		}

		v, err := strconv.Atoi(string(val))
		if err != nil {
			panic(fmt.Errorf("parsing %v to integer: %w", v, err))
		}

		size = v
	}

	return size, string(fields[1])
}

func newFsTree() *fsTree {
	tree := &fsTree{
		downEdges: make(map[*fsTreeNode]indexedFsTreeNodes),
		upEdges:   make(map[*fsTreeNode]*fsTreeNode),
	}

	tree.rootNode = &fsTreeNode{}

	tree.downEdges[tree.rootNode] = make(indexedFsTreeNodes, 0)

	return tree
}

type fsTree struct {
	rootNode *fsTreeNode
	// map parent nodes (tail) to their children nodes (head)
	downEdges map[*fsTreeNode]indexedFsTreeNodes
	// map children nodes (head) to their parent node (tail)
	upEdges map[*fsTreeNode]*fsTreeNode
}

// Indexes of nodes, belonging to a given parent, indexed by name to avoid
// populating duplicates.
type indexedFsTreeNodes map[string]*fsTreeNode

type fsTreeNode struct {
	typ  fsTreeType
	size int
}

type fsTreeType uint8

const (
	fsTreeTypeDir fsTreeType = iota
	fsTreeTypeFile
)

func (t *fsTree) connect(key string, child, parent *fsTreeNode) {
	if parent.typ != fsTreeTypeDir {
		panic(fmt.Errorf("attempt to connect node %q to parent %p that is not a directory", key, parent))
	}

	if t.downEdges[parent] == nil {
		t.downEdges[parent] = make(indexedFsTreeNodes, 1)
	}
	if _, exists := t.downEdges[parent][key]; exists {
		panic(fmt.Errorf("attempt to override existing child node %q for parent %p", key, parent))
	}

	t.downEdges[parent][key] = child

	if _, exists := t.upEdges[child]; exists {
		panic(fmt.Errorf("attempt to override existing parent for child node %q (%p)", key, child))
	}

	t.upEdges[child] = parent
}

// Only prints a few levels for quick verification.
func (t *fsTree) String() string {
	var b strings.Builder

	b.WriteString("- " + rootDir + " (dir)\n")

	for childName, child := range t.downEdges[t.rootNode] {
		b.WriteString("  - " + childName)
		if child.typ == fsTreeTypeDir {
			b.WriteString(" (dir)")
		} else {
			b.WriteString(fmt.Sprintf(" (file, size=%d)", child.size))
		}
		b.WriteByte('\n')

		if child.typ == fsTreeTypeDir {
			for childName, child := range t.downEdges[child] {
				b.WriteString("    - " + childName)
				if child.typ == fsTreeTypeDir {
					b.WriteString(" (dir)")
				} else {
					b.WriteString(fmt.Sprintf(" (file, size=%d)", child.size))
				}
				b.WriteByte('\n')

				if child.typ == fsTreeTypeDir {
					for childName, child := range t.downEdges[child] {
						b.WriteString("      - " + childName)
						if child.typ == fsTreeTypeDir {
							b.WriteString(" (dir)")
							if len(t.downEdges[child]) > 0 {
								b.WriteString("\n        - ...")
							}
						} else {
							b.WriteString(fmt.Sprintf(" (file, size=%d)", child.size))
						}
						b.WriteByte('\n')
					}
				}
			}
		}
	}

	return b.String()
}

func sumSizeAtFsNode(node *fsTreeNode, downEdges map[*fsTreeNode]indexedFsTreeNodes, dirsSizes *[]int) int {
	if node.typ == fsTreeTypeFile {
		return node.size
	}

	var sumSize int

	for _, child := range downEdges[node] {
		sumSize += sumSizeAtFsNode(child, downEdges, dirsSizes)
	}

	*dirsSizes = append(*dirsSizes, sumSize)

	return sumSize
}
