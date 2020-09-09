package main

import "fmt"

func reverse(x int) int {
	ret := 0
	for x != 0 {
		if tmp := int32(ret); (tmp*10)/10 != tmp {
			return 0
		}
		ret = ret*10 + x%10
		x = x / 10
	}
	return ret
}

func main() {
	tests := []struct {
		input int
		exp   int
	}{
		{123, 321},
		{-123, -321},
		{120, 21},
	}

	for _, item := range tests {
		res := reverse(item.input)
		fmt.Println(fmt.Sprintf("input:%v, output:%v, expect:%v", item.input, res, item.exp))
	}
}
