/**
 * 437. Path Sum III
 * https://leetcode.com/problems/path-sum-iii/
 *
 * submission: faster than 44%
 */
package main

func pathSum(root *TreeNode, sum int) int {
	count := 0
	traverse(root, 0, sum, false, &count)
	return count
}

func traverse(node *TreeNode, sum int, target int, chosen bool, count *int) {
	if node == nil {
		return
	}
	if sum + node.Val == target {
		*count++
	}

	if chosen == false {
		traverse(node.Left, sum, target, false, count)
		traverse(node.Right, sum, target, false, count)
	}
	traverse(node.Left, sum + node.Val, target, true, count)
	traverse(node.Right, sum + node.Val, target, true, count)
}
