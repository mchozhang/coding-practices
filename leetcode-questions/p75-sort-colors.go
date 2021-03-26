/**
 * LeetCode : Sort Colors
 * https://leetcode.com/problems/sort-colors/
 *
 * submission : faster than 100%
 */
package main

func sortColors(nums []int) {
	counter0 := 0
	counter1 := 0
	counter2 := 0
	for _, num := range nums {
		if num == 0 {
			counter0++
		} else if num == 1 {
			counter1++
		} else {
			counter2++
		}
	}
	i := 0
	j := counter0
	k := counter0 + counter1
	for i < counter0 {
		if nums[i] == 0 {
			i++
		} else if nums[i] == 1 {
			nums[i], nums[j] = nums[j], nums[i]
			j++
		} else {
			nums[i], nums[k] = nums[k], nums[i]
			k++
		}
	}

	for j < counter0 + counter1 {
		if nums[j] == 1 {
			j++
		} else {
			nums[j], nums[k] = nums[k], nums[j]
			k++
		}
	}
}
