// LeetCode : Single Number
// https://leetcode.com/problems/ugly-number/
// Write a program to check whether a given number is an ugly number.
//
// Ugly numbers are positive numbers whose prime factors only include 2, 3, 5.
// Example:
// input: 6
// output: true, 6 = 2 * 3
// faster than 33%

package main
import "fmt"

func isUgly(num int) bool {
	if num < 1 {
		return false
	}
	for num % 2 == 0 {
		num /= 2
	}

	for num % 3 == 0 {
		num /= 3
	}

	for num % 5 == 0 {
		num /= 5
	}
	return num == 1
}

func main()  {
	fmt.Println(isUgly(0))

	fmt.Println(isUgly(6))

}
