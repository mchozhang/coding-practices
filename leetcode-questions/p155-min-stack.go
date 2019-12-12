/**
 * LeetCode : Min Stack
 * https://leetcode.com/problems/min-stack/
 * Design a stack that supports push, pop, top, and retrieving the minimum element in constant time.
 * push(x) -- Push element x onto stack.
 * pop() -- Removes the element on top of the stack.
 * top() -- Get the top element.
 * getMin() -- Retrieve the minimum element in the stack.
 *
 * submission : faster than 82.35%
 */

package main

import "fmt"

type MinStack struct {
	stack []int // store the difference of the value with the minimum
	min   int
}

func Constructor() MinStack {
	return MinStack{stack: []int{}}
}

func (this *MinStack) Push(x int) {
	if len(this.stack) == 0 {
		this.stack = append(this.stack, 0)
		this.min = x
	} else {
		this.stack = append(this.stack, x - this.min)
		if x < this.min {
			this.min = x
		}
	}
}

func (this *MinStack) Pop() {
	if len(this.stack) == 0 {
		return
	}

	top := this.stack[len(this.stack) - 1]
	this.stack = this.stack[:len(this.stack) - 1]

	if top < 0 {
		this.min = this.min - top
	}
}

func (this *MinStack) Top() int {
	top := this.stack[len(this.stack) - 1]
	if top < 0 {
		return this.min
	} else {
		return this.min + top
	}
}

func (this *MinStack) GetMin() int {
	return this.min
}

func main() {
	obj := Constructor()
	obj.Push(1) // [0], min:1
	obj.Push(2) // [0, -1], min:1
	obj.Push(-2) // [0, -1, -3], min:-2
	a := obj.Top()       // -2
	min1 := obj.GetMin() // -2
	obj.Pop()
	b := obj.Top()       // 2
	min2 := obj.GetMin() // 1
	fmt.Println(a, b, min1, min2)
}
