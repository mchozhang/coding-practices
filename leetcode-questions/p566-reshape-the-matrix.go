/**
 * LeetCode: 566. Reshape the Matrix
 * https://leetcode.com/problems/reshape-the-matrix/
 *
 * submission: faster than 88%
 */
package main

func matrixReshape(mat [][]int, r int, c int) [][]int {
	if r * c != len(mat) * len(mat[0]) {
		return mat
	}
	res := make([][]int, r)
	for i := 0; i < r; i++ {
		res[i] = make([]int, c)
	}

	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[0]); j++ {
			num := i * len(mat[0]) + j
			res[num/c][num%c] = mat[i][j]
		}
	}
	return res
}
