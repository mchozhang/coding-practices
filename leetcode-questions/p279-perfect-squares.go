/**
 * LeetCode: Perfect Squares
 * https://leetcode.com/problems/perfect-squares/
 *
 * submission: faster than 88%
 */
package main

func numSquares(n int) int {
	dp := make([]int, n+1)

	for i := 1; i <= n; i++ {
		if i*i <= n {
			dp[i*i] = 1
		}

		if dp[i] != 1 {
			min := dp[1] + dp[i-1]
			for j := 1; j*j <= i; j++ {
				if dp[j*j]+dp[n-j*j] < min {
					min = dp[j*j] + dp[i-j*j]
				}
			}
			dp[i] = min
		}
	}
	return dp[n]
}
