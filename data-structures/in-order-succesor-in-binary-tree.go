/**
 * Given a node in a binary search tree, where left < parent < right,
 * find the in-order successor of that node.
 * Note:
 * an in-order traversal will result in the visiting nodes in sorted, ascending order.
 *        5
 *      /    \
 *    3       7
 *   / \     /  \
 *  2   4   6    8
 *  Given 5 -> 6
 *  Given 6 -> 7
 *  Given 8 -> null
 */
package main

import "fmt"

type Node struct {
	Val int
	Left *Node
	Right *Node
	Parent *Node
}

func inOrderSuccessor(node *Node) *Node {

	if node.Right != nil {
		// search right sub-tree if right child exists
		return findSmallest(node.Right)
	} else if node.Parent.Val > node.Val {
		// return parent if parent is greater than node value
		return node.Parent
	} else {
		return nil
	}
}

/**
 * find successor in right sub-tree
 */
func findSmallest(node *Node) *Node {
	if node.Left != nil {
		return findSmallest(node.Left)
	} else {
		return node
	}
}

func main() {
	var node1 = &Node{5, nil, nil,nil}
	var node2 = &Node{3, nil, nil, node1}
	var node3 = &Node{2, nil, nil, node2}
	var node4 = &Node{4, nil, nil, node2}
	var node5 = &Node{7, nil, nil, node1}
	var node6 = &Node{6, nil, nil, node5}
	var node7 = &Node{8, nil, nil, node5}
	node1.Left = node2
	node1.Right = node5
	node2.Left = node3
	node2.Right = node4
	node5.Left = node6
	node5.Right = node7

	// given node1, print 6
	fmt.Println(inOrderSuccessor(node1).Val)

	// given node6, print 7
	fmt.Println(inOrderSuccessor(node6).Val)

	// given node2, print 4
	fmt.Println(inOrderSuccessor(node2).Val)

	// given node7, print nil
	fmt.Println(inOrderSuccessor(node7))
}

