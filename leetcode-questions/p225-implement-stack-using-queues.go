/**
 * LeetCode : 225. Implement Stack using Queues
 * https://leetcode.com/problems/implement-stack-using-queues/
 *
 * submission : faster than 100%
 */
package main

type MyStack struct {
	Queue []int
}

func Constructor() MyStack {
	return MyStack{}
}

func (this *MyStack) Push(x int) {
	q := []int{x}
	q = append(q, this.Queue...)
	this.Queue = q
}

func (this *MyStack) Pop() int {
	value := this.Queue[0]
	this.Queue = this.Queue[1:]
	return value
}

func (this *MyStack) Top() int {
	return this.Queue[0]
}

func (this *MyStack) Empty() bool {
	return len(this.Queue) == 0
}
