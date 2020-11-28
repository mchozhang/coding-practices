/**
 * LeetCode : Unique Paths
 * https://leetcode.com/problems/unique-paths/
 * A robot is located at the top-left corner of a m x n grid (marked 'Start' in the diagram below).
 * The robot can only move either down or right at any point in time. The robot is trying to reach the bottom-right
 * corner of the grid
 *
 * Example:
 * Input: m = 3, n = 7
 * Output: 28
 *
 * submission : faster than 100%
 */


package main

import (
	"fmt"
)

func uniquePaths(m int, n int) int {
	if m == 1 || n == 1 {
		return 1
	}
	dp := make([]int, n)
	for i := range dp {
		dp[i] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[j] += dp[j-1]
		}
	}
	return dp[n - 1]
}

func main() {
	fmt.Println(uniquePaths(3, 3))

	fmt.Println(uniquePaths(3, 7))
	fmt.Println(uniquePaths(51, 9))

}