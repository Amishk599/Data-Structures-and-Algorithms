package main

import (
	"fmt"
)

func main() {
	for _, tc := range tcs {
		g := InitGraph(tc.numNodes)
		g.AddFirstNode(tc.edges[0][1])
		for _, edge := range tc.edges {
			g.AddNode(edge[0], edge[1])
		}
		// BFS Traversal
		fmt.Println("BFS : ")
		g.Bfs(1)
	}
}
