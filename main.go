package gokd

import (
	"errors"
	"math"
	"sort"
)

// KDTree represents a static miltidimensional binary search tree
type KDTree struct {
	dimensions int
	root       *Node
	nodes      int64
}

// Node represents a single leaf or edge node within a KDTree
type Node struct {
	Coordinates []float64
	Left        *Node
	Right       *Node
	Data        interface{}
	depth       int
	distance    float64
	tree        *KDTree
}

// ErrDimensionMismatch occurs when two or more Node or float arrays
// are expected to have the same number of dimensions but do not
var ErrDimensionMismatch = errors.New("dimensionality mismatch")

// New returns a new KDTree with d dimensions
func New(d int) *KDTree {
	t := &KDTree{
		dimensions: d,
	}

	return t
}

// NewNode returns a new Node
func (t *KDTree) NewNode() *Node {
	return &Node{
		tree: t,
	}
}

// Load does a balanced insertion of nodes into the KDTree. The order of the nodes
// within nodes may be changed during sorting
func (t *KDTree) Load(nodes []Node) (err error) {
	for nodeIndex := range nodes {
		if len(nodes[nodeIndex].Coordinates) != t.dimensions {
			return ErrDimensionMismatch
		}

		nodes[nodeIndex].tree = t
	}

	t.root = recursiveBuild(t, 0, nodes)
	return nil
}

// Nearest does a nearest neighbor search on KDTree t, returning a maximum of
// n Nodes
func (t *KDTree) Nearest(coords []float64, n int) (nodes []*Node, err error) {
	if len(coords) != t.dimensions {
		return nil, ErrDimensionMismatch
	}

	return nil, nil
}

// Clear removes all nodes in KDTree t
func (t *KDTree) Clear() {
	t.root = nil
}

func recursiveBuild(tree *KDTree, depth int, nodes []Node) *Node {
	for nodeIndex := range nodes {
		nodes[nodeIndex].depth = depth
	}

	sort.Sort(bydimension(nodes))

	left, right, median := getParts(nodes)

	newNode := tree.NewNode()
	newNode.Coordinates = median.Coordinates
	newNode.Data = median.Data

	if len(left) != 0 {
		newNode.Left = recursiveBuild(tree, depth+1, left)
	}

	if len(right) != 0 {
		newNode.Right = recursiveBuild(tree, depth+1, right)
	}

	return newNode
}

func euclideanDistance(p []float64, q []float64) (dist float64, err error) {
	dims := len(p)
	if len(p) != len(q) {
		return 0, ErrDimensionMismatch
	}

	var sum float64
	for x := 0; x < dims; x++ {
		sum += ((p[x] - q[x]) * (p[x] - q[x]))
	}

	return math.Sqrt(sum), nil
}

func getParts(nodes []Node) (left []Node, right []Node, median *Node) {
	if len(nodes) == 0 {
		return nil, nil, nil
	}

	if len(nodes) == 1 {
		return nil, nil, &nodes[0]
	}

	if len(nodes) == 2 {
		return nodes[:1], nil, &nodes[1]
	}

	medianIndex := 0
	if len(nodes)%2 == 0 {
		medianIndex = -1
	}

	medianIndex += len(nodes) / 2
	left = nodes[:medianIndex]
	median = &nodes[medianIndex]
	right = nodes[medianIndex+1:]

	return left, right, median
}
