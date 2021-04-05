/**
 * LeetCode: Delete Node in a Linked List
 * https://leetcode.com/problems/delete-node-in-a-linked-list/
 *
 * submission: faster than 100%
 */
package main

func deleteNode(node *ListNode) {
	for next := node.Next;true; {
		node.Val = next.Val
		if next.Next == nil {
			node.Next = nil
			return
		}
		node = node.Next
		next = node.Next
	}
}