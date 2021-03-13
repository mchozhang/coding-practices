/**
 * LeetCode : Diameter of Binary Tree
 * https://leetcode.com/problems/diameter-of-binary-tree/
 *
 * submission : faster than 97%
 */
package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {
	_, diameter := recursion(root)
	return diameter
}

func recursion(root *TreeNode) (int, int) {
	if root == nil {
		return 0, 0
	}
	leftDepth, leftDiameter := recursion(root.Left)    // 1, 2
	rightDepth, rightDiameter := recursion(root.Right) // 0,1
	maxDepth := leftDepth
	if leftDiameter == 0 && root.Left != nil {
		leftDiameter = 1
	}
	if rightDiameter == 0 && root.Right != nil {
		rightDiameter = 1
	}

	if leftDepth < rightDepth {
		maxDepth = rightDepth
	}
	maxDiameter := leftDiameter
	if leftDiameter < rightDiameter {
		maxDiameter = rightDiameter
	}
	diameter := leftDepth + rightDepth
	if maxDiameter > diameter {
		diameter = maxDiameter
	}
	return maxDepth + 1, diameter
}

func main() {
	node4 := &TreeNode{}
	node5 := &TreeNode{}
	node2 := &TreeNode{Left: node4, Right: node5}
	node6 := &TreeNode{}
	node3 := &TreeNode{Left: node6}
	node1 := &TreeNode{Left: node2, Right: node3}
	fmt.Println(diameterOfBinaryTree(node1))
}
