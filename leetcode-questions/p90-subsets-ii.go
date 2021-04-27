/**
 * LeetCode : Subsets II
 * https://leetcode.com/problems/subsets-ii/
 *
 * submission : faster than 100%
 */
package main

func subsetsWithDup(nums []int) [][]int {
	var res [][]int
	var counter = map[int]int{}
	var arr []int
	for _, num := range nums {
		counter[num]++
	}
	for i, _ := range counter {
		arr = append(arr, i)
	}
	for i := 0; i <= len(nums); i++ {
		findSubsets(0, nil, i, arr, counter, &res)
	}
	return res
}

func findSubsets(pos int, arr []int, size int, nums []int, counter map[int]int, res *[][]int) {
	if len(arr) == size {
		*res = append(*res, arr)
		return
	}
	if pos < len(nums) {
		for i := 0; i <= counter[nums[pos]] && len(arr)+i <= size; i++ {
			newArr := make([]int, len(arr)+i)
			copy(newArr, arr)
			for j := 0; j < i; j++ {
				newArr[len(arr)+j] = nums[pos]
			}
			findSubsets(pos+1, newArr, size, nums, counter, res)
		}
	}
}
