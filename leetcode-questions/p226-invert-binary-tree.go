/**
 * LeetCode : Invert Binary Tree
 * https://leetcode.com/problems/invert-binary-tree/
 * Invert a binary tree
 *
 * submission : faster than 100%
 */
package main

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	var tmpLeft *TreeNode
	var tmpRight *TreeNode
	if root.Left != nil {
		tmpLeft = invertTree(root.Left)
	}
	if root.Right != nil {
		tmpRight = invertTree(root.Right)
	}
	root.Right = tmpLeft
	root.Left = tmpRight
	return root
}

