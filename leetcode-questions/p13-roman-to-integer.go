/**
 * LeetCode: Roman to Integer
 * https://leetcode.com/problems/roman-to-integer/
 *
 * submission: faster than 61%
 */
package main

func romanToInt(s string) int {
	scores := map[uint8]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}
	sum := scores[s[0]]
	for i := 1; i < len(s); i++ {
		if scores[s[i]] > scores[s[i-1]] {
			sum += scores[s[i]] - 2 * scores[s[i-1]]
		} else {
			sum += scores[s[i]]
		}
	}
	return sum
}
