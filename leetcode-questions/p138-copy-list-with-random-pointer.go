/**
 * LeetCode: Copy List with Random Pointer
 * https://leetcode.com/problems/copy-list-with-random-pointer/
 *
 * submission: faster than 100%
 */
package main

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}

	mapper := map[*Node]*Node{head: {Val: head.Val}}
	preClone := &Node{}
	for current := head; current != nil; current = current.Next {
		var clone *Node
		if currentClone, isIn := mapper[current]; isIn {
			clone = currentClone
		} else {
			clone = &Node{Val: current.Val}
			mapper[current] = clone
		}
		if current.Random != nil {
			if currentRandom, isIn := mapper[current.Random]; isIn {
				clone.Random = currentRandom
			} else {
				cloneRandom := &Node{Val: current.Random.Val}
				mapper[current.Random] = cloneRandom
				clone.Random = cloneRandom
			}
		}

		preClone.Next = clone
		preClone = clone
	}
	return mapper[head]
}
