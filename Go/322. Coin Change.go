package leetcode

/*You are given coins of different denominations and a total amount of money amount. Write a function to compute the fewest number of coins that you need to make up that amount. If that amount of money cannot be made up by any combination of the coins, return -1.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/coin-change
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/
func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1, amount+1)
	for i := 0; i < amount+1; i++ {
		dp[i] = amount + 1
	}
	//base case
	dp[0] = 0
	for i := 0; i < amount+1; i++ {
		for _, v := range coins {
			if i-v < 0 {
				continue
			}
			if dp[i] > dp[i-v]+1 {
				dp[i] = dp[i-v] + 1
			}
		}
	}
	if dp[amount] == amount+1 {
		return -1
	} else {
		return dp[amount]
	}
}
