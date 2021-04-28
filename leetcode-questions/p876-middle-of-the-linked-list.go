/**
 * LeetCode : Middle of the Linked List
 * https://leetcode.com/problems/middle-of-the-linked-list/
 *
 * submission : faster than 100%
 */
package main

func middleNode(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}