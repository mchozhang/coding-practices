/**
 * LeetCode : Binary Tree Inorder Traversal
 * https://leetcode.com/problems/binary-tree-inorder-traversal/
 *
 * submission : faster than 99%
 */
package main

func inorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	traverse(root, &res)
	return res
}

func traverse(node *TreeNode, result *[]int) {
	if node == nil {
		return
	}
	traverse(node.Left, result)
	*result = append(*result, node.Val)
	traverse(node.Right, result)
}