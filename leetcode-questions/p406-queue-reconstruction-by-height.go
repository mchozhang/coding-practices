/**
 * LeetCode : Queue Reconstruction by Height
 * https://leetcode.com/problems/queue-reconstruction-by-height/
 *
 * submission : faster than 56%
 */
package main

import "sort"

func reconstructQueue(people [][]int) [][]int {
	sort.Slice(people, func(i, j int) bool {
		if people[i][0] == people[j][0] {
			return people[i][1] < people[j][1]
		} else {
			return people[i][0] > people[j][0]
		}
	})

	var queue [][]int
	for i := 0; i < len(people); i++ {
		insertIndex := people[i][1]
		queue = append(queue, []int {})
		copy(queue[insertIndex+1:], queue[insertIndex:])
		queue[insertIndex] = people[i]
	}
	return queue
}
