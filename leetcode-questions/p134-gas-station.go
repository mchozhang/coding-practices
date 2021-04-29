/**
 * LeetCode : Gas Station
 * https://leetcode.com/problems/gas-station/
 *
 * submission : faster than 90%
 */
package main

func canCompleteCircuit(gas []int, cost []int) int {
	sum, total, start := 0, 0, 0
	for i := 0; i < len(gas);i++ {
		diff := gas[i] - cost[i]
		total += diff
		sum += diff
		if sum < 0 {
			sum = 0
			start = i + 1
		}
	}
	if total < 0 {
		return -1
	}
	return start
}


