package main

import (
	"fmt"

	"int-fun/floor-cup/degree"
	"int-fun/floor-cup/node"
)

func main() {
	floor := 1000
	cup := 8

	dgr := degree.Calculate(floor, cup, degree.InnerCalculateA)
	root := node.Calculate(floor, cup, dgr)

	fmt.Printf("Floor=%d, cup=%d => at most try %d times\n", floor, cup, dgr)
	node.OutputJson(root)
}
