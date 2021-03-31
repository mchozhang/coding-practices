/**
 * LeetCode : Divisor Game
 * https://leetcode.com/problems/divisor-game/
 *
 * submission : faster than 40%
 */
package main

func divisorGame(n int) bool {
	dp := make([]bool, n + 1)
	dp[0] = true
	dp[1] = false
	for i:= 2; i <= n;i++ {
		for j := 1; j <= i/2;j++ {
			if i % j == 0 {
				if dp[i-j] == false {
					dp[i] = true
					break
				}
			}
		}
	}
	return dp[n]
}