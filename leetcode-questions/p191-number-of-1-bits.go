/**
 * LeetCode : Number of 1 Bits
 * https://leetcode.com/problems/number-of-1-bits/
 *
 * submission : faster than 94%
 */
package main

func hammingWeight(num uint32) int {
	count := 0
	for num != 0 {
		num &= num - 1
		count++
	}
	return count
}
