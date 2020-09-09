package main

import "fmt"

type MinStack struct {
	stack    []int
	minStack []int
}

func Constructor() MinStack {
	return MinStack{
		stack:    make([]int, 0),
		minStack: make([]int, 0),
	}
}
func (this *MinStack) Push(x int) {
	this.stack = append(this.stack, x)
	if len(this.minStack) == 0 || x < this.minStack[len(this.minStack)-1] {
		this.minStack = append(this.minStack, x)
	} else {
		this.minStack = append(this.minStack, this.minStack[len(this.minStack)-1])
	}
}

func (this *MinStack) Pop() {
	this.stack = this.stack[:len(this.stack)-1]
	this.minStack = this.minStack[:len(this.minStack)-1]
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	if len(this.minStack) == 0 {
		return nil
	}
	return this.minStack[len(this.minStack)-1]

}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
func main() {
	obj := Constructor()
	fmt.Println(obj.GetMin())
	obj.Push(-2)
	fmt.Println(obj.GetMin())
	obj.Push(0)
	fmt.Println(obj.GetMin())

}
