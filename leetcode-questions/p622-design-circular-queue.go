/**
 * LeetCode : Design Circular Queue
 * https://leetcode.com/problems/design-circular-queue/
 *
 * submission : faster than 93%
 */
package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

type MyCircularQueue struct {
	Head *ListNode
	Tail *ListNode
	Cap  int
	Len  int
}

func Constructor(k int) MyCircularQueue {
	return MyCircularQueue{Cap: k}
}

func (this *MyCircularQueue) EnQueue(value int) bool {
	if this.Cap == this.Len {
		return false
	}
	node := &ListNode{Val: value}
	if this.Len == 0 {
		this.Head = node
		this.Tail = node
		this.Len = 1
	} else {
		this.Tail.Next = node
		node.Next = this.Head
		this.Tail = node
		this.Len++
	}
	return true
}

func (this *MyCircularQueue) DeQueue() bool {
	if this.Len == 0 {
		return false
	}
	
	this.Tail.Next = this.Head.Next
	this.Head = this.Head.Next
	this.Len--
	return true
}

func (this *MyCircularQueue) Front() int {
	if this.Len == 0 {
		return -1
	}
	return this.Head.Val
}

func (this *MyCircularQueue) Rear() int {
	if this.Len == 0 {
		return -1
	}
	return this.Tail.Val
}

func (this *MyCircularQueue) IsEmpty() bool {
	return this.Len == 0
}

func (this *MyCircularQueue) IsFull() bool {
	return this.Len == this.Cap
}

func main() {
	queue := Constructor(8)
	fmt.Println(queue.EnQueue(3))
	fmt.Println(queue.EnQueue(9))
	fmt.Println(queue.EnQueue(5))
	fmt.Println(queue.EnQueue(0))
	fmt.Println(queue.DeQueue())
	fmt.Println(queue.DeQueue())
	fmt.Println(queue.Rear())
	fmt.Println(queue.Front())
}