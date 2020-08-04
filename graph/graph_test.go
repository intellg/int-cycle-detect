package graph

import "testing"

func TestBreadthFirstSearch(t *testing.T) {
	graph := [][]int{
		{0, 1, 0, 1, 0, 0},
		{1, 0, 0, 1, 1, 0},
		{0, 0, 0, 0, 1, 1},
		{1, 1, 0, 0, 1, 0},
		{0, 1, 1, 1, 0, 0},
		{0, 0, 1, 0, 0, 0},
	}
	queue := BreadthFirstSearch(graph, 0)
	t.Log(queue)
}
