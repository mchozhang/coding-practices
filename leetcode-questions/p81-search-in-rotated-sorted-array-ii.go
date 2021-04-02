/**
 * LeetCode : Search in Rotated Sorted Array II
 * https://leetcode.com/problems/search-in-rotated-sorted-array-ii/
 *
 * submission : faster than 90%
 */
package main

import "fmt"

func search(nums []int, target int) bool {
	// remove duplicates
	for i := 1; i < len(nums); {
		if nums[i] == nums[i-1] {
			nums = append(nums[:i], nums[i+1:]...)
		} else {
			i++
		}
	}

	// find the index of the smallest element
	low := 0
	high := len(nums) - 1
	for low < high {
		mid := (low + high) / 2
		if nums[mid] >= nums[high] {
			low = mid + 1
		} else {
			high = mid
		}
	}
	if target > nums[len(nums)-1] {
		high = low - 1
		low = 0
	} else {
		high = len(nums) - 1
	}
	for low < high {
		mid := (low + high) / 2
		if target == nums[mid] {
			return true
		} else if target > nums[mid] {
			low = mid + 1
		} else {
			high = mid
		}
	}
	return nums[low] == target
}

func main() {

	arr := []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 1, 1, 1, 1}
	fmt.Println(search(arr, 2))
}
