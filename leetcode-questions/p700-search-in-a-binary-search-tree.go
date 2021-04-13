/**
 * LeetCode: Search in a Binary Search Tree
 * https://leetcode.com/problems/search-in-a-binary-search-tree/
 *
 * submission: faster than 100%
 */
package main

func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == val {
		return root
	}
	if root.Val < val {
		return searchBST(root.Right, val)
	}else{
		return searchBST(root.Left, val)
	}
}