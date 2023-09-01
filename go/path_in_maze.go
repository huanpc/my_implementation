package main

// https://leetcode.com/problems/paths-in-maze-that-lead-to-same-room/submissions/

import (
	"sort"
	"strconv"
	"strings"
)

func numberOfPaths(n int, corridors [][]int) int {
	return sol2(n, corridors)
}

// approach 1, build adjacency matrix first, then iterate corridor and get common neighbor, but need to use unique path map to dedup
// time: O(E*V), space: O(V^2+E)
func sol1(n int, corridors [][]int) int {
	// create adjacency matrix
	adj := make([][]int, n+1)
	for _, corridor := range corridors {
		r1, r2 := corridor[0], corridor[1]
		if len(adj[r1]) == 0 {
			adj[r1] = make([]int, n+1)
		}
		if len(adj[r2]) == 0 {
			adj[r2] = make([]int, n+1)
		}
		adj[r1][r2] = 1
		adj[r2][r1] = 1
	}
	// iterate thru each corridor, find intersection neighbor rooms for 2 room in corridor
	// each intersection is account for a unique path
	cnt := 0
	existPath := make(map[string]bool)
	for _, corridor := range corridors {
		r1, r2 := corridor[0], corridor[1]
		// each intersection room form a path
		for i := 1; i < n+1; i++ {
			if adj[r1][i] == adj[r2][i] && adj[r1][i] == 1 {
				// check if unique path
				path := []string{strconv.Itoa(r1), strconv.Itoa(r2), strconv.Itoa(i)}
				sort.Strings(path)
				pathStr := strings.Join(path, "-")
				if existPath[pathStr] {
					continue
				}
				existPath[pathStr] = true
				cnt++
			}
		}
	}
	return cnt
}

// approach 2, interative way, build adjacency along with get common, no need to use unique path map
// time: O(E), space: O(V^2)
func sol2(n int, corridors [][]int) int {
	cnt := 0
	adj := make(map[int]map[int]bool, n)
	for _, corridor := range corridors {
		r1, r2 := corridor[0], corridor[1]
		if _, ok := adj[r1]; !ok {
			adj[r1] = make(map[int]bool)
		}
		if _, ok := adj[r2]; !ok {
			adj[r2] = make(map[int]bool)
		}
		adj[r1][r2] = true
		adj[r2][r1] = true
		cnt += countCommonNeighbors(adj[r1], adj[r2])
	}
	return cnt
}

func countCommonNeighbors(adj1, adj2 map[int]bool) int {
	cnt := 0
	for r := range adj1 {
		if _, ok := adj2[r]; ok {
			cnt += 1
		}
	}
	return cnt
}
