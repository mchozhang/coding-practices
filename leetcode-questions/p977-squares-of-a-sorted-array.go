/**
 * LeetCode : Squares of a Sorted Array
 * https://leetcode.com/problems/squares-of-a-sorted-array/
 *
 * submission : faster than 92%
 */
package main

func sortedSquares(nums []int) []int {
	res := make([]int, len(nums))
	positive := len(nums)
	for i, value := range nums {
		if value > 0 {
			positive = i
			break
		}
	}

	ans := 0
	i, j := positive, positive-1
	for i < len(nums) && j > -1 {
		if nums[i] < -nums[j] {
			res[ans] = nums[i]*nums[i]
			i++
		} else {
			res[ans] = nums[j]*nums[j]
			j--
		}
		ans++
	}
	for j > -1 {
		res[ans] = nums[j]*nums[j]
		j--
		ans++
	}
	for i < len(nums) {
		res[ans] = nums[i]*nums[i]
		ans++
		i++
	}
	return res
}
