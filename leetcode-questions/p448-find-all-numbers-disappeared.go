/**
 * LeetCode : Find All Disappeared in an Array
 * https://leetcode.com/problems/find-all-numbers-disappeared-in-an-array/
 *
 * submission : faster than 99%
 */
package main

func findDisappearedNumbers(nums []int) []int {
	result := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		result[i] = i + 1
	}
	for _, value := range nums {
		result[value-1] = 0
	}
	high := len(result) - 1
	low := 0
	for low <= high {
		if result[low] == 0 {
			result[low], result[high] = result[high], result[low]
			high--
		} else {
			low++
		}
	}
	return result[:high+1]
}
