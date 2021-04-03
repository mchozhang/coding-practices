/**
 * LeetCode : Binary Tree Paths
 * https://leetcode.com/problems/binary-tree-paths/
 *
 * submission : faster than 100%
 */
package main

import (
	"strconv"
	"strings"
)

func binaryTreePaths(root *TreeNode) []string {
	res := make([]string, 0)
	if root != nil {
		traverse(root, []string{}, &res)
	}
	return res
}

func traverse(root *TreeNode, path []string, res *[]string) {
	newPath := append([]string{}, path...)
	newPath = append(newPath, strconv.Itoa(root.Val))
	if root.Left == nil && root.Right == nil {
		*res = append(*res, strings.Join(newPath, "->") )
		return
	}
	if root.Left != nil {
		traverse(root.Left, newPath, res)
	}
	if root.Right != nil {
		traverse(root.Right, newPath, res)
	}
}