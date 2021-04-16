/**
 * LeetCode : Partition Equal Subset Sum
 * https://leetcode.com/problems/partition-equal-subset-sum/
 *
 * submission : faster than 82%
 */
package main

func canPartition(nums []int) bool {
	sum := 0
	for _, num := range nums {
		sum += num
	}

	if sum%2 == 1 {
		return false
	}
	sum /= 2
	dp := make([]bool, sum+1)
	dp[0] = true
	for _, num := range nums {
		for i := sum; i >= num; i-- {
			dp[i] = dp[i] || dp[i-num]
		}
	}
	return dp[sum]
}

/**
 * BFS, 928ms
 */
func canPartition2(nums []int) bool {
	if len(nums) < 2 {
		return false
	}
	queue := map[int]int{nums[0]: 0}

	for i := 1; i < len(nums)-1; i++ {
		temp := map[int]int{}
		for value, _ := range queue {
			diff1 := value + nums[i]
			diff2 := value - nums[i]
			if diff2 < 0 {
				diff2 = -diff2
			}
			if _, isIn := temp[diff1]; !isIn {
				temp[diff1] = 0
			}
			if _, isIn := temp[diff2]; !isIn {
				temp[diff2] = 0
			}
		}
		queue = temp
	}

	for num, _ := range queue {
		if num == nums[len(nums)-1] {
			return true
		}
	}
	return false
}
