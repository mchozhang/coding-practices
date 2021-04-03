/**
 * LeetCode : Construct Binary Tree from Inorder and Postorder Traversal
 * https://leetcode.com/problems/construct-binary-tree-from-inorder-and-postorder-traversal/
 *
 * submission : faster than 97%
 */
package main

func buildTree(inorder []int, postorder []int) *TreeNode {
	postIndex := len(postorder) - 1
	return build(&postIndex, 0, len(inorder)-1, inorder, postorder)
}

func build(postIndex *int, inorderLow int, inorderHigh int, inorder []int, postorder []int) *TreeNode {
	if *postIndex < 0 {
		return nil
	}
	value := postorder[*postIndex]
	node := &TreeNode{Val: value}
	for i := inorderLow; i <= inorderHigh; i++ {
		if inorder[i] == value {
			*postIndex--
			node.Right = build(postIndex, i+1, inorderHigh, inorder, postorder)
			node.Left = build(postIndex, inorderLow, i-1, inorder, postorder)
			return node
		}
	}
	return nil
}
