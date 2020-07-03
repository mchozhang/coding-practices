/**
 * 10. Regular Expression Matching
 * https://leetcode.com/problems/regular-expression-matching/
 * Given an input string (s) and a pattern (p), implement regular expression matching with support for '.' and '*'.
 * '.' Matches any single character.
 * '*' Matches zero or more of the preceding element.
 * Example:
 * Input: s = "aa", p = "a"
 * Output: false
 *
 * submission: faster than 93%
 */
package main

import "fmt"

func isMatch(s string, p string) bool {
	// 2D matrix, dp[i][j] represents whether the first i string matches the first j pattern
	dp := make([][]bool, len(s)+1)
	for i := 0; i < len(s)+1; i++ {
		dp[i] = make([]bool, len(p)+1)
	}

	// empty string matches empty pattern
	dp[0][0] = true

	// patterns that matches empty string
	for i := 1; i < len(p); i++ {
		if p[i] == '*' && dp[0][i-1] {
			dp[0][i+1] = true
		}
	}

	for i := 0; i < len(s); i++ {
		for j := 0; j < len(p); j++ {
			if s[i] == p[j] || p[j] == '.' {
				dp[i+1][j+1] = dp[i][j]
			}
			if p[j] == '*' {
				if s[i] != p[j-1] && p[j-1] != '.' {
					// 'a' repeats 0 time in 'a*'
					dp[i+1][j+1] = dp[i+1][j-1]
				} else {
					// 'a' repeats 0 time, 1 time and multiple times respectively in 'a*'
					dp[i+1][j+1] = dp[i+1][j-1] || dp[i+1][j] || dp[i][j+1]
				}
			}

		}
	}
	return dp[len(s)][len(p)]
}

func main() {
	fmt.Println(isMatch("aaa", "a*"))
	fmt.Println(isMatch("abc", "abc"))
	fmt.Println(isMatch("abb", "ab*"))
	fmt.Println(isMatch("1234567", ".*"))
	fmt.Println(isMatch("abcd", "a.c."))
	fmt.Println(isMatch("mississippi", "mis*is*p*."))
}
