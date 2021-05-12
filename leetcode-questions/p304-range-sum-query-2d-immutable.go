/**
 * LeetCode : Range Sum Query 2D - Immutable
 * https://leetcode.com/problems/range-sum-query-2d-immutable/
 *
 * submission : faster than 93%
 */

type NumMatrix struct {
	Matrix [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	obj := NumMatrix{}
	obj.Matrix = make([][]int, len(matrix)+1)
	for i, _ := range obj.Matrix {
		obj.Matrix[i] = make([]int, len(matrix[0])+1)
	}
	for i, row := range matrix {
		for j, value := range row {
			obj.Matrix[i+1][j+1] = value + obj.Matrix[i+1][j] + obj.Matrix[i][j+1] - obj.Matrix[i][j]
		}
	}
	return obj
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	return this.Matrix[row2+1][col2+1] - this.Matrix[row1][col2+1] - this.Matrix[row2+1][col1] + this.Matrix[row1][col1]
}