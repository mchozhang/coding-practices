/**
 * LeetCode : Spiral Matrix II
 * https://leetcode.com/problems/spiral-matrix-ii/
 *
 * submission : faster than 100%
 */
package main

func generateMatrix(n int) [][]int {
	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
	}
	max := n*n + 1
	x := 0
	y := 0
	pos := 0
	directions := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	for value := 1; value < max; value++ {
		matrix[x][y] = value
		nextX := x + directions[pos][0]
		nextY := y + directions[pos][1]
		if -1 < nextX && nextX < n && -1 < nextY && nextY < n && matrix[nextX][nextY] == 0 {
			x = nextX
			y = nextY
		} else {
			pos++
			pos %= 4
			x += directions[pos][0]
			y += directions[pos][1]
		}
	}
	return matrix
}
