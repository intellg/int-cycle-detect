package degree

import "math"

func InnerCalculateC(floor, cup int) (degree int) {
	dp := make([][]int, cup+1)
	for i := 0; i <= cup; i++ {
		dp[i] = make([]int, floor+1)
	}
	for i := 1; i <= floor; i++ {
		dp[1][i] = i
	}
	for i := 1; i <= cup; i++ {
		dp[i][1] = 1
	}

	for i := 2; i <= cup; i++ {
		for j := 2; j <= floor; j++ {
			res := math.MaxInt64
			for k := 1; k <= j; k++ {
				tmp := max(dp[i-1][k-1], dp[i][j-k])
				res = min(tmp, res)
				counter++
			}
			dp[i][j] = res + 1
		}
	}

	return dp[cup][floor]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
