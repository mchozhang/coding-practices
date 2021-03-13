/**
 * LeetCode : Plus One
 * https://leetcode.com/problems/plus-one/
 * Given a non-empty array of decimal digits representing a non-negative integer, increment one to the integer.
 * The digits are stored such that the most significant digit is at the head of the list, and each element in the array contains a single digit.
 *
 * submission : faster than 100%
 */
package main

func plusOne(digits []int) []int {
	i := len(digits)-1
	for i > -1 {
		if digits[i] != 9 {
			digits[i]++
			return digits
		} else {
			digits[i] = 0
			i--
		}
	}
	return append([]int{1}, digits...)
}
