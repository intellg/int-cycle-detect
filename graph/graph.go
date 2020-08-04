package graph

import "math"

const (
	WHITE = iota
	GRAY
	BLACK
)

type Node struct {
	Color  int
	Depth  int
	Parent *Node
}

func BreadthFirstSearch(graph [][]int, rootId int) (queue []int) {
	nodes := make([]*Node, len(graph))
	for i := 0; i < len(nodes); i++ {
		if i == rootId {
			nodes[i] = &Node{
				Color: GRAY,
				Depth: 0,
			}
		} else {
			nodes[i] = &Node{
				Color: WHITE,
				Depth: math.MaxInt64,
			}
		}
	}

	queue = append(queue, rootId)
	for i := 0; i < len(queue); {
		cur := queue[i]
		i++
		for j := 0; j < len(graph[cur]); j++ {
			if graph[cur][j] == 1 && nodes[j].Color == WHITE {
				nodes[j].Color = GRAY
				nodes[j].Depth = nodes[cur].Depth + 1
				nodes[j].Parent = nodes[cur]
				queue = append(queue, j)
			}
		}
		nodes[cur].Color = BLACK
	}
	return
}
