package main

import (
	"fmt"
)

func ProblemA() {
	var N, Y int
	fmt.Scanf("%d %d", &N, &Y)

	A := 0 // 10,000円札
	B := 0 // 5,000円札
	C := 0 // 1,000円札

	for i := 0; i <= N; i++ {
		for j := 0; j <= N; j++ {
			checkTotal := ((10_000 * i) + (5_000 * j) + (1_000 * (N - i - j))) == Y
			checkCount := i+j+(N-i-j) == N

			if checkTotal && checkCount && (N-i-j) >= 0 {
				A = i
				B = j
				C = N - i - j
				break
			}
		}
	}

	if (A == 0) && (B == 0) && (C == 0) {
		fmt.Println(-1, -1, -1)
	} else {
		fmt.Println(A, B, C)
	}
}
