/**
 * LeetCode : Search in Rotated Sorted Array
 * https://leetcode.com/problems/search-in-rotated-sorted-array/
 *
 * submission : faster than 100%
 */
package main

import "fmt"

func search(nums []int, target int) int {
	low := 0
	high := len(nums) - 1
	// find the index of the smallest element
	for low < high {
		mid := (low + high) / 2
		if nums[mid] > nums[high] {
			low = mid + 1
		} else {
			high = mid
		}
	}
	// "low" is the index of the smallest element
	if target > nums[len(nums)-1] {
		high = low - 1
		low = 0
	} else {
		high = len(nums) - 1
	}
	for low < high {
		mid := (low + high) / 2
		if nums[mid] < target {
			low = mid + 1
		} else {
			high = mid
		}
	}

	if nums[low] == target {
		return low
	} else {
		return -1
	}
}

func main() {
	arr := []int{ 4, 5,1,2,3}
	fmt.Println(search(arr, 1))
}
