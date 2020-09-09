package main

import (
	"fmt"
	"sort"
)

func largestSumAfterKNegations(A []int, K int) int {
	sort.Ints(A)

	minIdx := 0
	for K > 0 {
		A[minIdx] = -A[minIdx]
		if minIdx < len(A)-1 && A[minIdx] > A[minIdx+1] {
			minIdx++
		}
		K--
	}

	var sum int
	for _, v := range A {
		sum += v
	}
	return sum
}

func main() {
	tests := []struct {
		input []int
		k     int
		exp   int
	}{
		{[]int{2, -3, -1, 5, -4}, 2, 13},
		{[]int{4, 2, 3}, 1, 5},
		{[]int{3, -1, 0, 2}, 3, 6},
	}

	for _, test := range tests {
		input := test.input
		k := test.k
		exp := largestSumAfterKNegations(input, k)
		if exp == test.exp {
			fmt.Println("passed")
		}
	}
}
