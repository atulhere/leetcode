/*
Best Time to Buy and Sell Stock II
You are given an integer array prices where prices[i]
is the price of a given stock on the ith day.

On each day, you may decide to buy and/or sell the stock.
You can only hold at most one share of the stock at any time.
However, you can buy it then immediately sell it on the same day.

Find and return the maximum profit you can achieve.
*/

package main

import "fmt"

func main() {

	prices := []int{7, 1, 5, 3, 6, 4}

	maxProfit := maxProfit(prices)

	fmt.Println("Maximum Profit is ", maxProfit)
}

func maxProfit(prices []int) int {

	maxProfit := 0

	if len(prices) <= 1 {
		return 0

	}

	for i := 1; i < len(prices); i++ {

		if prices[i] > prices[i-1] {

			maxProfit = maxProfit + (prices[i] - prices[i-1])
		}
	}

	return maxProfit

}
