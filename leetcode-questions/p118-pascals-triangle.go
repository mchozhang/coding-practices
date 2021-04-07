/**
 * 118. Pascal's Triangle
 * https://leetcode.com/problems/pascals-triangle/
 *
 * submission: faster than 100%
 */
package main

func generate(numRows int) [][]int {
	res := make([][]int, 0)
	for i := 0; i < numRows; i++ {
		row := make([]int, i+1)
		row[0] = 1
		row[i] = 1
		for j := 1; j < i; j++ {
			row[j] = res[i-1][j-1] + res[i-1][j]
		}
		res = append(res, row)
	}
	return res
}
