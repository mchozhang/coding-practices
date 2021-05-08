/**
 * LeetCode : Largest Rectangle in Histogram
 * https://leetcode.com/problems/largest-rectangle-in-histogram/
 *
 * submission : faster than 88%
 */
package main

func largestRectangleArea(heights []int) int {
	heights = append(heights, 0)
	var stack = []int{}
	max := 0
	for i, value := range heights {
		for len(stack) != 0 && value < heights[stack[len(stack)-1]] {
			height := heights[stack[len(stack)-1]]
			stack = stack[:len(stack)-1]
			width := i
			if len(stack) != 0 {
				width = i - stack[len(stack)-1] - 1
			}

			if max < height*width {
				max = height * width
			}
		}
		stack = append(stack, i)
	}
	return max
}
