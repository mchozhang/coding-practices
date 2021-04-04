/**
 * LeetCode: Isomorphic Strings
 * https://leetcode.com/problems/isomorphic-strings/
 *
 * submission : faster than 100%
 */
package main

func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	dict := map[uint8]uint8{}
	dict2 := map[uint8]uint8{}

	for i := 0; i < len(s); i++ {
		if c2, isIn := dict[s[i]]; isIn {
			if c2 != t[i] {
				return false
			}
		} else {
			// create new mapping
			if _, isIn := dict2[t[i]]; isIn {
				return false
			}
			dict[s[i]] = t[i]
			dict2[t[i]] = s[i]
		}
	}
	return true
}
