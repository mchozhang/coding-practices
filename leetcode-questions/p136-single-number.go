/**
 * LeetCode : Single Number
 * https://leetcode.com/problems/single-number/
 * Given a non-empty array of integers nums, every element appears twice except for one. Find that single one.
 *
 * Example:
 * Input: nums = [4,1,2,1,2]
 * Output: 4
 *
 * submission : faster than 100%
 */

package main

import "fmt"

func singleNumber(nums []int) int {
	xor := 0
	for _, num := range nums {
		xor ^= num
	}
	return xor
}

func main() {
	nums := [] int {1,2,2,1,4,3,3}
	fmt.Println(singleNumber(nums))
}
