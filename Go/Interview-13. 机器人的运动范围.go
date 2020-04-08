package leetcode

/*地上有一个m行n列的方格，从坐标 [0,0] 到坐标 [m-1,n-1] 。一个机器人从坐标 [0, 0] 的格子开始移动，它每次可以向左、右、上、下移动一格（不能移动到方格外），也不能进入行坐标和列坐标的数位之和大于k的格子。例如，当k为18时，机器人能够进入方格 [35, 37] ，因为3+5+3+7=18。但它不能进入方格 [35, 38]，因为3+5+3+8=19。请问该机器人能够到达多少个格子？

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/ji-qi-ren-de-yun-dong-fan-wei-lcof
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

func movingCount(m int, n int, k int) int {
	cnt := 0
	martix := make([][]int, m)
	for i := 0; i < m; i++ {
		martix[i] = make([]int, n)
	}
	martix[0][0] = 1
	cnt++
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if martix[i][j] == 1 {
				if i-1 >= 0 && martix[i-1][j] == 0 && sum(i-1, j) <= k {
					martix[i-1][j] = 1
					cnt++
				}
				if i+1 < m && martix[i+1][j] == 0 && sum(i+1, j) <= k {
					martix[i+1][j] = 1
					cnt++
				}
				if j-1 >= 0 && martix[i][j-1] == 0 && sum(i, j-1) <= k {
					martix[i][j-1] = 1
					cnt++
				}
				if j+1 < n && martix[i][j+1] == 0 && sum(i, j+1) <= k {
					martix[i][j+1] = 1
					cnt++
				}
			}
		}
	}

	return cnt
}

func sum(x, y int) int {
	res := 0
	for x != 0 {
		res += x % 10
		x /= 10
	}
	for y != 0 {
		res += y % 10
		y /= 10
	}
	return res
}
