/**
 * LeetCode : Counting Bits
 * https://leetcode.com/problems/counting-bits/
 *
 * submission : faster than 94%
 */
package main

func countBits(num int) []int {
	res := make([]int, num+1)
	for i := 1; i <= num; i++ {
		res[i] = res[i&(i-1)] + 1
	}
	return res
}
