/**
 * 119. Pascal's Triangle II
 * https://leetcode.com/problems/pascals-triangle-ii/
 *
 * submission: faster than 100%
 */
package main

func getRow(rowIndex int) []int {
	row := []int{1}
	for i := 1; i < rowIndex; i++ {
		newRow := make([]int, len(row)+1)
		newRow[0] = 1
		newRow[i] = 1
		for j := 1; j < i; j++ {
			newRow[j] = row[j] + row[j-1]
		}
		row = newRow
	}
	return row
}
