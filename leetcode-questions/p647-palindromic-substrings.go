/**
 * LeetCode : Palindromic Substrings
 * https://leetcode.com/problems/palindromic-substrings/
 *
 * submission : faster than 100%
 */
package main

func countSubstrings(s string) int {
	if len(s) == 0 {
		return 0
	}
	sum := 0
	for i := 0; i < len(s); i++ {
		sum += extentPalindromicString(s, i, i)
		sum += extentPalindromicString(s, i, i+1)
	}
	return sum
}

func extentPalindromicString(s string, left int, right int) int {
	count := 0
	for left > -1 && right < len(s) {
		if s[left] == s[right] {
			count++
			left--
			right++
		} else {
			break
		}
	}
	return count
}
