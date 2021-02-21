/**
 * LeetCode : Single Number
 * https://leetcode.com/problems/top-k-frequent-elements/
 * Given a non-empty array of integers, return the k most frequent elements.
 * Example:
 * Input: nums = [1,1,1,2,2,3], k = 2
 * Output: [1,2]
 *
 * submission : faster than 100%
 */

package main

import "fmt"

func topKFrequent(nums []int, k int) []int {
	counter := map[int]int{}
	for _, num := range nums {
		if _, ok := counter[num]; ok {
			counter[num]++
		} else {
			counter[num] = 1
		}
	}

	frequencies := make([][]int, len(nums) + 1)
	for i, _ := range frequencies {
		frequencies[i] = make([]int, 0)
	}
	for num, frequency := range counter {
		frequencies[frequency] = append(frequencies[frequency], num)
	}
	result := make([]int, 0)
	j := 0
	for i := len(frequencies) - 1; i > -1 && j < k; i-- {
		if len(frequencies[i]) != 0 {
			result = append(result, frequencies[i]...)
			j += len(frequencies[i])
		}
	}
	return result
}

func main() {
	nums := []int{1, 1, 2, 2, 2, 3, 3, 3,3}
	// nums := []int{1}

	fmt.Println(topKFrequent(nums, 2))
}
