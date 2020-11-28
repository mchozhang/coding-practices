/**
 * LeetCode : Reverse Linked List II
 * https://leetcode.com/problems/reverse-linked-list-ii/
 * Reverse a linked list from position m to n. Do it in one-pass.
 * Note: 1 ≤ m ≤ n ≤ length of list.
 *
 * Example:
 * Input: 1->2->3->4->5->NULL, m = 2, n = 4
 * Output: 1->4->3->2->5->NULL
 *
 * submission : faster than 100%
 */

package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if head == nil {
		return nil
	}
	var start *ListNode
	current := head
	next := head.Next

	for i := 1 ; i < m; i++ {
		start = current
		current = current.Next
		next = current.Next
	}

	var pre *ListNode
	tail := current
	for i := m; i < n+1; i++ {
		next = current.Next
		current.Next = pre
		pre = current
		current = next
	}
	tail.Next = current
	if start != nil {
		start.Next = pre
	}
	if m > 1 {
		return head
	}
	return pre
}

func main() {
	node4 := &ListNode{Val: 4}
	node3 := &ListNode{Val: 3, Next: node4}
	node2 := &ListNode{Val: 2, Next: node3}
	node1 := &ListNode{Val: 1, Next: node2}

	// for cur := node1; cur != nil; cur = cur.Next {
	// 	fmt.Println(cur.Val)
	// }

	node1 = reverseBetween(node1, 2,4)
	for cur := node1; cur != nil; cur = cur.Next {
		fmt.Println(cur.Val)
	}
}

