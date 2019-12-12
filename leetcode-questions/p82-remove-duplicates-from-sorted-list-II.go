/**
 * LeetCode : Remove Duplicates from Sorted List II
 * https://leetcode.com/problems/remove-duplicates-from-sorted-list-ii/
 * Given a sorted linked list, delete all nodes that have duplicate numbers,
 * leaving only distinct numbers from the original list.
 *
 * Example:
 * Input: 1->2->3->3->4->4->5
 * Output: 1->2->5
 *
 * submission : faster than 89.70%
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

	var prev = &ListNode{0, head}
	var fakeHead = prev
	var duplicate *ListNode
	for current := head; current != nil; current = current.Next {
		// check the existing duplicate,
		if duplicate != nil && current.Val == duplicate.Val {
			// prune 3 or more consecutive duplicates
			prev.Next = current.Next
		} else if current.Next != nil && current.Val == current.Next.Val {
			// record the presence of the duplicate
			duplicate = current

			// prune double duplicates
			prev.Next = current.Next
		} else {
			// no duplicate, current node will be kept
			duplicate = nil
			prev = current
		}
	}

	return fakeHead.Next
}

func main() {
	var first = [] int{1, 2, 2,2, 3,4,4}

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
