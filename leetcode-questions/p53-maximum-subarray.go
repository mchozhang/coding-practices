/**
 * LeetCode : Maximum Subarray
 * https://leetcode.com/problems/maximum-subarray/
 * Given an integer array nums, find the contiguous subarray (containing at least one number) which has the largest sum and return its sum.
 *
 * submission : faster than 97%
 */
package main

func maxSubArray(nums []int) int {
	maxSoFar := make([]int, len(nums))
	maxEndHere := make([]int, len(nums))
	maxSoFar[0] = nums[0]
	maxEndHere[0] = nums[0]

	for i := 1; i < len(nums); i++ {
		if maxEndHere[i-1] > 0 {
			maxEndHere[i] = maxEndHere[i-1] + nums[i]
		} else {
			maxEndHere[i] = nums[i]
		}
		if maxEndHere[i] > maxSoFar[i-1]  {
			maxSoFar[i] = maxEndHere[i]
		} else {
			maxSoFar[i] = maxSoFar[i-1]
		}
	}

	return maxSoFar[len(nums) - 1]
}

func main() {
	nums := []int {-2,1,-3,4,-1,2,1,-5,4}

	// [4,-1,2,1] has the largest sum = 6
	println(maxSubArray(nums))
}
