package main

import (
	"fmt"
	"math"
)

type MinStack struct {
	min   int
	stack []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	m := MinStack{}
	m.min = math.MaxInt32
	return m
}

func (this *MinStack) Push(x int) {
	if x < this.min {
		this.min = x
	}
	this.stack = append(this.stack, x)
}

func (this *MinStack) Pop() {
	// 如果pop的是最小元素，那么要找到第二小的
	pop := this.stack[len(this.stack)-1]
	this.stack = this.stack[:len(this.stack)-1]
	if len(this.stack) == 0 {
		this.min = math.MaxInt32
		return
	}
	if this.min == pop {
		this.min = this.stack[0]
		for _, v := range this.stack {
			if v < this.min {
				this.min = v
			}
		}
	}
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.min
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */

func main() {
	obj := Constructor()
	obj.Push(-2)
	obj.Push(0)
	obj.Push(-3)
	fmt.Println(obj.GetMin())
	obj.Pop()
	fmt.Println(obj.Top())
	fmt.Println(obj.GetMin())
}
