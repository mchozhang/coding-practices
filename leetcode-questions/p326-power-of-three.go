/**
 * LeetCode : Power of Three
 * https://leetcode.com/problems/power-of-three/
 *
 * submission : faster than 100%
 */
package main

import "math"

func isPowerOfThree(n int) bool {
	max := int(math.Pow(3, 19))
	return n > 0 && max % n == 0
}