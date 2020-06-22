// 给定一个数组和一个临时的交换位置，将数组中前n项移到末尾
package main

import "fmt"

var total = 0

func main() {
	m := 21
	n := 9
	slice := initialize(m)
	adjust(slice, m, n)

	fmt.Printf("Total %d", total)
}

func initialize(m int) []int {
	slice := make([]int, m)
	for i := 0; i < m; i++ {
		slice[i] = i
	}
	fmt.Println(slice)
	return slice
}

func adjust(slice []int, m int, n int) {
	fmt.Printf("adjust %d, %d\n", m, n)
	c := m / n
	for i := 0; i < c; i++ {
		if move(slice, m-n*i, n) {
			fmt.Println(slice)
		}
	}

	m2 := m % n
	if m2 == 0 {
		return
	}
	n2 := n % m2
	if n2 == 0 {
		return
	}
	adjust(slice, m2, n2)
}

func move(slice []int, m int, n int) bool{
	if m == n {
		return false
	}
	for i := n - 1; i >= 0; i-- {
		to := m - n + i
		temp := slice[to]
		slice[to] = slice[i]
		slice[i] = temp
		total++
	}
	return true
}
