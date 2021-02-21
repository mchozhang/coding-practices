// LeetCode : Cousins in Binary Tree
// https://leetcode.com/problems/cousins-in-binary-tree/
// In a binary tree, the root node is at depth 0, and children of each depth k node are at depth k+1.
//
// Two nodes of a binary tree are cousins if they have the same depth, but have different parents.
//
// We are given the root of a binary tree with unique values, and the values x and y of two different nodes in the tree.
//
// Return true if and only if the nodes corresponding to the values x and y are cousins..
//
// Example:
// input: root = [1,2,3,4], x = 4, y = 3
// output: false
// faster than 100%

package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isCousins(root *TreeNode, x int, y int) bool {
	if root == nil {
		return false
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)

	for len(queue) != 0 {
		foundX := false
		foundY := false
		levelSize := len(queue)

		// traverse the nodes that are in the same level
		for i := 0; i < levelSize; i++ {
			current := queue[0]
			queue = queue[1:]
			if current.Val == x {
				foundX = true
			}
			if current.Val == y {
				foundY = true
			}
			if current.Left != nil && current.Right != nil {
				if current.Left.Val == x && current.Right.Val == y {
					return false
				}
				if current.Left.Val == y && current.Right.Val == x {
					return false
				}
			}
			if current.Left != nil {
				queue = append(queue, current.Left)
			}
			if current.Right != nil {
				queue = append(queue, current.Right)
			}
		}

		if foundX && foundY {
			return true
		} else if foundX || foundY {
			return false
		}
	}
	return false
}

func main() {
	node6 := &TreeNode{Val: 6}
	node7 := &TreeNode{Val: 7}
	node4 := &TreeNode{Val: 4}
	node5 := &TreeNode{Val: 5}
	node2 := &TreeNode{Val: 2, Left: node4, Right: node6}
	node3 := &TreeNode{Val: 3, Right: node5, Left: node7}
	node1 := &TreeNode{Val: 1, Left: node2, Right: node3}
	fmt.Println(isCousins(node1, 1 ,2))

}
