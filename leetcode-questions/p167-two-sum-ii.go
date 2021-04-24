/**
 * LeetCode : Two Sum II - Input array is sorted
 * https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/
 *
 * submission : faster than 97%
 */
package main

func twoSum(nums []int, target int) []int {
	low, high := 0, len(nums)-1
	for low < high {
		sum := nums[low] + nums[high]
		if sum == target {
			return []int{low+1, high+1}
		}
		if sum > target {
			high--
		} else {
			low++
		}
	}
	return nil
}

func twoSum2(numbers []int, target int) []int {
	start, end := 0, 1
	seen := map[int]int{target - numbers[0]: 0}
	for i := 1; i < len(numbers); i++ {
		num := numbers[i]
		if num == numbers[i-1] {
			continue
		}
		if value, isIn := seen[num]; isIn {
			start, end = value+1, i+1
		} else {
			seen[target-num] = i
		}
	}
	return []int{start, end}
}
