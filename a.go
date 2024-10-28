package main

import (
	"fmt"
)

func ProblemA() {
	var a, b, c int
	var s string

	fmt.Scanf("%d", &a)
	fmt.Scanf("%d %d", &b, &c)
	fmt.Scanf("%s", &s)

	result := a + b + c

	fmt.Println(result, s)
}
