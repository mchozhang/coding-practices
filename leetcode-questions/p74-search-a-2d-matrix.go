/**
 * LeetCode: Search a 2D Matrix
 * https://leetcode.com/problems/search-a-2d-matrix/
 *
 * submission: faster than 100%
 */
package main

func searchMatrix(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])
	low, high := 0, m*n-1
	for low < high {
		mid := (low + high) / 2
		midX, midY := mid/n, mid%n
		if matrix[midX][midY] == target {
			return true
		}
		if matrix[midX][midY] < target {
			low = midX*n + midY + 1
		} else {
			high = midX*n + midY - 1
		}
	}
	return matrix[low/n][low%n] == target
}
