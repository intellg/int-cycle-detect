package graph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBreadthFirstSearch(t *testing.T) {
	graph := [][]int{
		{0, 1, 0, 1, 0, 0},
		{1, 0, 0, 1, 1, 0},
		{0, 0, 0, 0, 1, 1},
		{1, 1, 0, 0, 1, 0},
		{0, 1, 1, 1, 0, 0},
		{0, 0, 1, 0, 0, 0},
	}

	expectQ := []int{0, 1, 3, 4, 2, 5}
	queue := BreadthFirstSearch(graph, 0)
	assert.Equal(t, queue, expectQ, "Wrong")
}
