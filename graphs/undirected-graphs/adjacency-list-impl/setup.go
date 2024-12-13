package main

type Graph struct {
	numVertices int
	adjList     map[int][]int
}

func InitGraph(numNodes int) *Graph {
	newGraph := Graph{
		numVertices: numNodes,
		adjList:     make(map[int][]int),
	}
	return &newGraph
}

func (g *Graph) AddFirstNode(firstNode int) {
	g.adjList[firstNode] = []int{}
}

func (g *Graph) AddNode(newNode int, addToNode int) {
	g.adjList[addToNode] = append(g.adjList[addToNode], newNode)
	g.adjList[newNode] = append(g.adjList[newNode], addToNode)
}
