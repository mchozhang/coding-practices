/**
 * LeetCode : Is Subsequence
 * https://leetcode.com/problems/is-subsequence/
 *
 * submission : faster than 100%
 */
package main

func isSubsequence(s string, t string) bool {
	j := 0
	for i := 0; i < len(t) && j < len(s); i++ {
		if s[j] == t[i] {
			j++
		}
	}
	return j == len(s)
}
