/**
 * LeetCode : Subsets
 * https://leetcode.com/problems/subsets/
 *
 * submission : faster than 100%
 */
package main

import "fmt"

func subsets(nums []int) [][]int {
	res := make([][]int, 0)
	for i := 0; i <= len(nums); i++ {
		scan(nums, 0, []int{}, i, &res)
	}
	return res
}

func scan(nums []int, pos int, current []int, remain int, res *[][]int) {
	if remain == 0 {
		arr := make([]int, len(current))
		copy(arr, current)
		*res = append(*res, arr)
		return
	}

	if pos < len(nums) {
		// skip the current position
		scan(nums, pos+1, current, remain, res)

		// pick the current position
		newArr := append(current, nums[pos])
		scan(nums, pos+1, newArr, remain-1, res)
	}
}

func main() {
	arr := []int{0,1,2,3,4}
	fmt.Println(subsets(arr))
}
