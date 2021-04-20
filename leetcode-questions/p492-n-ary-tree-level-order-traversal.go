/**
 * LeetCode : N-ary Tree Level Order Traversal
 * https://leetcode.com/problems/n-ary-tree-level-order-traversal/
 *
 * submission : faster than 100%
 */
package main

type Node struct {
	Val      int
	Children []*Node
}

func levelOrder(root *Node) [][]int {
	if root == nil {
		return [][]int{}
	}
	var res [][]int
	queue := []*Node{root}
	for len(queue) != 0 {
		var level []int
		var temp []*Node
		for _, node := range queue {
			level = append(level, node.Val)
			temp = append(temp, node.Children...)
		}
		res = append(res, level)
		queue = temp
	}
	return res
}
