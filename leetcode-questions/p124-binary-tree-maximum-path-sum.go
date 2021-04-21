/**
 * LeetCode : Binary Tree Maximum Path Sum
 * https://leetcode.com/problems/binary-tree-maximum-path-sum/
 *
 * submission : faster than 88%
 */
package main

import "math"

func maxPathSum(node *TreeNode) int {
	max := math.MinInt16
	maxPathDown(node, &max)
	return max
}

func maxPathDown(node *TreeNode, max *int) int {
	if node == nil {
		return 0
	}
	left := maxPathDown(node.Left, max)
	if left < 0 {
		left = 0
	}
	right := maxPathDown(node.Right, max)
	if right < 0 {
		right = 0
	}
	if *max < left + right + node.Val {
		*max = left + right + node.Val
	}
	if left > right {
		return node.Val + left
	}
	return node.Val + right
}

/**
 * 200ms
 */
func maxPathSum2(node *TreeNode) int {
	if node == nil {
		return 0
	}

	maxPath := math.MinInt16
	longestBranch := node.Val
	var leftBranch, rightBranch int

	if node.Left != nil {
		maxPath = maxPathSum(node.Left)
		leftBranch = maxBranchSum(node.Left)
		if leftBranch < 0 {
			leftBranch = 0
		}
		longestBranch += leftBranch
	}

	if node.Right != nil {
		rightPathSum := maxPathSum(node.Right)
		if maxPath < rightPathSum {
			maxPath = rightPathSum
		}
		rightBranch = maxBranchSum(node.Right)
		if rightBranch < 0 {
			rightBranch = 0
		}
		longestBranch += rightBranch
	}

	if longestBranch > maxPath {
		return longestBranch
	}
	return maxPath
}

func maxBranchSum(node *TreeNode) int {
	var max int
	if node.Left != nil {
		max = maxBranchSum(node.Left)
	}
	if node.Right != nil {
		right := maxBranchSum(node.Right)
		if right > max {
			max = right
		}
	}

	if max > 0 {
		return node.Val + max
	} else {
		return node.Val
	}
}
