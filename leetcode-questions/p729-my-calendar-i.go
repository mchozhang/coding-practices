/**
 * LeetCode: 729. My Calendar I
 * https://leetcode.com/problems/my-calendar-i/
 *
 * submission: faster than 88%
 */
package main

type Node struct {
	Start int
	End   int
	Left  *Node
	Right *Node
}

type MyCalendar struct {
	Root *Node
}

func Constructor() MyCalendar {
	return MyCalendar{}
}

func insert(node *Node, start int, end int) bool {
	if start > node.End {
		if node.Right != nil {
			return insert(node.Right, start, end)
		} else {
			node.Right = &Node{Start: start, End: end}
			return true
		}
	} else if end <= node.Start {
		if node.Left != nil {
			return insert(node.Left, start, end)
		} else {
			node.Left = &Node{Start: start, End: end}
			return false
		}
	} else {
		return false
	}
}

func (this *MyCalendar) Book(start int, end int) bool {
	if this.Root == nil {
		this.Root = &Node{Start: start, End: end}
		return true
	}
	return insert(this.Root, start, end)
}
