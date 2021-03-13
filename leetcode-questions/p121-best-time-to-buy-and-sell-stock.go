/**
 * LeetCode : Best Time to Buy and Sell Stock
 * https://leetcode.com/problems/best-time-to-buy-and-sell-stock/
 *
 * submission : faster than 65%
 */
package main

func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	max := prices[1] - prices[0]
	minIndex := 0
	for i := 2; i < len(prices); i++ {
		if prices[i-1] < prices[minIndex] {
			minIndex = i - 1
		}
		if prices[i] - prices[minIndex] > max {
			max = prices[i] - prices[minIndex]
		}
	}
	if max > 0 {
		return max
	}
	return 0
}
