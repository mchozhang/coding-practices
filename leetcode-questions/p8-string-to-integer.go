/**
 * LeetCode: String to Integer
 * https://leetcode.com/problems/string-to-integer-atoi/
 *
 * submission: faster than 100%
 */
package main

import (
	"math"
	"strings"
)

func myAtoi(s string) int {
	s = strings.TrimSpace(s)
	if s == "" {
		return 0
	}
	base := 0
	sign := 1 // 1: positive -1: negative
	i := 0
	// find sign
	if s[i] == '-' {
		sign = -1
		i++
	} else if s[i] == '+' {
		i++
	}

	for ;i < len(s) && s[i] >= '0' && s[i] <= '9';i++ {
		base = 10 * base + int(s[i] - '0') * sign
		if base > math.MaxInt32 {
			return math.MaxInt32
		} else if base < math.MinInt32 {
			return math.MinInt32
		}
	}

	return base
}
