/**
 * LeetCode: Valid Anagram
 * https://leetcode.com/problems/valid-anagram/
 *
 * submission: faster than 100%
 */
package main

import (
	"reflect"
	"strings"
)

/**
 * faster, only works for English letters
 */
func isAnagram(s string, t string) bool{
	return convertStr(s) == convertStr(t)
}

func convertStr(str string) string{
	counter := [26]int{}
	for _, c := range str {
		counter[c-'a']++
	}
	var stringBuilder strings.Builder
	for _, count := range counter {
		stringBuilder.WriteRune(rune('a' + count))
	}
	return stringBuilder.String()
}

func isAnagram2(s string, t string) bool {
	return reflect.DeepEqual(countLetters(s), countLetters(t))
}

func countLetters(s string) map[uint8]int {
	letters := map[uint8]int{}
	for i := 0; i < len(s); i++ {
		if _, isIn := letters[s[i]]; isIn {
			letters[s[i]]++
		} else {
			letters[s[i]] = 1
		}
	}
	return letters
}

