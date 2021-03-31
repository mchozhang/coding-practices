/**
 * LeetCode : Range Sum Query - Immutable
 * https://leetcode.com/problems/range-sum-query-immutable/
 *
 * submission : faster than 95%
 */
package main

type NumArray struct {
	Array []int
}


func Constructor(nums []int) NumArray {
	arr := make([]int, len(nums))
	if len(nums) > 0 {
		arr[0] = nums[0]
	}
	for i:=1;i < len(arr);i++ {
		arr[i] = nums[i] + arr[i-1]
	}
	return NumArray{Array: arr}
}


func (this *NumArray) SumRange(left int, right int) int {
	if left == 0 {
		return this.Array[right]
	}
	return this.Array[right] - this.Array[left-1]
}