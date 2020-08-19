package main

import (
	"fmt"
	"sort"
)

var result [][]int = make([][]int, 0)
var path []int = make([]int, 0)

func main() {
	candidates := []int{2, 3, 7, 6}
	target := 7

	combinationSum(candidates, target)
	fmt.Println(result)
}

func combinationSum(candidates []int, target int) {
	sort.Ints(candidates)
	DFS(candidates, 0, target)
}

func DFS(candidates []int, start, target int) {
	fmt.Println(fmt.Sprintf("target: %v", target))

	if target == 0 {
		tmp := make([]int, len(path))
		copy(tmp, path)
		result = append(result, tmp)
		return
	}

	for i := start; i < len(candidates); i++ {
		fmt.Println(fmt.Sprintf("target: %v", target))

		if target-candidates[i] >= 0 {
			path = append(path, candidates[i])
			fmt.Println(fmt.Sprintf("append %v, path %v", candidates[i], path))
			DFS(candidates, i, target-candidates[i])
			p := path[len(path)-1]
			path = path[:len(path)-1]
			fmt.Println(fmt.Sprintf("pop %v, path %v", p, path))
		} else {
			fmt.Println(fmt.Sprintf("target [%v] < [%v], continue", target, candidates[i]))
		}
	}

	/*
		for i := start; i < len(candidates) && target-candidates[i] >= 0; i++ {
			path = append(path, candidates[i])
			fmt.Println(fmt.Sprintf("append %v, path %v", candidates[i], path))
			DFS(candidates, i, target-candidates[i])
			p := path[len(path)-1]
			path = path[:len(path)-1]
			fmt.Println(fmt.Sprintf("pop %v, path %v", p, path))
		}

	*/
}
