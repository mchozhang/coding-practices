/**
 * LeetCode: Partition Labels
 * https://leetcode.com/problems/partition-labels/
 *
 * submission: faster than 100%
 */
package main

func partitionLabels(s string) []int {
	var res []int
	ends := map[uint8]int{}
	for i := 0; i < len(s); i++ {
		ends[s[i]] = i
	}

	part := 0
	for i := 0; i < len(s); {
		part = ends[s[i]] + 1
		for j := i + 1; j < part; j++ {
			if ends[s[j]] >= part {
				part = ends[s[j]] + 1
			}
		}
		res = append(res, part - i)
		i = part
	}
	return res
}
