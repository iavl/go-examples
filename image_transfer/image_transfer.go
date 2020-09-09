package main

import (
	"fmt"
	"math"
	"sync"
)

type task struct {
	res [][]int
	arr [][]int
	i   int
	j   int
	m   int
	n   int
}

func getByIndex(arr [][]int, i, j, m, n int) (int, bool) {
	if i == -1 || j == -1 || i == m || j == n {
		return 0, false
	}
	return arr[i][j], true
}

func getAvg(arr [][]int, i, j, m, n int) int {
	var count int
	var sum int
	if t, ok := getByIndex(arr, i-1, j-1, m, n); ok {
		sum += t
		count++
	}
	if t, ok := getByIndex(arr, i-1, j, m, n); ok {
		sum += t
		count++
	}
	if t, ok := getByIndex(arr, i-1, j+1, m, n); ok {
		sum += t
		count++
	}
	if t, ok := getByIndex(arr, i, j-1, m, n); ok {
		sum += t
		count++
	}
	if t, ok := getByIndex(arr, i, j, m, n); ok {
		sum += t
		count++
	}
	if t, ok := getByIndex(arr, i, j+1, m, n); ok {
		sum += t
		count++
	}
	if t, ok := getByIndex(arr, i+1, j-1, m, n); ok {
		sum += t
		count++
	}
	if t, ok := getByIndex(arr, i+1, j, m, n); ok {
		sum += t
		count++
	}
	if t, ok := getByIndex(arr, i+1, j+1, m, n); ok {
		sum += t
		count++
	}
	r := math.Floor((float64)(sum) / (float64)(count))
	return int(r)
}

func calc(res [][]int, arr [][]int, m, n int) {
	ch := make(chan task, 100)

	// generator
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			t := task{res, arr, i, j, m, n}
			ch <- t
		}
	}

	var waitGroup sync.WaitGroup
	waitGroup.Add(5)
	for i := 0; i < 5; i++ {
		go func(ch chan task, wg *sync.WaitGroup) {
			for t := range ch {
				t.res[t.i][t.j] = getAvg(t.arr, t.i, t.j, t.m, t.n)
			}
			wg.Done()
		}(ch, &waitGroup)
	}

	go func() {
		// Wait for everything to be processed.
		waitGroup.Wait()

		close(ch)
	}()
}

func main() {
	var arr = [][]int{
		{1, 1, 1},
		{1, 0, 1},
		{1, 1, 1},
	}
	res := arr
	m, n := len(arr), len(arr[0])

	calc(res, arr, m, n)

	for _, item := range res {
		fmt.Println(item)
	}
}
