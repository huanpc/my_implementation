package main

import (
	"container/heap"
)

func networkDelayTime(times [][]int, n, k int) int {
	nodeDelay := make(map[int]int)
	// build adjacency list
	adjList := make(map[int][]Node, n)
	for _, e := range times {
		src, dst, delay := e[0], e[1], e[2]
		adjList[src] = append(adjList[src], Node{id: dst, delay: delay})
	}

	// use min heap to choose next smallest delay node to visit
	minHeap := &MinHeap{&Node{id: k, delay: 0}}
	nodeDelay[k] = 0
	for minHeap.Len() > 0 {
		node := heap.Pop(minHeap).(*Node)
		if nodeDelay[node.id] > node.delay {
			nodeDelay[node.id] = node.delay
		}
		for _, nb := range adjList[node.id] {
			newDelay := nb.delay + node.delay
			// ignore path which lead to higher delay
			if val, ok := nodeDelay[nb.id]; ok && val < newDelay {
				continue
			}
			nodeDelay[nb.id] = newDelay
			heap.Push(minHeap, &Node{id: nb.id, delay: newDelay})
		}
	}
	if len(nodeDelay) < n {
		return -1
	}
	maxDelay := 0
	for _, delay := range nodeDelay {
		if maxDelay < delay {
			maxDelay = delay
		}
	}
	return maxDelay
}

type Node struct {
	id    int
	delay int
}

type MinHeap []*Node

func (p MinHeap) Len() int {
	return len(p)
}

func (p MinHeap) Less(i int, j int) bool {
	// we want min heap
	return p[i].delay < p[j].delay
}

// Swap swaps the elements with indexes i and j.
func (p MinHeap) Swap(i int, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p *MinHeap) Push(x any) {
	*p = append(*p, x.(*Node))
}

func (p *MinHeap) Pop() any {
	old := *p
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	*p = old[0 : n-1]
	return item
}
