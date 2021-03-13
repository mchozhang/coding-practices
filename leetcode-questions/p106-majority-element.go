/**
 * LeetCode : Majority Element
 * https://leetcode.com/problems/majority-element/
 *
 * submission : faster than 96%
 */
package main

func majorityElement(nums []int) int {
	major := nums[0]
	count := 1
	for i := 1; i < len(nums); i++ {
		if count == 0 {
			major = nums[i]
			count = 1
		} else if nums[i] == major {
			count++
		} else {
			count--
		}
	}
	return major
}
