/**
 * 125. Valid Palindrome
 * https://leetcode.com/problems/valid-palindrome/
 *
 * submission: faster than 100%
 */
package main

func isPalindrome(s string) bool {
	low := 0
	high := len(s) - 1
	for low < high {
		if isAlphanumeric(s[low]) == false {
			low++
			continue
		}
		if isAlphanumeric(s[high]) == false {
			high--
			continue
		}
		if isDigit(s[low]) {
			if s[low] != s[high] {
				return false
			}
		} else if isLetter(s[low]) {
			if s[low] != s[high] && s[low]+'A'-'a' != s[high] && s[high]+'A'-'a' != s[low] {
				return false
			}
		}
		low++
		high--
	}
	return true
}

func isAlphanumeric(c uint8) bool {
	return isDigit(c) || isLetter(c)
}

func isLetter(c uint8) bool {
	return ('a' <= c && c <= 'z') || ('A' <= c && c <= 'Z')
}

func isDigit(c uint8) bool {
	return '0' <= c && c <= '9'
}
