/**
 * LeetCode : Partition List
 * https://leetcode.com/problems/partition-list/
 *
 * submission : faster than 100%
 */
package main

func partition(head *ListNode, x int) *ListNode {
	dummy := &ListNode{Next: head}
	small := dummy
	pre := small
	current := head
	for current != nil {
		if current.Val < x {
			next := current.Next
			big := small.Next
			if small == pre {
				big = current.Next
				pre = current
			}
			small.Next = current
			small = current
			small.Next = big
			current = next
			pre.Next = current
		} else {
			pre = current
			current = current.Next
		}
	}
	return dummy.Next
}