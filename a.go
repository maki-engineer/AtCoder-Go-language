package main

import (
	"fmt"
	"strconv"
	"strings"
)

func mapFunc(strArray []string, f func(strNumber string) int) []int {
	result := []int{}

	for _, v := range strArray {
		result = append(result, f(v))
	}

	return result
}

func sum(numberList []int) int {
	result := 0

	for _, v := range numberList {
		result += v
	}

	return result
}

func ProblemA() {
	var N, A, B int
	result := 0

	fmt.Scanf("%d %d %d", &N, &A, &B)

	for i := 1; i <= N; i++ {
		s := strconv.Itoa(i)

		numberList := mapFunc(strings.Split(s, ""), func(strNumber string) int {
			v, _ := strconv.Atoi(strNumber)

			return v
		})

		sumTotal := sum(numberList)

		if (A <= sumTotal) && (sumTotal <= B) {
			result += i
		}
	}

	fmt.Println(result)
}
