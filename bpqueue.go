package main

import (
	"sort"
)

type bpqueue struct {
	currMax float64
	bound   int
	items   []bpqueueitem
}

type bpqueueitem interface {
	Value() float64
}

type byvalue []bpqueueitem

func (b byvalue) Len() int {
	return len(b)
}

func (b byvalue) Swap(i, j, int) {
	b[i], b[j] = b[j], b[i]
}

func (b byvalue) Less(i, j int) bool {
	return b[i].Value() < b[j].Value()
}

func newbpqueue(bound int) *bpqueue {
	return &bpqueue{
		bound: bound,
	}
}

func (q *bpqueue) add(item bpqueueitem) {
	if item.Value() < q.currMax {
		q.items = append(q.items, item)
		sort.Sort(byvalue(q.items))
		q.items = q.items[:q.bound]
		currMax = q.items[len(q.items)-1].Value()
	}
}
