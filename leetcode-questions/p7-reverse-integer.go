/**
 * LeetCode : reverse integer
 * https://leetcode.com/problems/reverse-integer/
 * Given a 32-bit signed integer, reverse digits of an integer.
 *
 * Example:
 * Input: 123
 * Output: 321
 *
 * submission : faster than 100%
 */

package main

import (
	"fmt"
)

func reverse(x int) int {
	result := 0
	for x != 0 {
		result *= 10
		result += x % 10

		// result not in [- 2^31, 2^31 - 1] is overflow
		if result > 2147483647  || result < -2147483648 {
			return 0
		}
		x /= 10
	}
	return result
}

func main() {
	a := 120
	fmt.Println(reverse(a))
}
