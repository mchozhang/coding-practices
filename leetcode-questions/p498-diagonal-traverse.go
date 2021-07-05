/**
 * LeetCode: 498. Diagonal Traverse
 * https://leetcode.com/problems/diagonal-traverse/
 *
 * submission: faster than 99%
 */
package main

func findDiagonalOrder(mat [][]int) []int {
	var res []int
	up := true
	x, y := 0, 0
	for x != len(mat) && y != len(mat[0]) {
		res = append(res, mat[x][y])
		if up {
			if y == len(mat[0])-1 {
				x++ // go down
				up = false
			} else if x == 0 {
				y++ // go right
				up = false
			} else {
				y++
				x--
			}
		} else {
			if x == len(mat)-1 {
				y++ // go right
				up = true
			} else if y == 0 {
				x++ // go down
				up = true
			} else {
				x++
				y--
			}
		}
	}
	return res
}
