/**
 * LeetCode : Length of Last Word
 * https://leetcode.com/problems/length-of-last-word/
 * Given a string s consists of some words separated by spaces, return the length of the last word in the string. If the last word does not exist, return 0.
 * A word is a maximal substring consisting of non-space characters only.
 * submission : faster than 100%
 */
package main

import "fmt"

func lengthOfLastWord(s string) int {
	length := 0
	for i := len(s) - 1; i > -1; i-- {
		if s[i] == ' ' {
			if length == 0 {
				continue
			} else {
				break
			}
		} else {
			length++
		}
	}
	return length
}

func main() {
	fmt.Println(lengthOfLastWord("hello world"))
	fmt.Println(lengthOfLastWord("hello "))
	fmt.Println(lengthOfLastWord("1  "))
}
