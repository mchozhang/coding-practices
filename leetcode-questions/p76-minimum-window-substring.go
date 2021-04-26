/**
 * LeetCode : Minimum Window Substring
 * https://leetcode.com/problems/minimum-window-substring/
 *
 * submission : faster than 100%
 */
package main

func minWindow(s string, t string) string {
	counter := make([]int, 'z'-'A'+1)
	diff := 0
	for _, c := range t {
		counter[c-'A']++
		diff++
	}
	low, high := 0, 0
	// right position for the sliding window
	for ; high < len(s); high++ {
		counter[s[high]-'A']--
		if counter[s[high]-'A'] >= 0 {
			diff--
		}
		if diff == 0 {
			break
		}
	}
	if diff != 0 {
		return ""
	}
	minLow, minHigh := low, high

	for ; high < len(s); {
		if diff == 0 {
			// adjust low position
			for low < high && counter[s[low]-'A'] < 0 {
				counter[s[low]-'A']++
				low++
			}
			if high-low < minHigh-minLow {
				minLow, minHigh = low, high
			}
		}
		if high < len(s)-1 {
			// shift right
			high++
			counter[s[high]-'A']--
			if counter[s[high]-'A'] >= 0 {
				diff--
			}
			counter[s[low]-'A']++
			if counter[s[low]-'A'] > 0 {
				diff++
			}
			low++
		} else {
			break
		}
	}
	return s[minLow : minHigh+1]
}
