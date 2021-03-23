/**
 * LeetCode : Rotate Image
 * https://leetcode.com/problems/rotate-image/
 *
 * submission : faster than 100%
 */
package main

func rotate(matrix [][]int) {
	n := len(matrix)
	for round := 0; round < n/2; round++ {
		for i := round; i < n-1-round; i++ {
			matrix[round][i], matrix[i][n-1-round], matrix[n-1-round][n-1-i], matrix[n-1-i][round] =
				matrix[n-1-i][round], matrix[round][i], matrix[i][n-1-round], matrix[n-1-round][n-1-i]

		}
	}
}
