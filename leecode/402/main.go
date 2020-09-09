package main

import "fmt"

func isUnique(astr string) bool {
	mp := make(map[rune]int, 0)

	for _, v := range astr {
		if mp[v] > 0 {
			return false
		}
		mp[v]++
	}
	return true
}

func main() {
	s := "leetcode"
	fmt.Println(isUnique(s))
}
