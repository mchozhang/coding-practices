/**
 * LeetCode : Range Sum of BST
 * https://leetcode.com/problems/range-sum-of-bst/
 *
 * submission : faster than 95%
 */
package main

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	} else if root.Left == nil {
		return 1 + minDepth(root.Right)
	} else if root.Right == nil {
		return 1 + minDepth(root.Left)
	} else {
		leftDepth := minDepth(root.Left)
		rightDepth := minDepth(root.Right)
		min := leftDepth
		if min > rightDepth {
			min = rightDepth
		}
		return 1 + min
	}
}

func minDepth2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)

	for i := 1; len(queue) != 0; i++ {
		size := len(queue)
		for j := 0; j < size; j++ {
			current := queue[0]
			if current.Left == nil && current.Right == nil {
				return i
			}
			if current.Left != nil {
				queue = append(queue, current.Left)
			}
			if current.Right != nil {
				queue = append(queue, current.Right)
			}
			queue = queue[1:]
		}
	}
	return 0
}
