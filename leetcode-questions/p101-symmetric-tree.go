/**
 * LeetCode : Symmetric Tree
 * https://leetcode.com/problems/symmetric-tree/
 * Given the root of a binary tree, check whether it is a mirror of itself (i.e., symmetric around its center).
 *
 * submission : faster than 100%
 */
package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return isSymmetricTree(root.Left, root.Right)
}

func isSymmetricTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p != nil && q != nil && p.Val == q.Val {
		return isSymmetricTree(p.Left, q.Right) && isSymmetricTree(p.Right, q.Left)
	}
	return false
}

func main() {
	n4 := &TreeNode{Val: 3}
	n5 := &TreeNode{Val:4}
	n6 := &TreeNode{Val: 4}
	n7 := &TreeNode{Val: 3}

	n3 := &TreeNode{1, n6, n7}
	n2 := &TreeNode{1, n4, n5}

	n1 := &TreeNode{1, n2,n3}
	fmt.Println(isSymmetric(n1))
}
