/**
 * LeetCode : two sum
 * https://leetcode.com/problems/two-sum/
 * Given an array of integers, return indices of the two numbers such that they add up to a specific target.
 * You may assume that each input would have exactly one solution, and you may not use the same element twice.
 * Example:
 * Given nums = [2, 7, 11, 15], target = 9,
 * Because nums[0] + nums[1] = 2 + 7 = 9,
 * return [0, 1].
 */

package main

import "fmt"

func twoSum(nums []int, target int) []int {
	var result []int

	for i := 0; i < len(nums); i++ {
		var first = nums[i]
		for j := i + 1; j < len(nums); j++ {
			var second = nums[j]
			if first + second == target {
				result = append(result, i)
				result = append(result, j)
			}
		}
	}

	return result
}

func main() {
	var arr = [] int {2, 7, 11, 15}
	var target = 9
	var result = twoSum(arr, target)

	for i := 0 ; i < len(result); i++ {
		fmt.Println(result[i])
	}
}