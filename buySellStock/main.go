/*
Best Time to Buy and Sell Stock
You are given an array prices where prices[i] is the
price of a given stock on the ith day.

You want to maximize your profit by choosing a single
day to buy one stock and choosing a different day in the
future to sell that stock.

Return the maximum profit you can achieve from this transaction.
If you cannot achieve any profit, return 0.
*/

package main

import "fmt"

func main() {

	prices := []int{7, 1, 5, 3, 6, 4}

	maxProfit := maxProfit(prices)

	fmt.Println("The Maximum profit is ", maxProfit)

}

func maxProfit(prices []int) int {

	minPurchase := prices[0]
	maxProfit := 0
	if len(prices) <= 1 {
		return 0
	}

	for i := 1; i < len(prices); i++ {

		if prices[i] < minPurchase {
			minPurchase = prices[i]
		}
		if prices[i]-minPurchase > maxProfit {

			maxProfit = prices[i] - minPurchase
		}

	}
	return maxProfit
}
