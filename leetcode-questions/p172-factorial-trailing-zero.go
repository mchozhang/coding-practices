/**
 * LeetCode : Unique Binary Search Trees II
 * https://leetcode.com/problems/unique-binary-search-trees-ii/
 *
 * submission : faster than 91%
 */
package main

func trailingZeroes(n int) int {
	count := 0
	for i := 5; i <= n; i *= 5 {
		count += n / i
	}
	return count
}
