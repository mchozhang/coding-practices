/**
 * LeetCode : Merge Two Binary Trees
 * https://leetcode.com/problems/merge-two-binary-trees/
 *
 * submission : faster than 99%
 */
package main

func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	var mergeNode *TreeNode 
	if root1 != nil && root2 != nil {
		mergeNode = &TreeNode{Val: root1.Val + root2.Val}
		mergeNode.Left = mergeTrees(root1.Left, root2.Left)
		mergeNode.Right = mergeTrees(root1.Right, root2.Right)

	} else if root1 != nil {
		mergeNode = root1
		mergeNode.Left = mergeTrees(root1.Left, nil)
		mergeNode.Right = mergeTrees(root1.Right, nil)
	} else if root2 != nil {
		mergeNode = root2
		mergeNode.Left = mergeTrees(nil, root2.Left)
		mergeNode.Right = mergeTrees(nil, root2.Right)
	}
	return mergeNode
}
