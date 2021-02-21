// LeetCode : Single Number
// https://leetcode.com/problems/ugly-number-ii/
// Write a program to find the n-th ugly number.
// Ugly numbers are positive numbers whose prime factors only include 2, 3, 5.
//
// Example:
// input: n = 10
// output: 12,
// 1, 2, 3, 4, 5, 6, 8, 9, 10, 12 is the sequence of the first 10 ugly numbers.
// faster than 100%

package main

import "fmt"

func nthUglyNumber(n int) int {
	queue2 := make([]int, 0)
	queue3 := make([]int, 0)
	queue5 := make([]int, 0)
	result := 1

	for i := 0; i < n - 1; i++ {
		queue2 = append(queue2, result * 2)
		queue3 = append(queue3, result * 3)
		queue5 = append(queue5, result * 5)

		head2 := queue2[0]
		head3 := queue3[0]
		head5 := queue5[0]
		min := head2
		if head3 < min {
			min = head3
		}

		if head5 < min {
			min = head5
		}

		if head2 == min {
			queue2 = queue2[1:]
		}

		if head3 == min {
			queue3 = queue3[1:]
		}

		if head5 == min {
			queue5 = queue5[1:]
		}

		result = min
	}
	return result

}

func main() {
	fmt.Println(nthUglyNumber(7))
}