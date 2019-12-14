/**
 * LeetCode : Rotate Array
 * https://leetcode.com/problems/rotate-array/
 * Given an array, rotate the array to the right by k steps, where k is non-negative.
 *
 * Example:
 * Input: [1,2,3,4,5,6,7] and k = 3
 * Output: [5,6,7,1,2,3,4]
 *
 * submission : faster than 92.70%
 */

package main

import "fmt"

func rotate(nums []int, k int) {
	for i:=0; i < k; i++ {
		last :=  nums[len(nums) - 1]
		copy(nums[1:],nums[:len(nums) - 1])
		nums[0] = last
	}
}

func rotate2(nums []int, k int)  {
	k %= len(nums)
	tmp := nums[len(nums) - k:]
	copy(nums, append(tmp, nums[:len(nums) - k]...))
}

func main() {
	array := [] int{1, 2, 3, 4, 5, 6, 7}

	rotate2(array, 3)
	fmt.Println(array)
}
