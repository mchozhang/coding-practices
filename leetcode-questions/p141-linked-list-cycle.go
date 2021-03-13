/**
 * LeetCode : Linked List Cycle
 * https://leetcode.com/problems/linked-list-cycle/
 * Given head, the head of a linked list, determine if the linked list has a cycle in it.
 * There is a cycle in a linked list if there is some node in the list that can be reached again by continuously following the next pointer. Internally, pos is used to denote the index of the node that tail's next pointer is connected to.
 * Note that pos is not passed as a parameter.
 * submission : faster than 58%
 */
package main

/**
 * O(n) space complexity
 */
func hasCycle(head *ListNode) bool {
	set := make(map[*ListNode] bool)
	for node := head; node != nil; node = node.Next {
		if _, isIn := set[node]; isIn {
			return true
		} else {
			set[node] = true
		}
	}
	return false
}

func hasCycle2(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	fast := head
	slow := head
	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			return true
		}
	}
	return false
}

