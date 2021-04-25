/**
 * LeetCode : Combination Sum IV
 * https://leetcode.com/problems/combination-sum-iv/
 *
 * submission : faster than 100%
 */
package main

func combinationSum4(nums []int, target int) int {
	seen := map[int]int{}
	return findCombinations(target, nums, seen)
}

func findCombinations(target int, nums []int, seen map[int]int) int {
	if value, isIn := seen[target]; isIn {
		return value
	}
	if target == 0 {
		return 1
	}
	res := 0
	for _, num := range nums {
		if target >= num {
			res += findCombinations(target - num, nums, seen)
		}
	}
	seen[target] = res
	return res
}
