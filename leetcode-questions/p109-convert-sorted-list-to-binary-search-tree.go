/**
 * LeetCode : Convert Sorted List to Binary Search Tree
 * https://leetcode.com/problems/convert-sorted-list-to-binary-search-tree/
 *
 * submission : faster than 34%
 */
package main

func sortedListToBST(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}

	var pre *ListNode
	slow := head
	fast := head
	for fast != nil && fast.Next != nil {
		pre = slow
		slow = slow.Next
		fast = fast.Next.Next
	}
	if pre != nil {
		pre.Next = nil
	} else {
		head = nil
	}
	node := &TreeNode{Val: slow.Val}
	node.Left = sortedListToBST(head)
	node.Right = sortedListToBST(slow.Next)
	return node
}