/**
 * LeetCode : Maximum Average Subarray
 * https://leetcode.com/problems/maximum-average-subarray-i/
 *
 * submission : faster than 83%
 */
package main

func findMaxAverage(nums []int, k int) float64 {
	max := 0
	for i := 0; i < k; i++ {
		max += nums[i]
	}
	runningSum := max
	for i := k; i < len(nums); i++ {
		runningSum += nums[k] - nums[i-k]
		if runningSum > max {
			max = runningSum
		}
	}
	return float64(max) / float64(k)
}