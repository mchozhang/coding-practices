/**
 * LeetCode : Rotting Oranges
 * https://leetcode.com/problems/rotting-oranges/
 *
 * submission : faster than 75%
 */
package main

func orangesRotting(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	count := 0
	var queue [][]int
	for i, row := range grid {
		for j, cell := range row {
			if cell == 2 {
				queue = append(queue, []int{i, j})
			} else if cell == 1 {
				count++
			}
		}
	}
	if count == 0 {
		return 0
	}
	dirs := [4][2]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
	time := 0
	for len(queue) != 0 {
		time++
		var nextQueue [][]int
		for _, pos := range queue {
			for _, dir :=range dirs {
				nextX, nextY := pos[0] + dir[0], pos[1] + dir[1]
				if -1 < nextX && nextX < m && -1 < nextY && nextY < n && grid[nextX][nextY] == 1 {
					nextQueue = append(nextQueue, []int{nextX, nextY})
					grid[nextX][nextY] = 2
					count--
				}
			}
		}
		queue = nextQueue
	}
	if count == 0 {
		return time - 1
	}
	return -1
}
