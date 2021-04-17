/**
 * LeetCode : Relative Ranks
 * https://leetcode.com/problems/relative-ranks/
 *
 * submission : faster than 56%
 */
package main

import (
	"sort"
	"strconv"
)

func findRelativeRanks(score []int) []string {
	index := map[int]int{}
	for i, num := range score {
		index[num] = i
	}
	sort.Ints(score)
	res := make([]string, len(score))
	for i := 3; i < len(score); i++ {
		res[index[score[i]]] = strconv.Itoa(len(score) - i)
	}
	if len(score) > 1 {
		res[index[score[len(score)-1]]] = "Gold Medal"
	}
	if len(score) > 2 {
		res[index[score[len(score)-2]]] = "Silver Medal"
	}
	if len(score) > 3 {
		res[index[score[len(score)-2]]] = "Bronze Medal"
	}
	return res
}
