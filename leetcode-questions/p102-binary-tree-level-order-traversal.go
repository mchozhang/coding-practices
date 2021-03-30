/**
 * LeetCode : Binary Tree Level Order Traversal
 * https://leetcode.com/problems/binary-tree-level-order-traversal/
 * Given a binary tree, return the level order traversal of its nodes' values. (ie, from left to right, level by level).
 *
 * Example:
 * Given binary tree [3,9,20,null,null,15,7],
 *   3
 *  / \
 * 9  20
 *   /  \
 *  15   7
 *
 * output: [[3], [9,20], [15,7]]
 *
 * submission : faster than 40%
 */

package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	levelCounter := map[int]int{}
	traversal := make([][]int, 0)
	queue := make([]*TreeNode, 0)

	queue = append(queue, root)
	level := 0
	levelCounter[0] = 1
	for len(queue) != 0 {
		current := queue[0]
		queue = queue[1:]

		if len(traversal) <= level {
			traversal = append(traversal, []int{current.Val})
		} else {
			traversal[level] = append(traversal[level], current.Val)
		}

		if current.Left != nil {
			queue = append(queue, current.Left)
			levelCounter[level+1]++
		}
		if current.Right != nil {
			queue = append(queue, current.Right)
			levelCounter[level+1]++
		}

		if len(traversal[level]) == levelCounter[level] {
			level++
		}
	}

	return traversal
}

func main() {
	root := &TreeNode{0, nil, nil}
	root.Left = &TreeNode{1, nil, nil}
	root.Right = &TreeNode{2, nil, nil}

	node1 := root.Left
	node2 := root.Right

	node1.Left = &TreeNode{3, nil, nil}
	node2.Right = &TreeNode{8, nil, nil}

	result := levelOrder(root)
	fmt.Println(result)
}
