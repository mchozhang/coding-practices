/**
 * LeetCode : Insert Interval
 * https://leetcode.com/problems/insert-interval/
 *
 * submission : faster than 80%
 */
package main

func insert(intervals [][]int, newInterval []int) [][]int {
	var res [][]int
	i := 0
	for i < len(intervals) && newInterval[0] > intervals[i][1] {
		i++
	}
	res = make([][]int, i)
	copy(res, intervals[0:i])
	for i < len(intervals) && newInterval[1] >= intervals[i][0] {
		if intervals[i][0] < newInterval[0] {
			newInterval[0] = intervals[i][0]
		}
		if intervals[i][1] > newInterval[1] {
			newInterval[1] = intervals[i][1]
		}
		i++
	}
	res = append(res, newInterval)
	res = append(res, intervals[i:]...)
	return res
}


func insert2(intervals [][]int, newInterval []int) [][]int {
 	var res [][]int
	start, end := -1, -1
	newStart, newEnd := newInterval[0], newInterval[1]
	for i, interval := range intervals {
		// determine the start position
		if start == -1 && newInterval[0] <= interval[1] {
			start = i
			if newInterval[0] >= interval[0] {
				newStart = interval[0]
			}
			// put front unchanged intervals to res
			if i > 0 {
				res = make([][]int, i)
				copy(res, intervals[:i])
			}
		}
		// determine the end position
		if newInterval[1] <= interval[1] {
			end = i
			if newInterval[1] >= interval[0] {
				newEnd = interval[1]
			}
			break
		}
	}

	if start == -1 {
		return append(intervals, newInterval)
	}
	res = append(res, []int{newStart, newEnd})
	if end != -1 {
		if newEnd == newInterval[1] {
			res = append(res, intervals[end:]...)
		} else {
			res = append(res, intervals[end+1:]...)
		}
	}
	return res
}
