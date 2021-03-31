/**
 * LeetCode : Best Time to Buy and Sell Stock II
 * https://leetcode.com/problems/best-time-to-buy-and-sell-stock-ii/
 *
 * submission : faster than 94%
 */
package main

func maxProfit(prices []int) int {
	sum := 0
	for i := 1; i < len(prices);i++ {
		diff := prices[i] - prices[i-1]
		if diff > 0 {
			sum += diff
		}
	}
	return sum
}