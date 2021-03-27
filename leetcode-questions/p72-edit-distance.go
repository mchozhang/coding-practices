/**
 * LeetCode : Edit Distance
 * https://leetcode.com/problems/edit-distance/
 *
 * submission : faster than 100%
 */
package main

func minDistance(word1 string, word2 string) int {
	dp := make([][]int, len(word1)+1)
	for i := 0; i <= len(word1); i++ {
		dp[i] = make([]int, len(word2)+1)
	}

	// first row
	for i := 0; i <= len(word2); i++ {
		dp[0][i] = i
	}

	// first column
	for i := 1; i <= len(word1); i++ {
		dp[i][0] = i
	}
	for i := 1; i <= len(word1); i++ {
		for j := 1; j <= len(word2); j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				min := dp[i-1][j-1]
				if dp[i-1][j] < min {
					min = dp[i-1][j]
				}
				if dp[i][j-1] < min {
					min = dp[i][j-1]
				}
				dp[i][j] = min + 1
			}
		}
	}
	return dp[len(word1)][len(word2)]
}
