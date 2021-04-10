/**
 * LeetCode: Word Break
 * https://leetcode.com/problems/word-break/
 *
 * submission: faster than 100%
 */
package main

func wordBreak(s string, wordDict []string) bool {
	dp := make([]bool, len(s)+1)
	dp[0] = true

	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			if dp[j] && inDict(s[j:i], wordDict) {
				dp[i] = true
				break
			}
		}
	}
	return dp[len(s)]
}

func inDict(s string, dict []string) bool {
	for _, word := range dict {
		if s == word {
			return true
		}
	}
	return false
}