/**
 * LeetCode : Merge Sorted Array
 * https://leetcode.com/problems/merge-sorted-array/
 *
 * submission : faster than 100%
 */
package main

func merge(nums1 []int, m int, nums2 []int, n int)  {
	i := m - 1
	j := n - 1
	k := m + n - 1
	for i > -1 && j > -1 {
		if nums1[i] > nums2[j] {
			nums1[k] = nums1[i]
			i--
		} else {
			nums1[k] = nums2[j]
			j--
		}
		k--
	}
	for ;j > -1;j-- {
		nums1[k] = nums2[j]
		k--
	}
}
