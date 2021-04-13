/**
 * LeetCode: Count Complete Tree Nodes
 * https://leetcode.com/problems/count-complete-tree-nodes/
 *
 * submission: faster than 99%
 */
package main

import "math"

func countNodes(node *TreeNode) int {
	if node == nil {
		return 0
	}
	left := node.Left
	right := node.Right
	leftHeight := 1
	rightHeight := 1
	for right != nil {
		leftHeight++
		rightHeight++
		left = left.Left
		right = right.Right
	}

	if left == nil {
		return int(math.Pow(2, float64(leftHeight))) - 1
	}
	return countNodes(node.Left) + countNodes(node.Right) + 1
}

