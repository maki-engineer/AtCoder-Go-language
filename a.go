package main

import (
	"fmt"
)

func ProblemA() {
	var A, B, C, X int
	result := 0

	fmt.Scan(&A)
	fmt.Scan(&B)
	fmt.Scan(&C)
	fmt.Scan(&X)

	for a := 0; a <= A; a++ {
		for b := 0; b <= B; b++ {
			for c := 0; c <= C; c++ {
				totalAmount := (500 * a) + (100 * b) + (50 * c)
				if totalAmount == X {
					result++
				}
			}
		}
	}

	fmt.Println(result)
}
