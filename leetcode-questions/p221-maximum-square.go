/**
 * LeetCode: Maximum Square
 * https://leetcode.com/problems/maximal-square/
 *
 * submission: faster than 100%
 */
package main

func maximalSquare(matrix [][]byte) int {
	max := 0
	m := len(matrix)
	n := len(matrix[0])
	dp := make([][]int, m+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, n+1)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == byte('1') {
				// find the min value of its adjacent cell
				min := dp[i][j]
				if dp[i+1][j] < min {
					min = dp[i+1][j]
				}
				if dp[i][j+1] < min {
					min = dp[i][j+1]
				}
				dp[i+1][j+1] = min + 1

				if dp[i+1][j+1] > max {
					max = dp[i+1][j+1]
				}
			} else {
				dp[i+1][j+1] = 0
			}
		}
	}
	return max * max
}
