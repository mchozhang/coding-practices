/**
 * LeetCode : Rotate Image
 * https://leetcode.com/problems/rotate-image/
 *
 * submission : faster than 99%
 */
package main
import "strings"

func groupAnagrams(strs []string) [][]string {
	dict := map[string]int{}
	result := make([][]string, 0)
	for _, str := range strs {
		key := convertStr(str)
		if index, isIn := dict[key]; isIn {
			result[index] = append(result[index], str)
		} else {
			group := []string {str}
			result = append(result, group)
			dict[key] = len(result) - 1
		}
	}
	return result
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
