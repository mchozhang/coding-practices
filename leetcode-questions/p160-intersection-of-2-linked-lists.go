/**
 * LeetCode : Intersection of Two Linked Lists
 * https://leetcode.com/problems/intersection-of-two-linked-lists/
 *
 * submission : faster than 86%
 */

package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	nodeA := headA
	nodeB := headB
	for nodeA != nil && nodeB != nil {
		if nodeA == nodeB {
			return nodeA
		}

		nodeA = nodeA.Next
		nodeB = nodeB.Next
	}

	longList := headA
	shortList := headB
	remain := nodeA
	diff := 0
	if nodeA == nil {
		remain = nodeB
		longList = headB
		shortList = headA
	}
	for remain != nil {
		diff++
		remain = remain.Next
	}

	for i := 0; i < diff; i++ {
		longList = longList.Next
	}
	for longList != nil && shortList != nil {
		if longList == shortList {
			return longList
		}
		longList = longList.Next
		shortList = shortList.Next
	}
	return nil
}
