package main

import (
	"fmt"
)

func findOddNumber(numberList []int) bool {
	for _, v := range numberList {
		if v%2 != 0 {
			return true
		}
	}

	return false
}

func mapFunc(numberList []int, f func(n int) int) []int {
	result := []int{}

	for _, v := range numberList {
		result = append(result, f(v))
	}

	return result
}

func ProblemA() {
	var N int
	fmt.Scan(&N)

	var A []int

	for i := 0; i < N; i++ {
		var a int
		fmt.Scan(&a)

		A = append(A, a)
	}

	result := 0

	for !findOddNumber(A) {
		A = mapFunc(A, func(n int) int {
			return n / 2
		})

		result += 1
	}

	fmt.Println(result)
}
