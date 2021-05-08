/**
 * LeetCode: 583. Delete Operation for Two Strings
 * https://leetcode.com/problems/delete-operation-for-two-strings/
 *
 * submission : faster than 100%
 */
package main

func minDistance(word1 string, word2 string) int {
	dp := make([][]int, len(word1)+1)
	for i := 0; i <= len(word1); i++ {
		dp[i] = make([]int, len(word2)+1)
	}
	for i := 1; i <= len(word1); i++ {
		dp[i][0] = i
	}
	for i := 1; i <= len(word2); i++ {
		dp[0][i] = i
	}
	for i := 0; i < len(word1); i++ {
		for j := 0; j < len(word2); j++ {
			if word1[i] == word2[j] {
				dp[i+1][j+1] = dp[i][j]
			} else {
				min := dp[i][j+1]
				if min > dp[i+1][j] {
					min = dp[i+1][j]
				}
				dp[i+1][j+1] = min + 1
			}
		}
	}
	return dp[len(word1)][len(word2)]
}
