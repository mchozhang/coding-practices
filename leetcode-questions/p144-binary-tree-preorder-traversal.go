// LeetCode : Binary Tree Preorder Traversal
// https://leetcode.com/problems/binary-tree-preorder-traversal/
// Given the root of a binary tree, return the preorder traversal of its nodes' values.
//
// Example:
// input: root = [1,null,2,3]
// output: [1,2,3]
// faster than 100%

package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	result := make([]int, 0)
	stack := make([]*TreeNode, 0)
	stack = append(stack, root)
	for len(stack) != 0 {
		lastIndex := len(stack) - 1
		current := stack[lastIndex]
		stack = stack[:lastIndex]
		result = append(result, current.Val)
		if current.Right != nil {
			stack = append(stack, current.Right)
		}
		if current.Left != nil {
			stack = append(stack, current.Left)
		}
	}
	return result
}

func postorderTraversal(root *TreeNode) []int {
	result := make([]int, 0)
	stack := make([]*TreeNode, 0)
	current := root
	for len(stack) != 0 || current != nil {
		// current is non-leaf
		for current != nil && (current.Left != nil || current.Right != nil) {
			stack = append(stack, current)
			current = current.Left
		}
		if current != nil {
			result = append(result, current.Val)
		}
		for len(stack) != 0 && current == stack[len(stack)-1].Right {
			current = stack[len(stack) - 1]
			stack = stack[:len(stack) - 1]
			result = append(result, current.Val)
		}
		if len(stack) == 0 {
			current = nil
		} else{
			current = stack[len(stack) - 1].Right
		}
	}
	return result
}

func main() {
	//    1
	//  2   3
	// 4 5 6 7
	node6 := &TreeNode{Val: 6}
	node7 := &TreeNode{Val: 7}
	node4 := &TreeNode{Val: 4}
	node5 := &TreeNode{Val: 5}
	node2 := &TreeNode{Val: 2, Left: node4, Right: node5}
	node3 := &TreeNode{Val: 3, Left: node6, Right: node7}
	node1 := &TreeNode{Val: 1, Left: node2, Right: node3}
	fmt.Println(postorderTraversal(node1))
}
