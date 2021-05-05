/**
 * LeetCode : Path Sum
 * https://leetcode.com/problems/path-sum-ii/
 *
 * submission : faster than 100%
 */
package main

type TreeNode struct {
	Left  *TreeNode
	Right *TreeNode
	Val   int
}

func pathSum(root *TreeNode, targetSum int) [][]int {
	res := make([][]int, 0)
	if root != nil {
		traverse(root, targetSum, 0, []int{}, &res)
	}
	return res
}

func traverse(root *TreeNode, targetSum int, currentSum int, currentPath []int, result *[][]int) {
	arr := append([]int{}, currentPath...)
	arr = append(arr, root.Val)
	currentSum += root.Val
	if root.Left == nil && root.Right == nil && currentSum == targetSum {
		*result = append(*result, arr)
		return
	}
	if root.Left != nil {
		traverse(root.Left, targetSum, currentSum, arr, result)
	}
	if root.Right != nil {
		traverse(root.Right, targetSum, currentSum, arr, result)
	}
}
