/**
 * LeetCode : Convert Sorted Array to Binary Search Tree
 * https://leetcode.com/problems/convert-sorted-array-to-binary-search-tree/
 *
 * submission : faster than 100%
 */
package main

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	mid := len(nums) / 2
	node := &TreeNode{Val: nums[mid]}
	node.Left = sortedArrayToBST(nums[:mid])
	node.Right = sortedArrayToBST(nums[mid+1:])
	return node
}
