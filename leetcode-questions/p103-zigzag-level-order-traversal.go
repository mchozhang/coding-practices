/**
 * LeetCode : Binary Tree Zigzag Level Order Traversal
 * https://leetcode.com/problems/binary-tree-zigzag-level-order-traversal/
 *
 * submission : faster than 100%
 */
package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
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
		res = append(res, level)
	}

	// reverse odd level result
	for i, arr := range res {
		if i%2 == 1 {
			for j := 0; j < len(arr)/2; j++ {
				arr[j], arr[len(arr)-1-j] = arr[len(arr)-1-j], arr[j]
			}
		}
	}
	return res
}

func main() {
	node3 := &TreeNode{Val: 20}
	node2 := &TreeNode{Val: 9}
	node1 := &TreeNode{Val: 3, Left: node2, Right: node3}
	fmt.Println(zigzagLevelOrder(node1))
}
