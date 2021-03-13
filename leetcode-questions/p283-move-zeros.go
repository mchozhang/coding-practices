/**
 * LeetCode : Move Zeroes
 * https://leetcode.com/problems/move-zeroes/
 *
 * submission : faster than 94%
 */
package main

func moveZeroes(nums []int)  {
	j := 0
	for i := 0; i < len(nums);i++ {
		if nums[i] != 0 {
			nums[i], nums[j] = nums[j], nums[i]
			j++
		}
	}
}

