/**
 * LeetCode: Longest Increasing Subsequence
 * https://leetcode.com/problems/longest-increasing-subsequence/
 *
 * submission: faster than 99%
 */
package main

func lengthOfLIS(nums []int) int {
	var sorted = []int{nums[0]}

	for i := 1; i < len(nums); i++ {
		if nums[i] > sorted[len(sorted)-1] {
			sorted = append(sorted, nums[i])
		} else {
			low, high := 0, len(sorted)-1
			found := false
			for low <= high {
				mid := (low + high) / 2
				if sorted[mid] == nums[i] {
					found = true
					break
				}
				if sorted[mid] > nums[i] {
					high = mid - 1
				} else if nums[i] > sorted[mid] {
					low = mid + 1
				}
			}
			if !found {
				sorted[low] = nums[i]
			}
		}
	}
	return len(sorted)
}
