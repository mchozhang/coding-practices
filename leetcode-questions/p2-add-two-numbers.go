/**
 * LeetCode : add two numbers
 * https://leetcode.com/problems/add-two-numbers/
 * You are given two non-empty linked lists representing two non-negative integers.
 * The digits are stored in reverse order and each of their nodes contain a single digit.
 * Add the two numbers and return it as a linked list.
 * You may assume the two numbers do not contain any leading zero, except the number 0 itself.
 * Example:
 * Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
 * Output: 7 -> 0 -> 8
 * Explanation: 342 + 465 = 807.
 *
 * submission : faster than 100%
 */

package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var node = &ListNode{0, nil}
	var head = node
	var carry = 0

	for {
		var sum, val1, val2 int

		if l1 != nil && l2 != nil {
			val1 = l1.Val
			val2 = l2.Val
			l1 = l1.Next
			l2 = l2.Next
		} else if l1 != nil {
			val1 = l1.Val
			val2 = 0
			l1 = l1.Next
		} else if l2 != nil {
			val1 = 0
			val2 = l2.Val
			l2 = l2.Next
		} else if carry == 0 {
			break
		}

		sum = val1 + val2 + carry
		carry = 0
		if sum >= 10 {
			carry = 1
			sum = sum % 10
		}
		node.Val = sum

		if l1 != nil || l2 != nil || carry == 1 {
			node.Next = &ListNode{}
			node = node.Next
		}
	}

	return head
}

func main() {
	var first = [] int{2, 4, 3}
	var second = [] int{5, 6, 4}

	var node1 = &ListNode{first[0], nil}
	var node2 = &ListNode{second[0], nil}
	var head1 = node1
	var head2 = node2

	for _, v := range first[1:] {
		node1.Next = &ListNode{v, nil}
		node1 = node1.Next
	}

	for _, v := range second[1:] {
		node2.Next = &ListNode{v, nil}
		node2 = node2.Next
	}

	var result = addTwoNumbers(head1, head2)
	for i := 0; result != nil; i++ {
		fmt.Println(result.Val)
		result = result.Next
	}
}
