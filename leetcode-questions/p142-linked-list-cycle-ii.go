/**
 * LeetCode : Linked List Cycle II
 * https://leetcode.com/problems/linked-list-cycle-ii/
 *
 * Given a linked list, return the node where the cycle begins. If there is no cycle, return null.
 * There is a cycle in a linked list if there is some node in the list that can be reached again by
 * continuously following the next pointer. Internally, pos is used to denote the index of the node that
 * tail's next pointer is connected to. Note that pos is not passed as a parameter.
 * Note that pos is not passed as a parameter.
 * submission : faster than 97%
 */
package main

func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	fast := head
	slow := head
	cycle := false
	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			cycle = true
			break
		}
	}
	if !cycle {
		return nil
	}

	for head != slow {
		head = head.Next
		slow = slow.Next
	}
	return head
}

