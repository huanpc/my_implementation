package main

import (
	"fmt"
	"math"
	"sort"
	"strings"
)

func networkDelayTime(times [][]int, n, k int) int {
	// Create an adjacency map
	adjacency := make(map[int][][]int)
	// Store source as key, and destination and time as values
	for _, time := range times {
		source := time[0]
		destination := time[1]
		travelTime := time[2]
		if _, ok := adjacency[source]; !ok {
			adjacency[source] = [][]int{}
		}
		adjacency[source] = append(adjacency[source], []int{destination, travelTime})
	}

	fmt.Println("\t Adjacency dictionary:", strings.Replace(fmt.Sprint(adjacency), " ", ", ", -1))

	pq := make([][]int, 0)        // Initialize our queue with (time, node)
	pq = append(pq, []int{0, k})  // Add the source node with a delay time of 0
	visited := make(map[int]bool) // To store the visited nodes
	delays := 0                   // To store the delay time

	for len(pq) > 0 {
		// Sort the priority queue based on time
		sort.Slice(pq, func(i, j int) bool {
			return pq[i][0] < pq[j][0]
		})
		current := pq[0]
		pq = pq[1:]

		// Get the minimum time node from the queue
		time := current[0]
		node := current[1]

		fmt.Println("\t Retrieved time:", time)
		fmt.Println("\t Retrieved node:", node)

		//  If node is already visited we will continue to next iteration
		if visited[node] {
			continue
		}

		visited[node] = true                                   // Mark the node as visited
		delays = int(math.Max(float64(delays), float64(time))) // Update the delay time if necessary
		neighbors := adjacency[node]

		fmt.Println("\t Neighbors:", strings.Replace(fmt.Sprint(neighbors), " ", ", ", -1))

		// Add all the unvisited neighbors of the current node to the queue with their new delay time
		for _, neighbor := range neighbors {
			neighborNode := neighbor[0]
			neighborTime := neighbor[1]
			if !visited[neighborNode] {
				newTime := time + neighborTime
				pq = append(pq, []int{newTime, neighborNode})
			}
		}
	}

	// If all nodes have been visited, return the delay time else return -1
	if len(visited) == n {
		return delays
	}

	return -1
}

func main() {
	times := [][][]int{
		{{2, 1, 1}, {3, 2, 1}, {3, 4, 2}},
		{{2, 1, 1}, {1, 3, 1}, {3, 4, 2}, {5, 4, 2}},
		{{1, 2, 1}, {2, 3, 1}, {3, 4, 1}},
		{{1, 2, 1}, {2, 3, 1}, {3, 5, 2}},
		{{1, 2, 2}},
	}

	n := []int{4, 5, 4, 5, 2}
	k := []int{3, 1, 1, 1, 2}

	for i := 0; i < len(times); i++ {
		fmt.Printf("%d.\t times = %s\n", i+1, strings.Replace(fmt.Sprint(times[i]), " ", ", ", -1))
		fmt.Println("\t number of nodes 'n' =", n[i])
		fmt.Println("\t starting node 'k' =", k[i], "\n")
		fmt.Println("\t Minimum amount of time required =", networkDelayTime(times[i], n[i], k[i]))
		fmt.Println(strings.Repeat("-", 100))
	}
}
