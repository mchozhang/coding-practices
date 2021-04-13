/**
 * LeetCode: Search a 2D Matrix II
 * https://leetcode.com/problems/search-a-2d-matrix-ii/
 *
 * submission: faster than 100%
 */
package main

import "fmt"

func searchMatrix(matrix [][]int, target int) bool {
	return search(matrix, target, 0, len(matrix)-1, 0, len(matrix[0])-1)
}

func search(matrix [][]int, target int, lowX int, highX int, lowY int, highY int) bool {
	if lowX > highX || lowY > highY {
		return false
	}
	fmt.Println(lowX, highX, lowY, highY)
	midX := (lowX + highX) / 2
	midY := (lowY + highY) / 2
	if matrix[midX][midY] == target {
		return true
	} else if lowX == highX && lowY == highY {
		return false
	} else if matrix[midX][midY] > target {
		return search(matrix, target, lowX, highX, lowY, midY-1) ||
			search(matrix, target, lowX, midX-1, midY, highY)
	} else {
		return search(matrix, target, midX+1, highX, lowY, midY) ||
			search(matrix, target, lowX, highX, midY+1, highY)
	}
}

/**
 * O(m+n)
 */
func searchMatrix2(matrix [][]int, target int) bool {
	i, j := 0, len(matrix[0])-1
	for i < len(matrix) && j > -1 {
		if matrix[i][j] == target {
			return true
		}
		if matrix[i][j] > target {
			j--
		} else {
			i++
		}
	}
	return false
}
