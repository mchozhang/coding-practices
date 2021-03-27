/**
 * LeetCode : Unique Paths II
 * https://leetcode.com/problems/unique-paths-ii/
 *
 * submission : faster than 100%
 */
package main

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	row := len(obstacleGrid)
	col := len(obstacleGrid[0])
	dp := make([]int, col)
	dp[0] = 1
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if obstacleGrid[i][j] == 1 {
				dp[j] = 0
			} else if j > 0 {
				dp[j] += dp[j-1]
			}
		}
	}
	return dp[col-1]
}
