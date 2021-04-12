/**
 * LeetCode: House Robber
 * https://leetcode.com/problems/house-robber/
 *
 * submission: faster than 100%
 */
package main

func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 2 {
		if nums[0] > nums[1] {
			return nums[0]
		}
		return nums[1]
	}

	dp := make([]int, len(nums))
	dp[0] = nums[0]
	dp[1] = nums[1]
	dp[2] = nums[2] + nums[0]
	for i := 3; i < len(nums); i++ {
		dp[i] = dp[i-2] + nums[i]
		if dp[i-3] > dp[i-2] {
			dp[i] = dp[i-3] + nums[i]
		}
	}
	if dp[len(nums)-1] > dp[len(nums)-2] {
		return dp[len(nums)-1]
	}
	return dp[len(nums)-2]
}
