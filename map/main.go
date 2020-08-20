package main

import "fmt"

func main() {
	m := make(map[int]int)
	mdMap(m)
	fmt.Println(m)
}

func mdMap(m map[int]int) {
	m = make(map[int]int)
	m[1] = 100
	m[2] = 200
}
