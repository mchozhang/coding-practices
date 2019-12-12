/**
 * LeetCode : Missing Number
 * https://leetcode.com/problems/missing-number/
 * Given an array containing n distinct numbers taken from 0, 1, 2, ..., n, find the one that is missing from the array.
 *
 * Example:
 * Input: [9,6,4,2,3,5,7,0,1]
 * Output: 8
 *
 * submission : faster than 92.16%
 */

package main

import "fmt"

func missingNumber(nums []int) int {
	all := len(nums) * (len(nums) + 1) / 2
	sum := 0
	for _, v := range nums {
		sum += v
	}
	return all - sum
}

func main() {
	var array = [] int{9,6,4,2,3,5,7,0,1}
	fmt.Println(missingNumber(array))
}
