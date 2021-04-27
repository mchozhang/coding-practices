/**
 * LeetCode : Split Array into Fibonacci Sequence
 * https://leetcode.com/problems/split-array-into-fibonacci-sequence/
 *
 * submission : faster than 100%
 */
package main

import (
	"math"
	"strconv"
)

func splitIntoFibonacci(s string) []int {
	var res []int
	dfs(0, nil, s, &res)
	return res
}

func dfs(start int, arr []int, s string, res *[]int) bool {
	if len(arr) > 1 {
		target := (arr)[len(arr)-1] + arr[len(arr)-2]
		targetLen := 1
		for i := target; i/10 != 0; i /= 10 {
			targetLen++
		}
		if start+targetLen-1 < len(s) {
			numStr := s[start : start+targetLen]
			if len(numStr) > 1 && numStr[0] == '0' {
				return false
			}
			num, _ := strconv.Atoi(numStr)
			if num > math.MaxInt32 {
				return false
			}
			if num == target {
				if start + targetLen == len(s) {
					*res = arr
					return true
				}
				arr = append(arr, target)
				if dfs(start+targetLen, arr, s, res) {
					return true
				}
				arr = arr[:len(arr)-1]
			}
		}
	} else {
		for i := 0; i < 11 && start+i < len(s); i++ {
			numStr := s[start : start+i+1]
			if len(numStr) > 1 && numStr[0] == '0' {
				return false
			}
			num, _ := strconv.Atoi(numStr)
			if num < math.MaxInt32 {
				arr = append(arr, num)
				if dfs(start+i+1, arr, s, res) {
					return true
				}
				arr = arr[:len(arr)-1]
			}
		}
	}
	return false
}