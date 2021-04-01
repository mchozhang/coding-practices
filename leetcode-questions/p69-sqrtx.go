/**
 * LeetCode : Sqrt(x)
 * https://leetcode.com/problems/sqrtx/
 *
 * submission : faster than 53%
 */
package main

func mySqrt(x int) int {
	if x < 2 {
		return x
	}
	low := 0
	high := x
	for low < high {
		mid := (high + low) / 2
		if low == mid {
			return low
		}
		res := mid * mid
		if res == x {
			return mid
		} else if res < x {
			low = mid
		} else {
			high = mid
		}
	}
	return low
}
