// 据说这是Amazon的面试问题
// 在一个二维数组中0代表节点，1代表隔离物
// 如果一个节点0的上下左右为0，那么这两个0为连续空间
// 问题：找出全部连续的空间，并且找到每个连续空间中，到达各个节点的距离最短的节点（可能是多个），以及这个距离的值
package main

import (
	"fmt"
	"math"
)

type point struct {
	X int
	Y int
}

const (
	x = 6
	y = 6
)

var array = [x][y]int{
	{0, 0, 0, 0, 0, 0},
	{0, 1, 1, 1, 1, 0},
	{0, 1, 1, 1, 1, 0},
	{0, 1, 1, 1, 1, 0},
	{0, 1, 1, 1, 1, 0},
	{0, 0, 0, 0, 0, 0},
}

func main() {
	blocks := detectBlocks()
	for i, block := range blocks {
		fmt.Printf("<%d> %v\n", i, block)
	}

	fmt.Println("================================================================")

	for i, block := range blocks {
		middlePoints, length := detectPath(block)
		fmt.Printf("<%d> middle points %v, length is %d\n", i, middlePoints, length-1)
	}
}

func detectBlocks() (blocks []map[point]int) {
	for i := 0; i < x; i++ {
		for j := 0; j < y; j++ {
			if array[i][j] == 0 {
				p := point{i, j}
				isNewBlock := true
				for _, b := range blocks {
					if _, ok := b[p]; ok {
						isNewBlock = false
						break
					}
				}
				if isNewBlock {
					block := make(map[point]int)
					blocks = append(blocks, block)

					putPointIntoBlock(i, j, p, block)
				}
			}
		}
	}

	return
}

func checkPoint(i, j int, block map[point]int) {
	if i < 0 || j < 0 || i >= x || j >= y {
		return
	}

	if array[i][j] == 0 {
		p := point{i, j}
		if _, ok := block[p]; !ok {
			putPointIntoBlock(i, j, p, block)
		}
	}
}

func putPointIntoBlock(i, j int, p point, block map[point]int) {
	block[p] = 0
	checkPoint(i, j+1, block)
	checkPoint(i+1, j, block)
	checkPoint(i, j-1, block)
	checkPoint(i-1, j, block)
}

func detectPath(block map[point]int) (middlePoints []point, maxLength int) {
	maxLength = math.MaxInt64
	for p := range block {
		newBlock := make(map[point]int, len(block))
		for k, v := range block {
			newBlock[k] = v
		}
		length := calculateLength(newBlock, p)

		if length < maxLength {
			maxLength = length
			middlePoints = make([]point, 0)
			middlePoints = append(middlePoints, p)
		} else if length == maxLength {
			middlePoints = append(middlePoints, p)
		}
	}
	return
}

func calculateLength(block map[point]int, p point) (maxLength int) {
	block[p] = 1
	markAroundPoints(block, 2, p)

	for _, val := range block {
		if val > maxLength {
			maxLength = val
		}
	}
	return
}

func markAroundPoints(block map[point]int, length int, points ...point) {
	newPoints := make([]point, 0, 4*len(points))
	for _, p := range points {
		newPoint := point{p.X, p.Y + 1}
		if markPoint(block, length, newPoint) {
			newPoints = append(newPoints, newPoint)
		}
		newPoint = point{p.X + 1, p.Y}
		if markPoint(block, length, newPoint) {
			newPoints = append(newPoints, newPoint)
		}
		newPoint = point{p.X, p.Y - 1}
		if markPoint(block, length, newPoint) {
			newPoints = append(newPoints, newPoint)
		}
		newPoint = point{p.X - 1, p.Y}
		if markPoint(block, length, newPoint) {
			newPoints = append(newPoints, newPoint)
		}
	}
	if len(newPoints) > 0 {
		markAroundPoints(block, length+1, newPoints...)
	}
}

func markPoint(block map[point]int, length int, p point) bool {
	if p.X < 0 || p.Y < 0 || p.X >= x || p.Y >= y {
		return false
	}

	if val, ok := block[p]; ok {
		if val == 0 {
			block[p] = length
			return true
		}
	}
	return false
}
