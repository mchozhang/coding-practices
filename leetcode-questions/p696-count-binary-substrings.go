/**
 * LeetCode : Count Binary Substrings
 * https://leetcode.com/problems/count-binary-substrings/
 *
 * submission : faster than 88%
 */
package main

func countBinarySubstrings(s string) int {
	pre, cur := 0, 1
	sum := 0
	for i := 1; i < len(s); i++ {
		if s[i] != s[i-1] {
			pre = cur
			cur = 1
		} else {
			cur++
		}
		if cur <= pre {
			sum++
		}
	}
	return sum
}
