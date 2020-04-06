package leetcode

/*Given two words word1 and word2, find the minimum number of operations required to convert word1 to word2.

You have the following 3 operations permitted on a word:

Insert a character
Delete a character
Replace a character

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/edit-distance
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

import "math"

func minDistance(word1 string, word2 string) int {
	n, m := len(word1), len(word2)

	// 有一个字符串为空
	if n*m == 0 {
		return n + m
	}

	// DP数组
	dp := make([][]int, n+1)

	// 边界状态初始化
	for i := 0; i < n+1; i++ {
		dp[i] = make([]int, m+1)
		dp[i][0] = i
	}
	for j := 0; j < m+1; j++ {
		dp[0][j] = j
	}

	// 计算所有DP值
	for i := 1; i < n+1; i++ {
		for j := 1; j < m+1; j++ {
			left := dp[i-1][j] + 1
			down := dp[i][j-1] + 1
			leftDown := dp[i-1][j-1]
			if word1[i-1] != word2[j-1] {
				leftDown++
			}
			dp[i][j] = int(math.Min(math.Min(float64(left), float64(down)), float64(leftDown)))
		}
	}
	return dp[n][m]
}
