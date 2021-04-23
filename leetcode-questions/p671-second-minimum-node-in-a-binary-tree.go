/**
 * LeetCode : Second Minimum Node In a Binary Tree
 * https://leetcode.com/problems/second-minimum-node-in-a-binary-tree/
 *
 * submission : faster than 100%
 */
package main

func findSecondMinimumValue(node *TreeNode) int {
	min, secondMin := node.Val, -1
	traverse(node, min, &secondMin)
	return secondMin
}

func traverse(node *TreeNode, min int, res *int) {
	if node.Left == nil {
		return
	}
	big, small := node.Left, node.Right
	if big.Val > min && small.Val > min {
		return
	}
	if node.Left.Val < node.Right.Val {
		big, small = small, big
	}
	if (big.Val < *res || *res == -1) && big.Val != min {
		*res = big.Val
	}
	if big.Val == small.Val {
		traverse(node.Left, min ,res)
		traverse(node.Right, min ,res)
	} else {
		traverse(node.Left, min, res)
	}
}