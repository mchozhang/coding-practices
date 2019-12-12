/**
 * LeetCode : Remove Duplicates from Sorted List
 * https://leetcode.com/problems/remove-duplicates-from-sorted-list/
 * Given a sorted linked list, delete all duplicates such that each element appear only once.
 *
 * Example:
 * Input: 1->1->2
 * Output: 1->2
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

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	for current := head; current.Next != nil; {
		next := current.Next

		// duplicate found
		if current.Val == next.Val {
			if next.Next != nil {
				current.Next = next.Next
			} else {
				current.Next = nil
			}
		}

		// move on if the next value differs
		if current.Val != next.Val {
			current = current.Next
		}
	}

	return head
}

func main() {
	var first = [] int{2, 3, 3, 3}

	var node = &ListNode{first[0], nil}
	var head = node

	for _, v := range first[1:] {
		node.Next = &ListNode{v, nil}
		node = node.Next
	}

	var result = deleteDuplicates(head)
	for i := 0; result != nil; i++ {
		fmt.Println(result.Val)
		result = result.Next
	}
}
