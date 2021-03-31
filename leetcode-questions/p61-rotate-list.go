/**
 * LeetCode : Rotate List
 * https://leetcode.com/problems/rotate-list/
 *
 * submission : faster than 100%
 */
package main

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	count := 1
	tail := head
	for ;tail.Next != nil; tail = tail.Next {
		count++
	}
	k %= count
	if k == 0 {
		return head
	}

	node := head
	for i := 0; i < count - k - 1; i++ {
		node = node.Next
	}
	res := node.Next
	node.Next = nil
	tail.Next = head
	return res
}
