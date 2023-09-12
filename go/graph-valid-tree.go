package main

// https://leetcode.com/problems/graph-valid-tree/

// time: O(n)
// space: O(n*m)
func validTree(n int, edges [][]int) bool {
	// base case, if n - 1 < or > number of edges, that means there is cycle or unconnected nodes
	if n-1 != len(edges) {
		return false
	}
	visited := make(map[int]bool)
	adj := make(map[int][]int)
	// build adjacency list
	for _, e := range edges {
		a, b := e[0], e[1]
		adj[a] = append(adj[a], b)
		adj[b] = append(adj[b], a)
	}
	stack := make([]int, 0)
	// use DFS with stack, start with first node, iterate and visit all neighbor, if visited nodes is less than n, return false
	stack = append(stack, 0)
	visited[0] = true
	for len(stack) > 0 {
		// pop
		curNode := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		// add neighbors to stack
		// if already visited, that mean cycle, return immediately
		for _, node := range adj[curNode] {
			if _, yes := visited[node]; !yes {
				visited[node] = true
				stack = append(stack, node)
			}
		}
	}
	return len(visited) == n
}
