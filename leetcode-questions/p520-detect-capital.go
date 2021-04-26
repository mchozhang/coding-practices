/**
 * LeetCode : Detect Capital
 * https://leetcode.com/problems/detect-capital/
 *
 * submission : faster than 100%
 */
package main

func detectCapitalUse(word string) bool {
	if len(word) == 1 {
		return true
	}
	sign := isUpperCase(word[0]) && isUpperCase(word[len(word)-1])
	for i := 1; i < len(word); i++ {
		if isUpperCase(word[i]) != sign {
			return false
		}
	}
	return true
}

func isUpperCase(char uint8) bool {
	return 'A' <= char && char <= 'Z'
}
