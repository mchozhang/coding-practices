/**
 * LeetCode : Find All Anagrams in a String
 * https://leetcode.com/problems/find-all-anagrams-in-a-string/
 *
 * submission : faster than 100%
 */
package main

func findAnagrams(s string, p string) []int {
	if len(p) > len(s) {
		return []int{}
	}

	var res []int
	chars := make([]int, 26)
	diff := len(p)
	for i := 0; i < len(p); i++ {
		chars[p[i]-'a']++
	}
	for i := 0; i < len(p);i++ {
		c := s[i]
		chars[c-'a']--
		if chars[c-'a'] >= 0 {
			diff--
		}
	}
	if diff == 0 {
		res = append(res, 0)
	}

	pre := s[0]
	for i := len(p); i < len(s); i++ {
		if chars[pre-'a'] >= 0 {
			diff++
		}
		chars[pre-'a']++

		chars[s[i]-'a']--
		if chars[s[i]-'a'] >= 0 {
			diff--
		}

		if diff == 0 {
			res = append(res, i-len(p)+1)
		}
		pre = s[i-len(p)+1]
	}
	return res
}
