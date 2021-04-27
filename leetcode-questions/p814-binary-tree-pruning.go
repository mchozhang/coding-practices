/**
 * LeetCode : Binary Tree Pruning
 * https://leetcode.com/problems/binary-tree-pruning/
 *
 * submission : faster than 100%
 */
package main

func pruneTree(root *TreeNode) *TreeNode {
	if prune(root) {
		return root
	}
	return nil
}

func prune(node *TreeNode) bool {
	if node == nil {
		return false
	}

	left := prune(node.Left)
	right := prune(node.Right)
	if left == false {
		node.Left = nil
	}
	if right == false {
		node.Right = nil
	}
	return node.Val == 1 || left || right
}
