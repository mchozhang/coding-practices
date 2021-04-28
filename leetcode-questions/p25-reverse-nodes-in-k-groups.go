/**
 * LeetCode : Reverse Nodes in k-Group
 * https://leetcode.com/problems/reverse-nodes-in-k-group/
 *
 * submission : faster than 95%
 */
package main

func reverseKGroup(head *ListNode, k int) *ListNode {
	nextGroupHead := head
	var tail *ListNode
	count := 0
	for count < k && nextGroupHead != nil {
		tail = nextGroupHead
		nextGroupHead = nextGroupHead.Next
		count++
	}
	if count != k {
		return head
	}
	nextGroupHead = reverseKGroup(nextGroupHead, k)
	tail.Next = nextGroupHead

	pre := nextGroupHead
	current := head
	for current != nextGroupHead {
		next := current.Next
		current.Next = pre
		pre = current
		current = next
	}
	return pre
}
