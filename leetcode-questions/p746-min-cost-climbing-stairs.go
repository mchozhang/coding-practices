/**
 * LeetCode : Min Cost Climbing Stairs
 * https://leetcode.com/problems/min-cost-climbing-stairs/
 *
 * submission : faster than 90%
 */
package main

func minCostClimbingStairs(cost []int) int {
	dp := make([]int, len(cost))
	dp[0] = cost[0]
	dp[1] = cost[1]

	for i := 2; i < len(cost); i++ {
		min := dp[i-2]
		if dp[i-1] < dp[i-2] {
			min = dp[i-1]
		}
		dp[i] = min + cost[i]
	}
	min := dp[len(cost) - 2]
	if min > dp[len(cost) - 1] {
		min = dp[len(cost) - 1]
	}
	return min
}
