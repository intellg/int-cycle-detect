package degree

import (
	"fmt"
	"math"
)

var counter int

func Calculate(floor, cup int, innerCalculate func(int, int) int) (degree int) {
	// 1.0 If eggs are enough then the binary tree is a non-hollow tree
	log2Floor := math.Log2(float64(floor))
	if float64(cup) > log2Floor {
		degree = int(log2Floor) + 1
		return
	}

	counter = 0
	degree = innerCalculate(floor, cup)
	fmt.Println(counter)
	return
}
