/**
 * LeetCode : Sum of Left Leaves
 * https://leetcode.com/problems/merge-two-sorted-lists/
 *
 * submission : faster than 100%
 */
package main

func addBinary(a string, b string) string {
	long, short := []byte(a), []byte(b)
	if len(b) > len(a) {
		long, short = short, long
	}
	sum := byte(0)
	for i, j := len(long)-1, len(short)-1; i > -1; i-- {
		sum += long[i] - '0'
		if j > -1 {
			sum += short[j] - '0'
			j--
		}
		long[i] = sum%2 + '0'
		sum /= 2
	}
	if sum != byte(0) {
		long = append([]byte{'1'}, long...)
	}
	return string(long)
}
