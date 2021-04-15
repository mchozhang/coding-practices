/**
 * LeetCode : Daily Temperatures
 * https://leetcode.com/problems/daily-temperatures/
 *
 * submission : faster than 91%
 */
package main

func dailyTemperatures(t []int) []int {
	res := make([]int, len(t))
	var stack []int

	for i, num := range t {
		for len(stack) != 0 {
			top := stack[len(stack)-1]
			if num <= t[top] {
				break
			}
			res[top] = i - top
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}
	return res
}
