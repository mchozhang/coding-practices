/**
 * LeetCode: Maximum Product of Three Numbers
 * https://leetcode.com/problems/maximum-product-of-three-numbers/
 *
 * submission: faster than 100%
 */
package main

import "math"

func maximumProduct(nums []int) int {
	max1, max2, max3 := math.MinInt32, math.MinInt32, math.MinInt32
	min1, min2 := math.MaxInt32, math.MaxInt32

	for i := 0; i < len(nums); i++ {
		if nums[i] > max1 {
			max1, max2, max3 = nums[i], max1, max2
		} else if nums[i] > max2 {
			max2, max3 = nums[i], max2
		} else if nums[i] > max3 {
			max3 = nums[i]
		}

		if nums[i] < min1 {
			min1, min2 = nums[i], min1
		} else if nums[i] < min2 {
			min2 = nums[i]
		}
	}

	product := max1 * max2 * max3
	if product < max1 * min1 * min2 {
		return max1 * min1 * min2
	}
	return product
}
