/**
 * LeetCode : Kth Smallest Element in a BST
 * https://leetcode.com/problems/kth-smallest-element-in-a-bst/
 * Given a binary search tree, write a function kthSmallest to find the kth smallest element in it.
 *
 * Example:
 * root = [3,1,4,null,2], k = 1
 *   3
 *  / \
 * 1   4
 *  \
 *   2
 *
 * output: 1
 *
 * submission : faster than 95%
 */

package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var count = 0
var result = 0

func kthSmallest(root *TreeNode, k int) int {
    count = 0
    result = 0
	inOrderSearch(root, k)
    return result
}

func inOrderSearch(node *TreeNode, k int) {
	if node.Left != nil {
		inOrderSearch(node.Left, k)
	}
	count++
	if count == k {
		result = node.Val
		return
	}

	if node.Right != nil {
		inOrderSearch(node.Right, k)
	}
}

func main() {
	node1 := &TreeNode{Val: 5}
	node2 := &TreeNode{Val: 3}
	node3 := &TreeNode{Val: 6}
	node4 := &TreeNode{Val: 2}
	node5 := &TreeNode{Val: 4}
	node6 := &TreeNode{Val: 1}

	node1.Left = node2
	node1.Right = node3
	node2.Left = node4
	node2.Right = node5
	node4.Left = node6
	fmt.Println(kthSmallest(node1, 5))
}



