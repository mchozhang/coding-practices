/**
 * LeetCode : Target Sum
 * https://leetcode.com/problems/target-sum/
 *
 * submission : faster than 54%
 */
package main

import "sort"

func findTargetSumWays(nums []int, target int) int {
	sort.Ints(nums)
	sum := 0
	for _, num := range nums {
		sum += num
	}
	if (sum + target)%2 == 1 {
		return 0
	}
	targetSum := (sum + target) / 2
	count := 0
	dfs(nums, 0, targetSum, 0, &count)
	return count
}

func dfs(nums []int, sum int, target int, pos int, count *int) {
	newSum := sum + nums[pos]
	if pos == len(nums) -1 {
		if sum == target {
			*count++
		}
		if newSum == target {
			*count++
		}
		return
	}
	if sum > target {
		return
	}
	dfs(nums, newSum, target, pos+1, count)
	dfs(nums, sum, target, pos+1, count)
}
