/**
 * LeetCode : Valid Parentheses
 * https://leetcode.com/problems/valid-parentheses/
 * Given a string containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.
 * An input string is valid if:
 *	1. Open brackets must be closed by the same type of brackets.
 *	2. Open brackets must be closed in the correct order.
 *
 * Example:
 * Input: "()[]{}"
 * Output: true
 *
 * submission : faster than 100%
 */

package main

import "fmt"

func isValid(s string) bool {
	stack := ""
	for _, v := range s {
		if v == '[' || v == '(' || v == '{' {
			stack = stack + string(v)
		} else {
			if len(stack) == 0 {
				return false
			}

			last := stack[len(stack) - 1]
			if (v == ')' && last != '(') || (v == ']' && last != '[') || (v == '}' && last != '{') {
				return false
			} else {
				stack = stack[:len(stack) - 1]
			}
		}
	}
	return stack == ""
}

func main() {
	fmt.Println(isValid("()[]{}"))
	fmt.Println(isValid(""))
	fmt.Println(isValid("(}"))
	fmt.Println(isValid("}"))
	fmt.Println(isValid("("))
}
