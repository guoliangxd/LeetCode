package leetcode

/*Say you have an array for which the ith element is the price of a given stock on day i.

If you were only permitted to complete at most one transaction (i.e., buy one and sell one share of the stock), design an algorithm to find the maximum profit.

Note that you cannot sell a stock before you buy one.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

func maxProfit(prices []int) int {
	length := len(prices)
	rlt := 0
	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			if prices[j]-prices[i] > rlt {
				rlt = prices[j] - prices[i]
			}
		}
	}
	return rlt
}

/* 单次遍历
func maxProfit(prices []int) int {
    rlt := 0
    min := 0xefffffff

    for i := 0; i < len(prices); i++ {
        if min > prices[i] {
            min = prices[i]
        } else if prices[i] - min > rlt {
            rlt = prices[i]  - min
        }
    }

    return rlt
}*/
