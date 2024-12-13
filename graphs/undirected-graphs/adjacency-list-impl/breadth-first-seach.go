package main

import (
	"container/list"
	"fmt"
)

/*
Uses a queue (FIFO) and visited list.
Dequeue a node from queue and visit each of
its neighbours and mark it as visted and also
push to the queue. Print each dequeued node
*/
func (g *Graph) Bfs(startNode int) {
	visited := make([]bool, g.numVertices+1)
	queue := list.New()

	queue.PushBack(startNode)
	visited[startNode] = true

	for queue.Len() > 0 {
		// dequeue a node from the front
		node := queue.Front().Value.(int)
		queue.Remove(queue.Front())
		fmt.Printf("%d ", node)
		// retieve all neighbours of this node
		neighbours := g.adjList[node]

		for _, neighbour := range neighbours {
			if !visited[neighbour] {
				visited[neighbour] = true
				queue.PushBack(neighbour)
			}
		}
	}
}
