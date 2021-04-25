/**
 * LeetCode : Task Scheduler
 * https://leetcode.com/problems/counting-bits/
 *
 * submission : faster than 100%
 */
package main

func fib(n int) int {
	if n < 2 {
		return n
	}
	pre, cur := 0,1
	for i := 2; i <= n; i++ {
		temp := cur
		cur += pre
		pre = temp
	}
	return cur
}
