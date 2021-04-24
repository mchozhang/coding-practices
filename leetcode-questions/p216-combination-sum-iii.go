/**
 * LeetCode : Combination Sum III
 * https://leetcode.com/problems/combination-sum-iii/
 *
 * submission : faster than 100%
 */
package main

func combinationSum3(k int, n int) [][]int {
	var res [][]int
	findCombinations(1, k, n, 0, []int{}, &res)
	return res
}

func findCombinations(pos int, k int, n int, sum int, arr []int, res *[][]int) {
	if sum == n && len(arr) == k {
		*res = append(*res, arr)
		return
	}
	if k-len(arr) > 9-pos+1 || sum > n {
		return
	}
	newArr := make([]int, len(arr)+1)
	copy(newArr, arr)
	newArr[len(newArr)-1] = pos
	findCombinations(pos+1, k, n, sum+pos, newArr, res)
	findCombinations(pos+1, k, n, sum, arr, res)
}

