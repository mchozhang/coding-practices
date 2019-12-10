/**
 * LeetCode : Sudoku Solver
 * https://leetcode.com/problems/sudoku-solver/
 * Write a program to solve a Sudoku puzzle by filling the empty cells.
 *
 * submission : faster than 100%
 */

package main

import "fmt"

func solveSudoku(board [][]byte) {
	solve(board)
}

func solve(board [][]byte) bool {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				// place a number to the first empty cell
				for c := 0; c < 9; c++ {
					char := byte('1' + c)
					if isValid(board, i, j, char) {
						board[i][j] = char
						if solve(board) {
							return true
						} else {
							board[i][j] = '.'
						}
					}
				}
				return false
			}
		}
	}
	return true
}

func isValid(board [][]byte, row int, col int, c byte) bool {
	for i := 0; i < 9; i++ {
		//check column
		if board[i][col] != '.' && board[i][col] == c {
			return false
		}
		// check row
		if board[row][i] != '.' && board[row][i] == c {
			return false
		}
		// check block
		if board[3*(row/3)+i/3][ 3*(col/3)+i%3] != '.' &&
			board[3*(row/3)+i/3][3*(col/3)+i%3] == c {
			return false
		}
	}
	return true
}

func main() {
	board := [][] byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}

	solveSudoku(board)
	a := [9][9] int{}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			a[i][j] = int(board[i][j] - '0')
		}
		fmt.Println(a[i])
	}
}
