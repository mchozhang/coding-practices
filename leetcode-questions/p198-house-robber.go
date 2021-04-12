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

	pre3, pre2, pre1 := nums[0], nums[1], nums[0] + nums[2]
	for i := 3; i < len(nums); i++ {
		value := pre2 + nums[i]
		if pre3 > pre2 {
			value = pre3 + nums[i]
		}
		pre3, pre2, pre1 = pre2, pre1, value
	}
	if pre1 > pre2 {
		return pre1
	}
	return pre2
}
