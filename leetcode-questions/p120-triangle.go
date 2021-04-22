/**
 * LeetCode : Triangle
 * https://leetcode.com/problems/triangle/
 *
 * submission : faster than 95%
 */
package main

func minimumTotal(triangle [][]int) int {
	dp := make([]int, len(triangle))
	dp[0] = triangle[0][0]

	for i := 1; i < len(triangle); i++ {
		j := len(triangle[i]) - 1
		dp[j] += dp[j-1] + triangle[i][j]
		for j--; j > 0; j-- {
			min := dp[j]
			if dp[j-1] < min {
				min = dp[j-1]
			}
			dp[j] = min + triangle[i][j]
		}
		dp[0] += triangle[i][0]
	}

	min := dp[0]
	for i := 0; i < len(dp); i++ {
		if dp[i] < min {
			min = dp[i]
		}
	}
	return min
}
