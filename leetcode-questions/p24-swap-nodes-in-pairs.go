/**
 * LeetCode :Swap Nodes in Pairs
 * https://leetcode.com/problems/swap-nodes-in-pairs/
 *
 * submission : faster than 100%
 */
package main

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	pre := &ListNode{}
	dummy := pre
	current := head
	paired := head.Next
	next := paired.Next

	for true {
		pre.Next = paired
		paired.Next = current
		current.Next = next

		pre = current
		current = next
		if current != nil {
			paired = current.Next
		} else {
			break
		}

		if paired != nil {
			next = paired.Next
		} else {
			break
		}
	}
	return dummy.Next
}

func swapPairs2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	next := head.Next
	head.Next = swapPairs(next.Next)
	next.Next = head
	return next
}
