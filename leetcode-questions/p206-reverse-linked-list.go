/**
 * LeetCode : reverse linked list
 * https://leetcode.com/problems/reverse-linked-list/
 * Reverse a singly linked list.
 *
 * Example:
 * Input: 1->2->3->4->5->NULL
 * Output: 5->4->3->2->1->NULL
 *
 * submission : faster than 100%
 */

package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	var pre *ListNode
	current := head
	next := head.Next

	for ; current != nil; {
		next = current.Next
		current.Next = pre
		pre = current
		current = next
	}
	return pre
}

func main() {
	node4 := &ListNode{Val: 4}
	node3 := &ListNode{Val: 3, Next: node4}
	node2 := &ListNode{Val: 2, Next: node3}
	node1 := &ListNode{Val: 1, Next: node2}

	for cur := node1; cur != nil; cur = cur.Next {
		fmt.Println(cur.Val)
	}

	node1 = reverseList(node1)
	for cur := node1; cur != nil; cur = cur.Next {
		fmt.Println(cur.Val)
	}
}
