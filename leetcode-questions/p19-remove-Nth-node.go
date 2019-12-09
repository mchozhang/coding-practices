/**
 * LeetCode : Remove Nth Node From End of List
 * https://leetcode.com/problems/remove-nth-node-from-end-of-list/
 * Given a linked list, remove the n-th node from the end of list and return its head.
 *
 * Example:
 * Given linked list: 1->2->3->4->5, and n = 2.
 * Output:  1->2->3->5.
 *
 * submission : faster than 100%
 */

package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	node := head
	nodes := make([]*ListNode, 0, n + 1)

	for ;node != nil; {
		if len(nodes) == n + 1 {
			nodes = append(nodes[1:], node)
		} else {
			nodes = append(nodes, node)
		}
		node = node.Next
	}

	length := len(nodes)
	if length - n == 0 {
		// remove first node or remove the only node
		return head.Next
	} else if length - n + 1 == length {
		// remove last node
		nodes[length - n - 1].Next = nil
	} else if length - n > 0 {
		// remove node in middle
		nodes[length - n - 1].Next = nodes[length - n + 1]
	}

	return head
}

func main() {
	var first = [] int{1, 2, 3}
	var node = &ListNode{first[0], nil}
	var head = node

	for _, v := range first[1:] {
		node.Next = &ListNode{v, nil}
		node = node.Next
	}

	var result = removeNthFromEnd(head, 3)
	for ; result != nil; {
		fmt.Println(result.Val)
		result = result.Next
	}
}
