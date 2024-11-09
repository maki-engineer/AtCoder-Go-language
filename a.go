package main

import (
	"fmt"
	"sort"
)

func ProblemA() {
	var N int
	var D []int
	fmt.Scan(&N)

	result := 0
	n := 0

	for i := 0; i < N; i++ {
		var d int

		fmt.Scan(&d)

		D = append(D, d)
	}

	// 降順にソート
	sort.Slice(D, func(i, j int) bool {
		return D[i] > D[j]
	})

	for _, v := range D {
		if n == v {
			continue
		}

		n = v
		result++
	}

	fmt.Println(result)
}
