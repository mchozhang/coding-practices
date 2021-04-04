/**
 * LeetCode: Remove Linked List Elements
 * https://leetcode.com/problems/remove-linked-list-elements/
 *
 * submission : faster than 85%
 */
package main

func removeElements(node *ListNode, val int) *ListNode {
	if node == nil {
		return nil
	}
	if node.Val == val {
		return removeElements(node.Next, val)
	} else {
		node.Next = removeElements(node.Next, val)
		return node
	}
}
