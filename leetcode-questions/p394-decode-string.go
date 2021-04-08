/**
 * 394. Decode String
 * https://leetcode.com/problems/decode-string/
 *
 * submission: faster than 100%
 */
package main

func decodeString(s string) string {
	stack := []int{1}
	stringStack  := []string{""}
	base := 0
	for i := 0; i < len(s); i++ {
		c := s[i]

		if '0' <= c && c <= '9' {
			base = base * 10 + int(c-'0')
			if s[i+1] == '[' {
				stack = append(stack, base)
				stringStack = append(stringStack, "")
				base = 0
				i++
			}
		} else if c == ']' {
			k := stack[len(stack) - 1]
			stack = stack[:len(stack)-1]
			code := stringStack[len(stringStack) - 1]
			stringStack = stringStack[:len(stringStack) - 1]
			for j := 0; j < k; j++ {
				stringStack[len(stringStack)-1] += code
			}
		} else {
			stringStack[len(stringStack) - 1] += string(c)
		}
	}
	return stringStack[0]
}
