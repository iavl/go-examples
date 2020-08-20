package main

import "fmt"

func maxSubArray(nums []int) int {
	dp := make([]int, len(nums))
	for i := range dp {
		dp[i] = -1
	}

	dp[0] = nums[0]
	res := dp[0]
	for i := 1; i < len(nums); i++ {
		if dp[i-1] < 0 {
			dp[i] = nums[i]
		} else {
			dp[i] = dp[i-1] + nums[i]
		}
		res = max(res, dp[i])
	}
	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {

	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println(maxSubArray(nums))
}
