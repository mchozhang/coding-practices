/**
 * LeetCode : longest substring without repeating character
 * https://leetcode.com/problems/longest-substring-without-repeating-characters/
 * Given a string, find the length of the longest substring without repeating characters.
 * Example 1:
 *
 * Input: "abcabcbb"
 * Output: 3
 * Explanation: The answer is "abc", with the length of 3.
 *
 * submission : faster than 70%
 */
package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	var max = 0
	var sub = map[uint8]int{}

	// j is the index of the substring starts
	var j = 0
	for i := 0; i < len(s); i++ {
		var letter = s[i]
		if target, ok := sub[letter]; ok && target >= j {
			j = target + 1
		}

		sub[letter] = i

		if i - j + 1 > max {
			max = i - j + 1
		}
	}
	return max
}

func main() {
	var s = "abcaefgh"
	fmt.Println(lengthOfLongestSubstring(s))


}
