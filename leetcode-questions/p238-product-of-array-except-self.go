/**
 * LeetCode : Product of Array Except Self
 * https://leetcode.com/problems/product-of-array-except-self/
 *
 * submission : faster than 24%
 */
package main

func productExceptSelf(nums []int) []int {
	ans := make([]int, len(nums))
	product := 1
	count := 0
	onlyZero := -1
	for i, v := range nums {
		if v != 0 {
			product *= v
		} else {
			count++
			onlyZero = i
			if count > 1 {
				return ans
			}
		}
	}
	if count == 1 {
		ans[onlyZero] = product
		return ans
	}
	for i, _ := range ans {
		ans[i] = product / nums[i]
	}
	return ans
}
