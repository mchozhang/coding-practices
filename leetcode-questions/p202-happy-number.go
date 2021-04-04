/**
 * LeetCode: Happy Number
 * https://leetcode.com/problems/happy-number/
 *
 * submission : faster than 100%
 */
package main

func isHappy(n int) bool {
	res := map[int]int{}
	num := n
	for true {
		sum := 0
		for i := num; i > 0; i /= 10 {
			sum += (i % 10) * (i % 10)
		}
		if sum == 1 {
			return true
		} else if _, isIn := res[sum]; isIn {
			return false
		} else {
			res[sum] = 0
			num = sum
		}
	}
	return false
}
