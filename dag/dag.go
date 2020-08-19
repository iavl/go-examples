package main

import (
	"fmt"
)

func main() {
	var edge = map[string]string{
		"重庆": "新疆",
		"广州": "南京",
		"上海": "广州",
		"西藏": "重庆",
		"北京": "上海",
		"南京": "西藏",
	}

	var result []string
	visited := make(map[string]bool, 0)

	var visitAll func(items []string)
	visitAll = func(items []string) {
		fmt.Println(fmt.Sprintf("visit items: %v", items))
		for _, item := range items {
			if !visited[item] {
				visited[item] = true
				if v, ok := edge[item]; ok {
					fmt.Println(fmt.Sprintf("%s -> %s", item, v))
					visitAll([]string{v})
				}

				result = append(result, item)
				fmt.Println(fmt.Sprintf("append %s, result: %v", item, result))

			} else {
				fmt.Println(fmt.Sprintf("%s visited", item))
			}
		}

	}

	var keys []string
	for key := range edge {
		keys = append(keys, key)
	}

	fmt.Println("keys:", keys)
	visitAll(keys)

	reverse(result)
	fmt.Println(fmt.Sprintf("result: %v", result))
}

func reverse(arr []string) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}
