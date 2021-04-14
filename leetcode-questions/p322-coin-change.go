/**
 * LeetCode: Coin Change
 * https://leetcode.com/problems/coin-change/
 *
 * submission: faster than 88%
 */
package main

import (
	"math"
)

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)

	for i := 1; i <= amount; i++ {
		min := math.MaxInt16
		for _, coin := range coins {
			if coin > i {
				continue
			}
			if dp[i-coin]+1 < min {
				min = dp[i-coin] + 1
			}
		}
		dp[i] = min
	}
	if dp[amount] > amount {
		return -1
	}
	return dp[amount]
}
