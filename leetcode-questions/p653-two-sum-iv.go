/**
 * LeetCode : Two Sum IV - Input is a BST
 * https://leetcode.com/problems/two-sum-iv-input-is-a-bst/
 *
 * submission : faster than 88%
 */
package main

/**
 * O(n) time complexity, O(logN) space complexity
 */
func findTarget(root *TreeNode, k int) bool {
	left, right := root, root
	leftVal, rightVal := 0, 0
	var leftStack, rightStack []*TreeNode
	for left != nil {
		leftStack = append(leftStack, left)
		leftVal = left.Val
		left = left.Left
	}
	for right != nil {
		rightStack = append(rightStack, right)
		rightVal = right.Val
		right = right.Right
	}

	for leftVal < rightVal {
		sum := leftVal + rightVal
		if sum == k {
			return true
		}
		if sum < k {
			for left != nil {
				leftStack = append(leftStack, left)
				leftVal = left.Val
				left = left.Left
			}
			leftVal = leftStack[len(leftStack)-1].Val
			left = leftStack[len(leftStack)-1].Right
			leftStack = leftStack[:len(leftStack)-1]
		} else {
			for right != nil {
				rightStack = append(rightStack, right)
				rightVal = right.Val
				right = right.Right
			}
			rightVal = rightStack[len(rightStack)-1].Val
			right = rightStack[len(rightStack)-1].Left
			rightStack = rightStack[:len(rightStack)-1]
		}
	}
	return false
}
