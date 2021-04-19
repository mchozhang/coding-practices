/**
 * LeetCode : Subarray Sum Equals K
 * https://leetcode.com/problems/subarray-sum-equals-k/
 *
 * submission : faster than 80%
 */
package main

func subarraySum(nums []int, k int) int {
	seen := map[int]int{0: 1}
	sum := 0
	count := 0
	for _, num := range nums {
		sum += num

		if value, isIn := seen[sum-k]; isIn {
			count += value
		}

		if _, isIn := seen[sum]; isIn {
			seen[sum]++
		} else {
			seen[sum] = 1
		}
	}
	return count
}
