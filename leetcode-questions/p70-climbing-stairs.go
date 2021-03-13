/**
 * LeetCode : Climbing Stairs
 * https://leetcode.com/problems/climbing-stairs/
 * You are climbing a staircase. It takes n steps to reach the top.
 * Each time you can either climb 1 or 2 steps. In how many distinct ways can you climb to the top?
 * submission : faster than 100%
 */
package main

func climbStairs(n int) int {
	a := 1
	b := 1
	for i := 1; i < n; i++ {
		a, b = b, a+b
	}
	return a
}
