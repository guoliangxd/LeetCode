package leetcode

/*Say you have an array for which the ith element is the price of a given stock on day i.

Design an algorithm to find the maximum profit. You may complete as many transactions as you like (i.e., buy one and sell one share of the stock multiple times).

Note: You may not engage in multiple transactions at the same time (i.e., you must sell the stock before you buy again).*/

func maxProfit(prices []int) int {
	profit := 0
	isIncrease := false
	buy := 0
	for i := 0; i < len(prices)-1; i++ {
		if isIncrease {
			if prices[i] > prices[i+1] {
				profit += prices[i] - buy
				isIncrease = false
			}
		} else {
			if prices[i] < prices[i+1] {
				buy = prices[i]
				isIncrease = true
			}
		}
	}

	if isIncrease {
		profit += prices[len(prices)-1] - buy
	}

	return profit
}
