package main

type TestCase struct {
	numNodes  int
	edges     [][]int // Each edge is represented as [child, parent]
	startNode int     // Start node for BFS and DFS
}

var tcs = []TestCase{
	{
		numNodes:  8,
		edges:     [][]int{{2, 1}, {3, 1}, {4, 2}, {5, 2}, {6, 3}, {7, 6}, {8, 7}},
		startNode: 1,
	},
}
