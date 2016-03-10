package gokd

type bydimension []Node

func (b bydimension) Len() int {
	return len(b)
}

func (b bydimension) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func (b bydimension) Less(i, j int) bool {
	dim := b[i].depth % b[i].tree.dimensions
	return b[i].Coordinates[dim] < b[j].Coordinates[dim]
}
