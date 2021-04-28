/**
 * LeetCode : Delete Node in a BST
 * https://leetcode.com/problems/delete-node-in-a-bst/
 *
 * submission : faster than 85%
 */
package main

func deleteNode(node *TreeNode, key int) *TreeNode {
	if node == nil {
		return nil
	}
	if key > node.Val {
		node.Right = deleteNode(node.Right, key)
	} else if key < node.Val {
		node.Left = deleteNode(node.Left, key)
	} else {
		if node.Left == nil {
			return node.Right
		}
		if node.Right == nil {
			return node.Left
		}
		rightMin := node.Right
		for rightMin.Left != nil {
			rightMin = rightMin.Left
		}
		rightMin.Left = node.Left
		return node.Right
	}
	return node
}
