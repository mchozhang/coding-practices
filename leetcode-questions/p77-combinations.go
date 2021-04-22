/**
 * LeetCode : Combinations
 * https://leetcode.com/problems/combinations/
 *
 * submission : faster than 100%
 */
package main

func combine(n int, k int) [][]int {
	var res [][]int
	findCombinations(1, n, k, []int{}, &res)
	return res
}

func findCombinations(start int, n int, k int, arr []int, res *[][]int) {
	if len(arr) == k {
		*res = append(*res, arr)
		return
	}
	if k-len(arr) > n-start+1 {
		return
	}
	newArr := make([]int, len(arr)+1)
	copy(newArr, arr)
	newArr[len(newArr)-1] = start
	findCombinations(start+1, n, k, newArr, res)
	findCombinations(start+1, n, k, arr, res)
}
