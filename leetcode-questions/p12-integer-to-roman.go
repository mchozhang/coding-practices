/**
 * LeetCode: Integer to Roman
 * https://leetcode.com/problems/integer-to-roman/
 *
 * submission: faster than 100%
 */
package main

func intToRoman(num int) string {
	letters := map[int]string{
		1: "I", 4: "IV", 5: "V", 9: "IX", 10: "X",
		40: "XL", 50: "L", 90: "XC", 100: "C",
		400: "CD", 500: "D", 900: "CM", 1000: "M",
	}
	values := []int{1, 4, 5, 9, 10, 40, 50, 90, 100, 400, 500, 900, 1000, 4000, 5000, 9000, 10000}
	res := ""
	base := 1
	pos := 4
	for num > 0 {
		flag := num % 10 * base
		if flag == values[pos-1] {
			res = letters[values[pos-1]] + res
		} else if flag >= values[pos-2] {
			s := ""
			k := (flag - values[pos-2]) / values[pos-4]
			for i := 0; i < k; i++ {
				s += letters[values[pos-4]]
			}
			res = letters[values[pos-2]] + s + res
		} else if flag == values[pos-3] {
			res = letters[values[pos-3]] + res
		} else {
			s := ""
			k := flag / values[pos-4]
			for i:=0;i<k;i++{
				s += letters[values[pos-4]]
			}
			res = s + res
		}
		base *= 10
		pos += 4
		num /= 10
	}
	return res
}
