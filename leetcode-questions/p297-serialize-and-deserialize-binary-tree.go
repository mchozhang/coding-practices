/**
 * LeetCode :  Serialize and Deserialize Binary Tree
 * https://leetcode.com/problems/serialize-and-deserialize-binary-tree/
 *
 * submission : faster than 82%
 */
package main

import (
	"strconv"
	"strings"
)

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	var nodeStrings []string
	queue := []*TreeNode{root}
	for len(queue) != 0 {
		var temp []*TreeNode
		for _, node := range queue {
			if node == nil {
				nodeStrings = append(nodeStrings, "")
			} else {
				nodeStrings = append(nodeStrings, strconv.Itoa(node.Val))
				temp = append(temp, node.Left)
				temp = append(temp, node.Right)
			}
		}
		queue = temp
	}

	return strings.Join(nodeStrings, ",")
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	if data == "" {
		return nil
	}
	nodeStrings := strings.Split(data, ",")
	nodes := make([]*TreeNode, len(nodeStrings))
	for i, s := range nodeStrings {
		value, _ := strconv.Atoi(s)
		nodes[i] = &TreeNode{Val: value}
	}

	queue := []int{0}
	nextPos := 1
	for len(queue) != 0 {
		pos := queue[0]
		queue = queue[1:]
		node := nodes[pos]
		if nodeStrings[nextPos] != "" {
			node.Left = nodes[nextPos]
			queue = append(queue, nextPos)
		}
		if nodeStrings[nextPos+1] != "" {
			node.Right = nodes[nextPos+1]
			queue = append(queue, nextPos + 1)
		}
		nextPos += 2
	}
	return nodes[0]
}

