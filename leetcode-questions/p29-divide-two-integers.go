/**
 * LeetCode : Divide Two Integers
 * https://leetcode.com/problems/divide-two-integers/
 *
 * submission : faster than 95%
 */
package main

import "math"

func divide(dividend int, divisor int) int {
	sign := 1
	if dividend < 0 {
		dividend = -dividend
		sign = -1
	}
	if divisor < 0 {
		divisor = -divisor
		sign *= -1
	}
	ans := div(dividend, divisor) * sign
	if ans < math.MinInt32 {
		return math.MinInt32
	}
	if ans > math.MaxInt32 {
		return math.MaxInt32
	}
	return ans
}

func div(dividend int, divisor int) int {
	if divisor > dividend {
		return 0
	}
	sum := divisor
	multiple := 1
	for sum + sum <= dividend {
		sum += sum
		multiple += multiple
	}
	return multiple + div(dividend-sum, divisor)
}