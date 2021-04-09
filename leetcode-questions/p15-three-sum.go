/**
 * LeetCode : 3 Sum
 * https://leetcode.com/problems/three-sum/
 *
 * submission : faster than 100%
 */
package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	var ans [][]int
	sort.Ints(nums)
	for i, value := range nums {
		if i != 0 && nums[i-1] == value {
			continue
		}
		if value > 0 {
			break
		}
		target := -value
		low := i + 1
		high := len(nums) - 1
		for low < high {
			sum := nums[low]+nums[high]
			if sum == target {
				ans = append(ans, []int{value, nums[low], nums[high]})
				low++
				high--
				for low < high && nums[low] == nums[low-1] {
					low++
				}
				for low < high && nums[high] == nums[high+1] {
					high--
				}
			} else if sum < target {
				low++
			} else {
				high--
			}
		}
	}
	return ans
}


func main() {
	nums := []int{-1, 0, 1, 2, -1, -4}
	// nums := []int{-6, -6, 2, 2, 3, 3, 3, 3, 4, 4, 5}
	fmt.Println(threeSum(nums))
}
