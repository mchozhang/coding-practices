/**
 * LeetCode : Maximum Depth of Binary Tree
 * https://leetcode.com/problems/maximum-depth-of-binary-tree/
 * Given the root of a binary tree, return its maximum depth.
 * submission : faster than 95%
 */
package main

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	max := maxDepth(root.Left)
	right := maxDepth(root.Right)
	if right > max {
		max = right
	}
	return 1 + max
}