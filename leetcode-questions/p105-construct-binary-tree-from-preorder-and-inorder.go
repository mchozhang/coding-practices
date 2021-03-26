/**
 * LeetCode : Construct Binary Tree from Preorder and Inorder Traversal
 * https://leetcode.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/
 *
 * submission : faster than 96%
 */
package main
func buildTree(preorder []int, inorder []int) *TreeNode {
	return build(&preorder, inorder)
}

func build(preorder *[]int, inorder []int) *TreeNode {
	if len(*preorder) == 0 {
		return nil
	}
	node := &TreeNode{Val: (*preorder)[0]}
	*preorder = (*preorder)[1:]
	for i, num := range inorder {
		if num == node.Val {
			node.Left = build(preorder, inorder[:i])
			node.Right = build(preorder, inorder[i+1:])
			return node
		}
	}
	return nil
}
