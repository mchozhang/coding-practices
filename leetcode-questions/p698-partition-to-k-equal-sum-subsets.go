/**
 * LeetCode : Partition to K Equal Sum Subsets
 * https://leetcode.com/problems/partition-to-k-equal-sum-subsets/
 *
 * submission : faster than 0%
 */
package main

import "sort"

func canPartitionKSubsets(nums []int, k int) bool {
	sum := 0
	for _, value := range nums {
		sum += value
	}
	if sum%k != 0 {
		return false
	}
	sum /= k
	sort.Sort(sort.Reverse(sort.IntSlice(nums)))
	visited := make([]bool, len(nums))
	return dfs(k, 0, 0, sum, visited, nums)
}

func dfs(k int, sum int, pos int, target int, visited []bool, nums []int) bool {
	if k == 1 {
		return true
	}
	if sum == target {
		return dfs(k-1, 0, 0, target, visited, nums)
	}

	for i := pos; i < len(visited); i++ {
		newSum := sum + nums[i]
		if visited[i] == false && newSum <= target {
			visited[i] = true
			if dfs(k, newSum, i+1, target, visited, nums) {
				return true
			}
			visited[i] = false
		}
	}
	return false
}
