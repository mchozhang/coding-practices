/**
 * LeetCode : Partition to K Equal Sum Subsets
 * https://leetcode.com/problems/partition-to-k-equal-sum-subsets/
 *
 * submission : faster than 82%
 */
package main

func canPartitionKSubsets(nums []int, k int) bool {
	sum := 0
	for _, value := range nums {
		sum += value
	}
	if sum%k != 0 {
		return false
	}
	dp := make([][]bool, k)
	for i := 0; i < k; i++ {
		dp[i] = make([]bool, sum+1)
	}
	dp[k]

	for _, value := range nums {
		for i := sum; i >= value; i-- {

		}
	}
	return dp[sum]
}
