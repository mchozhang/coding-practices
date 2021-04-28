/**
 * LeetCode : Set Matrix Zeroes
 * https://leetcode.com/problems/set-matrix-zeroes/
 *
 * submission : faster than 88%
 */
package main

/**
 * O(m+n)
 */
func setZeroes(matrix [][]int) {
	var rows, cols = map[int]int{}, map[int]int{}
	for i, row := range matrix {
		for j, cell := range row {
			if cell == 0 {
				rows[i] = 0
				cols[j] = 0
			}
		}
	}
	for row := range rows {
		matrix[row] = make([]int, len(matrix[0]))
	}
	for col := range cols {
		for i := 0; i < len(matrix);i++ {
			matrix[i][col] = 0
		}
	}
}

/**
 * O(1) space complexity
 */
func setZeroes2(matrix [][]int) {
	firstRow, firstCol := false, false
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == 0 {
				if i == 0 {
					firstRow = true
				}
				if j == 0 {
					firstCol = true
				}
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}
	// set cells except the first row and column
	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[0]); j++ {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
	}
	// set the first row and column
	if firstRow {
		matrix[0] = make([]int, len(matrix[0]))
	}
	if firstCol {
		for i := 0; i < len(matrix); i++ {
			matrix[i][0] = 0
		}
	}
}
