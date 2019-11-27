/**
 * LeetCode : two sum
 * https://leetcode.com/problems/two-sum/
 * Given an array of integers, return indices of the two numbers such that they add up to a specific target.
 * You may assume that each input would have exactly one solution, and you may not use the same element twice.
 * Example:
 * Given nums = [2, 7, 11, 15], target = 9,
 * Because nums[0] + nums[1] = 2 + 7 = 9,
 * return [0, 1].
 *
 * submission : faster than 100%
 */

package main

import "fmt"

func twoSum(nums []int, target int) []int {
	var result []int

	var arrMap = map[int]int{}
	for index, num := range nums {
		arrMap[num] = index
	}

	for i := 0; i < len(nums); i++ {
		var first = nums[i]
		var complement = target - first
		if val, ok := arrMap[complement]; ok && i != val {
			result = append(result, i)
			result = append(result, val)
			break
		}
	}

	return result
}

func main() {
	var arr = [] int{3, 2, 4, 15}
	var target = 6
	var result = twoSum(arr, target)

	for i := 0; i < len(result); i++ {
		fmt.Println(result[i])
	}
}
