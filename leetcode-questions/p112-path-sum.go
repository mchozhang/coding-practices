/**
 * LeetCode : Path Sum
 * https://leetcode.com/problems/path-sum/
 * Given a binary tree and a sum, determine if the tree has a root-to-leaf path such that adding up all the values along
the path equals the given sum.
 *
 * Example:
 * Input:
 *       5
 *      / \
 *     4   8
 *    /   / \
 *   11  13  4
 *  /  \      \
 * 7    2      1
 * Output: 5->4->11->2
 *
 * submission : faster than 96.34%
 */

package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}


func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}

	if root.Left == nil && root.Right == nil && root.Val == sum {
		return true
	}

	left := hasPathSum(root.Left, sum - root.Val)
	right := hasPathSum(root.Right, sum - root.Val)

	return left != false || right != false
}

func main() {
	node14 := &TreeNode{1, nil, nil}
	node7 := &TreeNode{7, nil, nil}
	node8 := &TreeNode{2, nil, nil}

	node3 := &TreeNode{11, node7, node8}
	node5 := &TreeNode{13, nil, nil}
	node6 := &TreeNode{4, nil, node14}

	node1 := &TreeNode{4, node3, nil}
	node2 := &TreeNode{8, node5, node6}
	node0 := &TreeNode{5, node1, node2}

	fmt.Println(hasPathSum(node0, 20))
}
