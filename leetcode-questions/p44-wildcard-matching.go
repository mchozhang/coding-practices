/**
 * LeetCode : 44. Wildcard Matching
 * https://leetcode.com/problems/wildcard-matching/
 *
 * submission : faster than 84%
 */
package main

func isMatch(s string, p string) bool {
	dp := make([][]bool, len(p)+1)
	for i := 0; i <= len(p); i++ {
		dp[i] = make([]bool, len(s)+1)
	}
	dp[0][0] = true
	for i := 1; i <= len(s); i++ {
		dp[0][i] = false
	}
	for i := 1; i <= len(p); i++ {
		if p[i-1] == '*' {
			dp[i][0] = dp[i-1][0]
		}
	}

	for i := 0; i < len(p); i++ {
		for j := 0; j < len(s); j++ {
			if p[i] == s[j] || p[i] == '?' {
				dp[i+1][j+1] = dp[i][j]
			} else if p[i] == '*' {
				dp[i+1][j+1] = dp[i][j+1] || dp[i][j] || dp[i+1][j]
			}
		}
	}
	return dp[len(p)][len(s)]
}
