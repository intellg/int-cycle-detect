// 给定一个有向图，图中可能含有多个环，请检测
package main

import (
	"fmt"
	"strings"
)

var (
	items = [][]string{
		//{"o", "a"},
		//{"o", "b"},
		//{"a", "c"},
		//{"b", "d"},
		//{"c", "e"},
		//{"e", "d"},
		//{"d", "c"},
		//{"c", "f"},
		//{"f", "c"},
		//{"d", "f"},
		//{"f", "d"},
		{"a", "b"},
		{"b", "c"},
		{"c", "d"},
		{"d", "a"},
		{"c", "b"},
		{"b", "d"},
		{"a", "c"},
	}
	stack    = make([]string, 0)
	pathArr  = make([]string, 0)
	cycleArr = make([]string, 0)
)

func main() {
	detect(items[0][0])
	stack = stack[0 : len(stack)-1]

	fmt.Println(cycleArr)
}

func detect(current string) bool {
	// Check whether the current path is ever visited
	path := strings.Join(stack, "") + current
	for _, eachPath := range pathArr {
		if eachPath == path {
			return false
		}
	}
	pathArr = append(pathArr, path)

	// Check whether the cycle is detected
	for i, item := range stack {
		if item == current {
			handleCycle(i, current)
			return false
		}
	}
	stack = append(stack, current)

	// Recursively call detect() for each child
	for _, item := range items {
		if item[0] == current {
			if detect(item[1]) {
				stack = stack[0 : len(stack)-1]
			}
		}
	}
	return true
}

func handleCycle(index int, current string) {
	from := index
	for i := index; i < len(stack); i++ {
		if stack[i] < current {
			current = stack[i]
			from = i
		}
	}

	var cycle string
	for count := 0; count < len(stack)-index+1; count++ {
		if from == len(stack) {
			from = index
		}
		cycle += stack[from]
		from++
	}

	for _, eachCycle := range cycleArr {
		if eachCycle == cycle {
			return
		}
	}
	cycleArr = append(cycleArr, cycle)
}
