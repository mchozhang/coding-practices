/**
 * LeetCode: Find the Duplicate Number
 * https://leetcode.com/problems/find-the-duplicate-number/
 *
 * submission: faster than 99%
 */
package main

func findDuplicate(nums []int) int {
	for i := 0; i < len(nums); {
		if nums[i] != i + 1{
			if nums[nums[i]] == nums[i] {
				return nums[i]
			}
			nums[i], nums[nums[i]] = nums[nums[i]], nums[i]
		} else {
			i++
		}
	}
	return nums[len(nums) - 1]
}
