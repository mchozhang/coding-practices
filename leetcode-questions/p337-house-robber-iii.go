/**
 * LeetCode: House Robber III
 * https://leetcode.com/problems/house-robber-iii/
 *
 * submission: faster than 99%
 */
package main

func rob(root *TreeNode) int {
	robCurrent, skipCurrent := robNode(root)
	return findMax(robCurrent, skipCurrent)
}

func robNode(node *TreeNode) (int,int) {
	if node == nil {
		return 0, 0
	}
	robLeft, skipLeft := robNode(node.Left)
	robRight, skipRight := robNode(node.Right)

	robCurrent := node.Val + skipLeft + skipRight
	skipCurrent := findMax(robLeft, skipLeft) + findMax(robRight, skipRight)

	return robCurrent, skipCurrent
}

func findMax(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
