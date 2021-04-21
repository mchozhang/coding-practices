/**
 * LeetCode : Valid Perfect Square
 * https://leetcode.com/problems/valid-perfect-square/
 *
 * submission : faster than 88%
 */
package main

func isPerfectSquare(num int) bool {
	for i := 1; i*i <= num;i++ {
		if i*i == num {
			return true
		}
	}
	return false
}