/**
 * LeetCode: Maximum Product Subarray
 * https://leetcode.com/problems/maximum-product-subarray/
 *
 * submission: faster than 100%
 */
package main

func maxProduct(nums []int) int {
	res := nums[0]
	max := res
	min := res
	for i := 1; i < len(nums); i++ {
		if nums[i] < 0 {
			min, max = max, min
		}

		max = nums[i] * max
		if nums[i] > max {
			max = nums[i]
		}

		min = nums[i] * min
		if nums[i] < min {
			min = nums[i]
		}

		if max > res {
			res = max
		}
	}
	return res
}
