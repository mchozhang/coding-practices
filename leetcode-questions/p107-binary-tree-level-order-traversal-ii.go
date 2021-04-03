/**
 * LeetCode : Binary Tree Level Order Traversal II
 * https://leetcode.com/problems/binary-tree-level-order-traversal-ii/
 *
 * submission : faster than 100%
 */
package main

func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	res := make([][]int, 0)
	queue := []*TreeNode{root}

	for len(queue) != 0 {
		size := len(queue)
		var level []int
		for i := 0; i < size; i++ {
			current := queue[0]
			queue = queue[1:]
			level = append(level, current.Val)
			if current.Left != nil {
				queue = append(queue, current.Left)
			}

			if current.Right != nil {
				queue = append(queue, current.Right)
			}
		}

		res = append(res, []int{})
		copy(res[1:], res)
		res[0] = level
	}
	return res
}
