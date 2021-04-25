/**
 * LeetCode : Task Scheduler
 * https://leetcode.com/problems/counting-bits/
 *
 * submission : faster than 93%
 */
package main

import "sort"

func leastInterval(tasks []byte, n int) int {
	counter := make([]int, 26)
	for _, task := range tasks {
		counter[task-'A']++
	}
	sort.Ints(counter)
	maxFrequency := counter[25]
	patternLen := 0
	for i := 25; i > -1 && counter[i] == maxFrequency; i-- {
		patternLen++
	}
	if len(tasks) > (n+1)*(maxFrequency-1)+patternLen {
		return len(tasks)
	}
	return (n+1)*(maxFrequency-1) + patternLen
}
