/**
 * LeetCode : N Queens II
 * https://leetcode.com/problems/n-queens-ii/
 *
 * submission : faster than 100%
 */
package main

func totalNQueens(n int) int {
	count := 0
	var grid [][]byte
	for i := 0; i < n; i++ {
		row := make([]byte, n)
		grid = append(grid, row)
	}
	dfs(grid, 0, &count)
	return count
}

func dfs(grid [][]byte, i int, count *int) {
	// place a queen to a row
	for j := 0; j < len(grid); j++ {
		if isValid(grid, i, j) {
			grid[i][j] = 1
			if i == len(grid)-1 {
				*count++
			} else {
				dfs(grid, i+1, count)
			}
			grid[i][j] = 0
		}
	}
}

func isValid(grid [][]byte, x int, y int) bool {
	// check column
	for i := x; i > -1; i-- {
		if grid[i][y] == 1 {
			return false
		}
	}
	// check diagonal
	for i, j := x, y; j < len(grid) && i > -1; {
		if grid[i][j] == 1 {
			return false
		}
		j++
		i--
	}
	for i, j := x, y; j > -1 && i > -1; {
		if grid[i][j] == 1 {
			return false
		}
		j--
		i--
	}
	return true
}
