/**
 * LeetCode : 1431. Kids With the Greatest Number of Candies
 * https://leetcode.com/problems/kids-with-the-greatest-number-of-candies/
 *
 * submission : faster than 100%
 */

package main

import "math"

func kidsWithCandies(candies []int, extraCandies int) []bool {
	max := math.MinInt32
	for _, v := range candies {
		if v > max {
			max = v
		}
	}

	max -= extraCandies

	res := make([]bool, len(candies))
	for i := 0; i < len(res); i++ {
		res[i] = candies[i] >= max
	}
	return res
}
