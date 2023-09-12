package main

import "container/list"

/**
 * Definition for a Node.
 */
type Node struct {
	Val       int
	Neighbors []*Node
}

var (
	cloned map[int]*Node
)

func cloneGraph(node *Node) *Node {
	// using DFS to clone, start with root node, if neighbor is cloned, recursively clone it by adding to stack, after that, add it to root's neighbor
	// a new map to store cloned node
	// cloned = make(map[int]*Node)
	// return dfs(node)
	if node == nil {
		return nil
	}
	return bfs(node)
}

// recursion
// time: O(N)
// space: O(N)
func dfs(node *Node) *Node {
	if node == nil {
		return nil
	}
	clone := &Node{Val: node.Val, Neighbors: make([]*Node, 0)}
	cloned[node.Val] = clone
	// iterate neighbors
	for _, nb := range node.Neighbors {
		clonedNB, exist := cloned[nb.Val]
		if !exist {
			clonedNB = dfs(nb)
		}
		clone.Neighbors = append(clone.Neighbors, clonedNB)
	}
	return clone
}

// BFS, interactive, clone level by level, clone node first, neighbor later
// time: O(N)
// space: O(N)
func bfs(node *Node) *Node {
	cloned := make(map[int]*Node)
	queue := list.New()
	// start with root
	root := cloneNode(node)
	cloned[node.Val] = root
	queue.PushBack(node)
	for queue.Len() > 0 {
		e := queue.Front()
		queue.Remove(e)
		cur := e.Value.(*Node)
		neighbors := make([]*Node, 0)
		for _, nb := range cur.Neighbors {
			// clone if not cloned
			if _, exist := cloned[nb.Val]; !exist {
				queue.PushBack(nb)
				cloned[nb.Val] = cloneNode(nb)
			}
			neighbors = append(neighbors, cloned[nb.Val])
		}
		cloned[cur.Val].Neighbors = neighbors
	}
	return root
}

func cloneNode(node *Node) *Node {
	return &Node{Val: node.Val}
}
