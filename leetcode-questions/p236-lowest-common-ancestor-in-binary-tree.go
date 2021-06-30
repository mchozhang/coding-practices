/**
 * LeetCode : Lowest Common Ancestor of a Binary Tree
 * https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-tree/
 *
 * submission : faster than 88.05%
 */
package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		return root
	}

	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)

	if left == nil && right == nil {
		return nil
	} else if left != nil && right == nil {
		return left
	} else if left == nil && right != nil {
		return right
	} else {
		return root
	}
}

func main() {
	node9 := &TreeNode{7, nil, nil}
	node10 := &TreeNode{4, nil, nil}

	node3 := &TreeNode{6, nil, nil}
	node4 := &TreeNode{2, node9, node10}
	node5 := &TreeNode{0, nil, nil}
	node6 := &TreeNode{8, nil, nil}

	node1 := &TreeNode{5, node3, node4}
	node2 := &TreeNode{1, node5, node6}
	node0 := &TreeNode{3, node1, node2}

	fmt.Println(lowestCommonAncestor(node0, node1, node2).Val)
	fmt.Println(lowestCommonAncestor(node0, node1, node10).Val)
	fmt.Println(lowestCommonAncestor(node0, node3, node6).Val)
	fmt.Println(lowestCommonAncestor(node0, node9, node10).Val)
}
