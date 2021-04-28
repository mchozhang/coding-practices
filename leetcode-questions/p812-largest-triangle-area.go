/**
 * LeetCode : Largest Triangle Area
 * https://leetcode.com/problems/largest-triangle-area/
 *
 * submission : faster than 34%
 */
package main

func largestTriangleArea(points [][]int) float64 {
	max := 0
	for _, p1 := range points {
		for _, p2 := range points {
			for _, p3 := range points {
				area := p1[0]*p2[1] + p2[0]*p3[1] + p3[0]*p1[1] - p1[0]*p3[1] - p2[0]*p1[1] - p3[0]*p2[1]
				if area > max {
					max = area
				}
			}
		}
	}
	return float64(max) / 2
}
