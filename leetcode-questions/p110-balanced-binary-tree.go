/**
 * LeetCode : Balanced Binary Tree
 * https://leetcode.com/problems/balanced-binary-tree/
 *
 * submission : faster than 80%
 */
package main

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	leftBal := isBalanced(root.Left)
	rightBal := isBalanced(root.Right)
	if leftBal && rightBal {
		leftDepth := getDepth(root.Left)
		rightDepth := getDepth(root.Right)
		if leftDepth - rightDepth > 1 || rightDepth - leftDepth > 1 {
			return false
		}
		return true
	}

	return false
}

func getDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftDepth := getDepth(root.Left)
	rightDepth := getDepth(root.Right)
	max := leftDepth
	if max < rightDepth {
		max = rightDepth
	}
	return 1 + max
}