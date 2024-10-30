package main

import (
	"fmt"
	"strings"
)

func ProblemA() {
	var s string
	fmt.Scanf("%s", &s)

	marbleList := strings.Split(s, "")
	result := 0

	for _, v := range marbleList {
		if v == "1" {
			result += 1
		}
	}

	fmt.Println(result)
}
