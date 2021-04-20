/**
 * LeetCode : N Queens
 * https://leetcode.com/problems/n-queens/
 *
 * submission : faster than 100%
 */
package main

func solveNQueens(n int) [][]string {
	var res [][]string
	var grid [][]byte
	for i := 0; i < n; i++ {
		row := make([]byte, n)
		for j := 0; j < n; j++ {
			row[j] = '.'
		}
		grid = append(grid, row)
	}
	dfs(grid, 0, &res)
	return res
}

func dfs(grid [][]byte, i int, res *[][]string) {
	// place a queen to a row
	for j := 0; j < len(grid); j++ {
		if isValid(grid, i, j) {
			grid[i][j] = 'Q'
			if i == len(grid)-1 {
				ans := make([]string,len(grid))
				for r := range ans {
					ans[r] = string(grid[r])
				}
				*res = append(*res, ans)
			} else {
				dfs(grid, i + 1, res)
			}
			grid[i][j] = '.'
		}
	}
}

func isValid(grid [][]byte, x int, y int) bool {
	// check column
	for i := x; i > -1; i-- {
		if grid[i][y] != '.' {
			return false
		}
	}
	// check diagonal
	for i, j := x, y; j < len(grid) && i > -1; {
		if grid[i][j] != '.' {
			return false
		}
		j++
		i--
	}
	for i, j := x, y; j > -1 && i > -1; {
		if grid[i][j] != '.' {
			return false
		}
		j--
		i--
	}
	return true
}
