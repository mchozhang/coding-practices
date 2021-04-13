/**
 * LeetCode: Binary Search
 * https://leetcode.com/problems/binary-search/
 *
 * submission: faster than 63%
 */
package main

func search(nums []int, target int) int {
	low, high := 0, len(nums)-1
	for low < high {
		mid := (low + high + 1) / 2
		if nums[mid] > target {
			high = mid - 1
		} else {
			low = mid
		}
	}
	if nums[low] == target {
		return low
	}
	return -1
}
