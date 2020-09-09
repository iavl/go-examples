package main

import (
	"fmt"
	"math"
)

func minSubArrayLen(s int, nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	start, end := 0, 0
	sum := 0
	ans := math.MaxInt32
	for end < n {
		sum += nums[end]
		for sum >= s {
			ans = min(ans, end-start+1)
			sum -= nums[start]
			start++
		}
		end++
	}

	if ans == math.MaxInt32 {
		return 0
	}

	return ans
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func main() {
	tests := []struct {
		input []int
		s     int
		exp   int
	}{
		{[]int{2, 3, 1, 2, 4, 3}, 7, 2},
	}

	for _, test := range tests {
		input := test.input
		s := test.s
		exp := minSubArrayLen(s, input)
		if exp == test.exp {
			fmt.Println("passed")
		}
	}
}
