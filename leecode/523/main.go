package main

import "fmt"

func checkSubarraySum(nums []int, k int) bool {
	mp := make(map[int]int)

	sum := 0
	for i, v := range nums {
		sum += v
		if k != 0 {
			sum = sum % k
		}
		if val, ok := mp[sum]; ok {
			if i-val > 1 {
				return true
			}
		} else {
			mp[sum] = i
		}
	}
	return false
}

func main() {
	//arr := []int{23, 2, 4, 6, 7}
	arr := []int{0, 0}
	k := 0
	fmt.Println(checkSubarraySum(arr, k))
}
