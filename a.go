package main

import (
	"fmt"
	"math"
	"sort"
)

func ProblemA() {
	var N int
	var A []int

	fmt.Scan(&N)

	scoreList := []int{0, 0}

	for i := 0; i < N; i++ {
		var a int
		fmt.Scan(&a)

		A = append(A, a)
	}

	// 降順にソート
	sort.Slice(A, func(i, j int) bool {
		return A[i] > A[j]
	})

	// スコアを計算
	for i := 0; i < N; i++ {
		if i%2 == 0 {
			scoreList[0] += A[i]
		} else {
			scoreList[1] += A[i]
		}
	}

	fmt.Println(int(math.Abs(float64(scoreList[0]) - float64(scoreList[1]))))
}
