/**
 * LeetCode : Continuous Subarray Sum
 * https://leetcode.com/problems/continuous-subarray-sum/
 *
 * submission : faster than 93%
 */

package main

func checkSubarraySum(nums []int, k int) bool {
	remainders := map[int] int{}
	sum := 0
	for i, num := range nums {
		sum += num
		if k != 0 {
			sum %= k
		}
		if v, isIn := remainders[sum]; isIn {
			if v - i > 1 {
				return true
			}
		} else {
			remainders[sum] = i
		}
	}
	return false
}
