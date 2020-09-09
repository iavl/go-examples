package main

import (
	"fmt"
	"sort"
)

func pairSums(nums []int, target int) [][]int {
	sort.Ints(nums)

	res := make([][]int, 0)

	start := 0
	end := len(nums) - 1
	for start < end {
		sum := nums[start] + nums[end]
		if sum > target {
			end--
		} else if sum < target {
			start++
		} else {
			tmp := []int{nums[start], nums[end]}
			res = append(res, tmp)

			start++
			end--
		}
	}
	return res
}

func main() {
	tests := []struct {
		input []int
		k     int
	}{
		{[]int{5, 6, 5}, 11},
		{[]int{5, 6, 5, 6}, 11},
	}

	for _, test := range tests {
		input := test.input
		k := test.k
		exp := pairSums(input, k)
		fmt.Println(exp)
	}
}
