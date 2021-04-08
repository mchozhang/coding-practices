/**
 * 199. Binary Tree Right Side View
 * https://leetcode.com/problems/binary-tree-right-side-view/
 *
 * submission: faster than 100%
 */
package main

func rightSideView(root *TreeNode) []int {
	var ans []int
	if root != nil {
		traverse(0, root, &ans)
	}
	return ans
}

func traverse(layer int, node *TreeNode, ans *[]int) {
	if len(*ans) <= layer {
		*ans = append(*ans, node.Val)
	} else {
		(*ans)[layer] = node.Val
	}
	if node.Left != nil {
		traverse(layer+1, node.Left, ans)
	}

	if node.Right != nil {
		traverse(layer+1, node.Right, ans)
	}
}

