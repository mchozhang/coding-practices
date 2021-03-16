/**
 * LeetCode : Sum of Left Leaves
 * https://leetcode.com/problems/sum-of-left-leaves/
 *
 * submission : faster than 95%
 */
package main

func sumOfLeftLeaves(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left != nil && root.Left.Left == nil && root.Left.Right == nil {
		return root.Left.Val + sumOfLeftLeaves(root.Right)
	}
	return sumOfLeftLeaves(root.Left) + sumOfLeftLeaves(root.Right)
}
