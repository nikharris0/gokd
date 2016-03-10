package gokd

import (
	"testing"
)

func TestNew(t *testing.T) {
	kd := New(2)
	if kd.root == nil {
		t.Errorf("tree.root == nil")
		t.Errorf("tree.dimensions == %d, wanted 2", kd.dimensions)
	}
}

func TestEuclideanDistance(t *testing.T) {
	p := []float64{1, 2}
	q := []float64{1, 2}
	dist, err := euclideanDistance(p, q)
	if err != nil {
		t.Errorf("euclideanDistance(p,q) had dimension mismatch")
	}

	if dist != 0 {
		t.Errorf("euclideanDistance(p,q) = %f, wanted 0", dist)
	}
}

func TestGetParts(t *testing.T) {
	nodes := []Node{
		Node{
			Coordinates: []float64{1, 1},
		},
		Node{
			Coordinates: []float64{2, 2},
		},
		Node{
			Coordinates: []float64{3, 3},
		},
	}

	left, right, median := getParts(nodes)
	if len(left) != 1 {
		t.Errorf("len(left) = %d, wanted 1", len(left))
	}

	if len(right) != 1 {
		t.Errorf("len(right) = %d, wanted 1", len(left))
	}

	if left[0].Coordinates[0] != 1 {
		t.Errorf("left[0].dimensions[0] = %f, wanted 1", left[0].Coordinates[0])
	}

	if right[0].Coordinates[0] != 3 {
		t.Errorf("right[0].dimensions[0] = %f, wanted 3", right[0].Coordinates[0])
	}

	if median.Coordinates[0] != 2 {
		t.Errorf("median[0] = %f, wanted 2", median.Coordinates[0])
	}
}
