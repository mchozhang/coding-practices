/**
 * LeetCode : 3Sum Closest Set
 * https://leetcode.com/problems/3sum-closest/
 *
 * submission : faster than 98%
 */
package main

import "sort"

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	res := nums[0] + nums[1] + nums[2]
	min :=  res - target
	if min < 0 {
		min = -min
	}
	for i, value := range nums {
		if i != 0 && value == nums[i-1] {
			continue
		}

		subTarget := target - value
		low := i + 1
		high := len(nums) - 1
		for low < high {
			sum := nums[low] + nums[high]
			diff := subTarget - sum
			if diff == 0 {
				return sum + value
			} else if diff > 0 {
				if diff < min {
					min = diff
					res = sum + value
				}
				low++
			} else {
				if -diff < min {
					min = -diff
					res = sum + value
				}
				high--
			}
		}
	}
	return res
}
