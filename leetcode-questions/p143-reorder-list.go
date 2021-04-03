/**
 * LeetCode: Reorder List
 * https://leetcode.com/problems/reorder-list/
 *
 * submission : faster than 100%
 */
package main

func reorderList(head *ListNode)  {
	slow := head
	fast := head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	var pre *ListNode
	for slow != nil {
		next := slow.Next
		slow.Next = pre
		pre = slow
		slow = next
	}

	node := head
	node2 := pre
	for node2 != nil || node != nil {
		if node != nil {
			next := node.Next
			node.Next = node2
			node = next
		}

		if node2 != nil {
			next2 := node2.Next
			node2.Next = node
			node2 = next2
		}
	}
}