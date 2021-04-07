/**
 * 796. Rotate String
 * https://leetcode.com/problems/rotate-string/
 *
 * submission: faster than 100%
 */
package main

import (
	"strings"
)

func rotateString(A string, B string) bool {
	return len(A) == len(B) && strings.Contains(A+A, B)
}

func rotateString2(A string, B string) bool {
	if len(A) != len(B) {
		return false
	}
	if len(A) == 0 && len(B) == 0 {
		return true
	}

	for i := 0; i < len(A); i++ {
		if A == B {
			return true
		}
		A = string(A[len(A)-1]) + A[:len(A)-1]
	}
	return false
}
