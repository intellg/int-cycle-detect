// Question：一个监狱，关押着50名囚犯，狱长要执行特赦。
// 他选出一个代表，告诉他，每天随机有一个犯人可以出来放风，这个犯人可以开关操场上的一盏灯。
// 需要代表调查：什么时候，这50名犯人全部放过风。
//
// Solution：设定灯的初始状态是关。
// 每名放风的犯人，如果第一次出来放风，并且发现灯是关着的，就打开，否则不做任何操作。
// 代表每次出来，如果灯是开着的，就计数+1，并且把灯关上，否则不做任何操作。
// 经过若干次（代表至少出来49次）放风之后，代表的计数达到49，则表明全部囚犯都放过风了。
//
// Further Investigation：计算概率
package main

import (
	"fmt"
	"math/big"
	"sort"
	"time"
)

type ValInv struct {
	invalid *big.Int
	valid   *big.Int
}

var dict = make(map[int]ValInv)

// 组合公式，计算C(n, m)
func comb(n int, m int) *big.Int {
	if n < m {
		return big.NewInt(0)
	}

	scope := m
	// 下面三行代码，是为了计算组合公式时，避免分子分母出现相同数字时，重复计算。
	// 性能提升20%（day=1500, person=200)
	if m > n/2 {
		scope = n - m
	}

	result := big.NewInt(1)
	for i := 1; i <= scope; i++ {
		result = new(big.Int).Mul(result, big.NewInt(int64(n-i+1)))
		result = new(big.Int).Div(result, big.NewInt(int64(i)))
	}
	return result
}

func rate(day int, person int) *big.Int {
	invalid := big.NewInt(0)
	for i := 1; i < person; i++ {
		result, ok := dict[person-i]
		subValid := result.valid
		if !ok {
			subValid = rate(day, person-i)
		}

		semiPossibility := comb(person, i)
		invalid = new(big.Int).Add(invalid, new(big.Int).Mul(semiPossibility, subValid))
	}

	total := new(big.Int).Exp(big.NewInt(int64(person)), big.NewInt(int64(day)), nil)
	valid := new(big.Int).Sub(total, invalid)
	dict[person] = ValInv{valid: valid, invalid: invalid}
	return valid
}

func output(day int, person int, startTime time.Time, endTime time.Time) {
	var keys []int
	for k := range dict {
		if k > person - 10 {
			keys = append(keys, k)
		}
	}
	sort.Ints(keys)

	for _, k := range keys {
		result := dict[k]
		valid := new(big.Float).SetInt(result.valid)
		invalid := new(big.Float).SetInt(result.invalid)
		total := new(big.Float).Add(valid, invalid)
		probability := new(big.Float).Quo(valid, total)
		fmt.Printf("person: %d\n  invalid: %e\n  valid:   %e\n  total:   %e\n  probability: %f\n", k, invalid, valid, total, probability)
	}
	fmt.Printf("Days: %d\n", day)
	fmt.Printf("Spent %d ms\n", endTime.Sub(startTime).Milliseconds())
}

func main() {
	day := 421
	person := 50

	fmt.Println("Calculating...")

	startTime := time.Now()
	rate(day, person)
	endTime := time.Now()

	output(day, person, startTime, endTime)
}
