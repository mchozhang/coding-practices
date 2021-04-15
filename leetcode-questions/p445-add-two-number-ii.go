/**
 * LeetCode : Add Two Numbers II
 * https://leetcode.com/problems/add-two-numbers-ii/
 *
 * submission : faster than 90%
 */
package main

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	node1, node2 := l1, l2
	len1, len2 := 0, 0
	for node1 != nil {
		node1 = node1.Next
		len1++
	}

	for node2 != nil {
		node2 = node2.Next
		len2++
	}

	node1, node2 = l1, l2
	if len1 < len2 {
		len1, len2 = len2, len1
		node1, node2 = node2, node1
	}

	node1 = addTwoNodes(node1, node2, len1-len2,0)
	if node1.Val > 9 {
		node1.Val -= 10
		return &ListNode{1, node1}
	}
	return node1
}

func addTwoNodes(node1 *ListNode, node2 *ListNode, diff int, pos int) *ListNode {
	if node1 == nil {
		return nil
	}

	var next *ListNode
	if pos < diff {
		next = addTwoNodes(node1.Next, node2, diff, pos+1)
	} else {
		next = addTwoNodes(node1.Next, node2.Next, diff, pos+1)
		node1.Val += node2.Val
	}

	if next != nil {
		node1.Val += next.Val/10
		next.Val %= 10
	}
	return node1
}