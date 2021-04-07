/**
 * 997. Find the Town Judge
 * https://leetcode.com/problems/find-the-town-judge/
 *
 * submission: faster than 100%
 */
package main

type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}
	built := map[int]*Node{node.Val: {Val: node.Val}}

	queue := []*Node{node}

	for len(queue) != 0 {
		current := queue[0]
		clone := built[current.Val]
		queue = queue[1:]

		for _, neighbor := range current.Neighbors {
			var cloneNeighbor *Node
			if node, isIn:= built[neighbor.Val];isIn {
				cloneNeighbor = node
			} else {
				cloneNeighbor = &Node{Val:neighbor.Val}
				built[neighbor.Val] = cloneNeighbor
				queue = append(queue, neighbor)
			}
			clone.Neighbors = append(clone.Neighbors, cloneNeighbor)
		}
	}
	return built[node.Val]
}
