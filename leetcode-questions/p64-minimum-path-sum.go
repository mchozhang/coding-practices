/**
 * LeetCode : Minimum Path Sum
 * https://leetcode.com/problems/minimum-path-sum/
 *
 * submission : faster than 88%
 */
package main

func minPathSum(grid [][]int) int {
	m := len(grid) // row number
	n := len(grid[0]) // column
	for i:= 1; i < n; i++ {
		grid[0][i] += grid[0][i-1]
	}
	for i := 1;i < m; i++ {
		grid[i][0] += grid[i-1][0]
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			min := grid[i-1][j]
			if grid[i][j-1] < min {
				min = grid[i][j-1]
			}
			grid[i][j] += min
		}
	}
	return grid[m-1][n-1]
}

