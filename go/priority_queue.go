package main

import (
	"container/heap"
	"fmt"
)

type Item struct {
	value    string
	priority int
}

type PriorityQueue []*Item

// Len is the number of elements in the collection.
func (p PriorityQueue) Len() int {
	return len(p)
}

// Less reports whether the element with index i
// must sort before the element with index j.
//
// If both Less(i, j) and Less(j, i) are false,
// then the elements at index i and j are considered equal.
// Sort may place equal elements in any order in the final result,
// while Stable preserves the original input order of equal elements.
//
// Less must describe a transitive ordering:
//   - if both Less(i, j) and Less(j, k) are true, then Less(i, k) must be true as well.
//   - if both Less(i, j) and Less(j, k) are false, then Less(i, k) must be false as well.
//
// Note that floating-point comparison (the < operator on float32 or float64 values)
// is not a transitive ordering when not-a-number (NaN) values are involved.
// See Float64Slice.Less for a correct implementation for floating-point values.
func (p PriorityQueue) Less(i int, j int) bool {
	// we want max heap
	return p[i].priority < p[j].priority
}

// Swap swaps the elements with indexes i and j.
func (p PriorityQueue) Swap(i int, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p *PriorityQueue) Push(x any) {
	item := x.(*Item)
	*p = append(*p, item)
}

func (p *PriorityQueue) Pop() any {
	old := *p
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	*p = old[0 : n-1]
	return item
}

func main() {
	// Some items and their priorities.
	items := map[string]int{
		"banana": 3, "apple": 2, "pear": 4, "joy": 1,
	}

	// Create a priority queue, put the items in it, and
	// establish the priority queue (heap) invariants.
	pq := make(PriorityQueue, len(items))
	i := 0
	for value, priority := range items {
		pq[i] = &Item{
			value:    value,
			priority: priority,
		}
		i++
	}
	heap.Init(&pq)
	fmt.Printf("%v\n", pq[0])

	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*Item)
		fmt.Printf("%.2d:%s ", item.priority, item.value)
	}
}
