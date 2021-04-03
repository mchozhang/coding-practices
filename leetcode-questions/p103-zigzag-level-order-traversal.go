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
	levelQueue := []int{0}
	for len(queue) != 0 {
		current := queue[0]
		currentLevel := levelQueue[0]
		queue = queue[1:]
		levelQueue = levelQueue[1:]
		if current.Left != nil {
			queue = append(queue, current.Left)
			levelQueue = append(levelQueue, currentLevel+1)
		}

		if current.Right != nil {
			queue = append(queue, current.Right)
			levelQueue = append(levelQueue, currentLevel+1)
		}

		// append value to result
		if currentLevel < len(res) {
			res[currentLevel] = append(res[currentLevel], current.Val)
		} else {
			res = append(res, []int{current.Val})
		}
	}

	// reverse odd level result
	for i, arr := range res {
		if i % 2 == 1 {
			for j := 0; j < len(arr)/2;j++ {
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
