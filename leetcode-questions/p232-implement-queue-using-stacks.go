/**
 * LeetCode : 232. Implement Queue using Stacks
 * https://leetcode.com/problems/implement-queue-using-stacks/
 *
 * submission : faster than 100%
 */
package main

type MyQueue struct {
	Input  []int
	Output []int
}

func Constructor() MyQueue {
	return MyQueue{}
}

func (this *MyQueue) Push(x int) {
	this.Input = append(this.Input, x)
}

func (this *MyQueue) Pop() int {
	this.Peek()
	res := this.Output[len(this.Output)-1]
	this.Output = this.Output[:len(this.Output)-1]
	return res
}

func (this *MyQueue) Peek() int {
	if len(this.Output) == 0 {
		for len(this.Input) != 0 {
			this.Output = append(this.Output, this.Input[len(this.Input)-1])
			this.Input = this.Input[:len(this.Input)-1]
		}
	}
	return this.Output[len(this.Output)-1]
}

func (this *MyQueue) Empty() bool {
	return len(this.Output) == 0 && len(this.Input) == 0
}
