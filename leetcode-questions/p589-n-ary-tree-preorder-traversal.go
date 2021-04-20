/**
 * LeetCode : N-ary Tree Preorder Traversal
 * https://leetcode.com/problems/n-ary-tree-preorder-traversal/
 *
 * submission : faster than 100%
 */
package main

type Node struct {
	Val      int
	Children []*Node
}

func preorder(root *Node) []int {
	var res []int
	traverse(root, &res)
	return res
}

func traverse(node *Node, res *[]int) {
	if node == nil {
		return
	}
	*res = append(*res, node.Val)
	for _, child := range node.Children {
		traverse(child, res)
	}
}