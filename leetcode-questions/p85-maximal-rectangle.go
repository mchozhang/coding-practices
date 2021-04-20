/**
 * LeetCode : Maximal Rectangle
 * https://leetcode.com/problems/maximal-rectangle/
 *
 * submission : faster than 100%
 */
package main

func maximalRectangle(matrix [][]byte) int {
	if len(matrix) == 0 {
		return 0
	}

	heights := make([][]int, len(matrix)+1)
	max := 0
	for i := 0; i < len(matrix)+1; i++ {
		heights[i] = make([]int, len(matrix[0])+1)
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == '1' {
				heights[i+1][j] = heights[i][j] + 1
			}
		}

		var stack []int
		for j, value := range heights[i+1] {
			for len(stack) != 0 && value < heights[i+1][stack[len(stack)-1]] {
				height := heights[i+1][stack[len(stack)-1]]
				stack = stack[:len(stack)-1]
				width := j
				if len(stack) != 0 {
					width = j - stack[len(stack)-1] - 1
				}
				area := height * width
				if max < area {
					max = area
				}
			}
			stack = append(stack, j)
		}
	}
	return max
}
