/**
 * LeetCode : Balanced Binary Tree
 * https://leetcode.com/problems/balanced-binary-tree/
 *
 * submission : faster than 100%
 */
package main

func addStrings(num1 string, num2 string) string {
	long, short := []byte(num1), []byte(num2)
	if len(num1) < len(num2) {
		long, short = short, long
	}

	sum := byte(0)
	for i, j := len(long) - 1, len(short) - 1; i > - 1; i-- {
		sum += long[i] - '0'
		if j > -1 {
			sum += short[j] - '0'
			j--
		}
		long[i] = sum%10 + '0'
		sum /= 10
	}
	if sum != byte(0) {
		long = append([]byte{'1'}, long...)
	}
	return string(long)
}
