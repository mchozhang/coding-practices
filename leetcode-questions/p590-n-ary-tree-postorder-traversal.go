/**
 * LeetCode : N-ary Tree Postorder Traversal
 * https://leetcode.com/problems/n-ary-tree-postorder-traversal/
 *
 * submission : faster than 100%
 */
package main

type Node struct {
	Val      int
	Children []*Node
}

func postorder(root *Node) []int {
	var res []int
	traverse(root, &res)
	return res
}

func traverse(node *Node, res *[]int) {
	if node == nil {
		return
	}
	for _, child := range node.Children {
		traverse(child, res)
	}
	*res = append(*res, node.Val)
}