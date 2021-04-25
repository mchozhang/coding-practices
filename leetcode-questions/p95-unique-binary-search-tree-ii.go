/**
 * LeetCode : Unique Binary Search Trees II
 * https://leetcode.com/problems/unique-binary-search-trees-ii/
 *
 * submission : faster than 100%
 */
package main

func generateTrees(n int) []*TreeNode {
	return buildBinarySearchTree(1, n)
}

func buildBinarySearchTree(low int, high int) []*TreeNode {
	var res []*TreeNode
	if low > high {
		res = append(res, nil)
	}
	for i := low; i <= high; i++ {
		leftList := buildBinarySearchTree(low, i-1)
		rightList := buildBinarySearchTree(i+1, high)
		for _, leftNode := range leftList {
			for _, rightNode := range rightList {
				node := &TreeNode{Val: i}
				node.Left = leftNode
				node.Right = rightNode
				res = append(res, node)
			}
		}
	}
	return res
}
