/**
 * LeetCode : Reverse String
 * https://leetcode.com/problems/reverse-string/
 * Write a function that reverses a string. The input string is given as an array of characters char[].
 * Do not allocate extra space for another array, you must do this by modifying the input array in-place with O(1) extra memory.
 * submission : faster than 72%
 */
package main

func reverseString(s []byte)  {
	low := 0
	high := len(s) - 1
	for low < high {
		s[low], s[high] = s[high], s[low]
		low++
		high--
	}
}
