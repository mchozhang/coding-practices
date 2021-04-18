/**
 * LeetCode : Minimum Absolute Difference in BST
 * https://leetcode.com/problems/minimum-absolute-difference-in-bst/
 *
 * submission : faster than 74%
 */
package main

import "math"

func getMinimumDifference(node *TreeNode) int {
	min, preVal := math.MaxInt32, -1
	inorder(node, &preVal, &min)
	return min
}

func inorder(node *TreeNode, preVal *int, min *int) {
	if node.Left != nil {
		inorder(node.Left, preVal, min)
	}

	if *preVal >= 0 && node.Val - *preVal < *min {
		*min = node.Val - *preVal
	}
	*preVal = node.Val
	if node.Right != nil {
		inorder(node.Right, preVal, min)
	}
}

func getMinimumDifference2(node *TreeNode) int {
	_, _, diff := findMaxMin(node)
	return diff
}

func findMaxMin(node *TreeNode) (int, int, int) {
	max, min, diff := node.Val, node.Val, math.MaxInt16
	if node.Left != nil {
		leftMax, leftMin, leftDiff := findMaxMin(node.Left)
		min = leftMin
		diff = leftDiff
		if node.Val-leftMax < leftDiff {
			diff = node.Val - leftMax
		}
	}
	if node.Right != nil {
		rightMax, rightMin, rightDiff := findMaxMin(node.Right)
		max = rightMax
		if rightMin-node.Val < rightDiff {
			rightDiff = rightMin - node.Val
		}
		if rightDiff < diff {
			diff = rightDiff
		}
	}
	return max, min, diff
}
