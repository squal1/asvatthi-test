package main

import (
	"container/heap"
	"fmt"
)

type Node struct {
	x, y       int     // coordinates
	g, h, f    float64 // costs
	parent     *Node
	open, done bool
}

type PriorityQueue []*Node

func (pq PriorityQueue) Len() int           { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool { return pq[i].f < pq[j].f }
func (pq PriorityQueue) Swap(i, j int)      { pq[i], pq[j] = pq[j], pq[i] }

func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*Node)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func heuristic(current, goal *Node) float64 {
	return float64(abs(current.x-goal.x) + abs(current.y-goal.y))
}

func AStar(start, goal *Node) []*Node {
	openSet := make(PriorityQueue, 0)
	heap.Init(&openSet)

	start.g = 0
	start.h = heuristic(start, goal)
	start.f = start.g + start.h

	heap.Push(&openSet, start)

	for openSet.Len() > 0 {
		current := heap.Pop(&openSet).(*Node)

		if current == goal {
			path := make([]*Node, 0)
			for current != nil {
				path = append([]*Node{current}, path...)
				current = current.parent
			}
			return path
		}

		current.done = true

		neighbors := []*Node{
			{current.x - 1, current.y, 0, 0, 0, current, false, false},
			{current.x + 1, current.y, 0, 0, 0, current, false, false},
			{current.x, current.y - 1, 0, 0, 0, current, false, false},
			{current.x, current.y + 1, 0, 0, 0, current, false, false},
		}

		for _, neighbor := range neighbors {
			if neighbor.x < 0 || neighbor.x >= 10 || neighbor.y < 0 || neighbor.y >= 10 {
				continue
			}

			if neighbor.done {
				continue
			}

			if neighbor.open {
				continue
			}

			neighbor.g = current.g + 1
			neighbor.h = heuristic(neighbor, goal)
			neighbor.f = neighbor.g + neighbor.h
			neighbor.parent = current

			heap.Push(&openSet, neighbor)
			neighbor.open = true
		}
	}

	return nil
}

func main() {
	start := &Node{0, 0, 0, 0, 0, nil, false, false}
	goal := &Node{9, 9, 0, 0, 0, nil, false, false}

	path := AStar(start, goal)

	if path != nil {
		for _, node := range path {
			fmt.Printf("(%d, %d) -> ", node.x, node.y)
		}
		fmt.Println("Goal")
	} else {
		fmt.Println("No path found")
	}
}
