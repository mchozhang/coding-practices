/**
 * LeetCode : Shortest Unsorted Continuous Subarray
 * https://leetcode.com/problems/shortest-unsorted-continuous-subarray/
 *
 * submission : faster than 92%
 */
package main

func findUnsortedSubarray(nums []int) int {
	if len(nums) < 2 {
		return 0
	}
	left, right := 0, len(nums)-1
	for left < len(nums)-1 && nums[left] <= nums[left+1] {
		left++
	}
	for right > 0 && nums[right] >= nums[right-1] {
		right--
	}

	if left > right {
		return 0
	}
	max, min := nums[left], nums[left]
	for i := left + 1; i <= right; i++ {
		if nums[i] < min {
			min = nums[i]
		} else if nums[i] > max {
			max = nums[i]
		}
	}

	for left > 0 && nums[left-1] > min {
		left--
	}

	for right < len(nums)-1 && nums[right+1] < max {
		right++
	}
	return right - left + 1
}

/**
 * binary search, 28ms
 */
func findUnsortedSubarray2(nums []int) int {
	max := nums[0]
	length := 0
	start, end := -1, -1
	for i := 1; i < len(nums); i++ {
		if nums[i] < max {
			if length == 0 || nums[i] < nums[start] {
				// binary search the appropriate start position
				low, high := 0, i-1
				if start != -1 {
					high = start
				}
				for low < high {
					mid := (low + high) / 2
					if nums[mid] <= nums[i] {
						low = mid + 1
					} else {
						high = mid
					}
				}
				length = i - low + 1
				start = low
			} else {
				length += i - end
			}
			end = i
		} else {
			max = nums[i]
		}
	}
	return length
}
