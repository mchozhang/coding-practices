/**
 * LeetCode: House Robber II
 * https://leetcode.com/problems/house-robber-ii/
 *
 * submission: faster than 100%
 */
package main

func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	max := robRange(nums, 0, len(nums)-1)
	skipFirst := robRange(nums, 1, len(nums))
	if max < skipFirst {
		return skipFirst
	}
	return max
}

func robRange(nums []int, low int, high int) int {
	robPre := 0
	skipPre := 0
	for i := low; i < high; i++ {
		temp := skipPre
		if skipPre < robPre {
			skipPre = robPre
		}
		robPre = temp + nums[i]
	}
	if skipPre > robPre {
		return skipPre
	}
	return robPre
}
