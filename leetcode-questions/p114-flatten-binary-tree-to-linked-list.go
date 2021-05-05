/**
 * LeetCode: Flatten Binary Tree to Linked List
 * https://leetcode.com/problems/flatten-binary-tree-to-linked-list/
 *
 * submission: faster than 75%
 */
package main

type TreeNode struct {
	Left  *TreeNode
	Right *TreeNode
	Val   int
}

func flatten(root *TreeNode) {
	if root != nil {
		root = traverse(root, nil)
	}
}

func traverse(node *TreeNode, next *TreeNode) *TreeNode {
	if node.Left != nil {
		node.Right = traverse(node.Left, node.Right)
		node.Left = nil
		traverse(node.Right, next)
	} else if node.Right != nil {
		node.Right = traverse(node.Right, next)
	} else {
		node.Right = next
	}
	return node
}
