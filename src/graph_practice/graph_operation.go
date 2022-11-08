package main

import (
	"fmt"
	"sync"
)

// Node a single node that composes one value.
type Node struct {
	value string
}

// String to print out the value of the Node
func (n *Node) String() string {
	return fmt.Sprintf("%v", n.value)
}

// ItemGraph the Items graph
type ItemGraph struct {
	nodes []*Node
	edges map[Node][]*Node
	lock  sync.RWMutex
}

// AddNode adds a node to the graph
func (g *ItemGraph) AddNode(n *Node) {
	g.lock.Lock()
	g.nodes = append(g.nodes, n)
	g.lock.Unlock()
}

// AddEdge adds an edge to the graph and store it to the edges map
func (g *ItemGraph) AddEdge(n1, n2 *Node) {
	g.lock.Lock()
	if g.edges == nil {
		g.edges = make(map[Node][]*Node)
	}

	// this is the bidrectional graaph, so the edge will be added to two map items
	g.edges[*n1] = append(g.edges[*n1], n2)
	g.edges[*n2] = append(g.edges[*n2], n1)
	g.lock.Unlock()
}

// String shows the whole graph structure as adjacency list
func (g *ItemGraph) String() {
	// RLock(): multiple go routine can read(not write) at a time by acquiring the lock.
	g.lock.RLock()
	s := ""
	for i := 0; i < len(g.nodes); i++ {
		s += g.nodes[i].String() + " -> "
		near := g.edges[*g.nodes[i]]
		for j := 0; j < len(near); j++ {
			s += near[j].String() + " "
		}
		s += "\n"
	}
	fmt.Println(s)
	g.lock.RUnlock()
}

// Traverse implements the BFS traversing algorithm
func (g *ItemGraph) Traverse() {
	// traverse is a reading operation. Use RLock() to allow multiple go routine
	// read(not write) at a time by acquiring the lock.
	fmt.Println("\ntraverse passes through the following nodes in order:")
	g.lock.RLock()
	q := NodeQueue{}
	q.New()
	n := g.nodes[0]
	q.Enqueue(*n)
	visited := make(map[*Node]bool)
	for {
		if q.IsEmpty() {
			break
		}
		node := q.Dequeue()
		visited[node] = true
		fmt.Println("visited node: ", node.value)
		near := g.edges[*node]

		for i := 0; i < len(near); i++ {
			j := near[i]
			if !visited[j] {
				q.Enqueue(*j)
				visited[j] = true

			}
		}

	}
	g.lock.RUnlock()
}
