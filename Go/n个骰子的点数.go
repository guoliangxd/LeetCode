package leetcode

/*把n个骰子扔在地上，所有骰子朝上一面的点数之和为s。输入n，打印出s的所有可能的值出现的概率。



你需要用一个浮点数数组返回答案，其中第 i 个元素代表这 n 个骰子所能掷出的点数集合中第 i 小的那个的概率。



来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/nge-tou-zi-de-dian-shu-lcof
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

func twoSum(n int) []float64 {
	size := 6 + 5*(n-1)
	res := make([]int, n*6)
	base := 1.0
	for i := 0; i < n; i++ {
		base *= 6
	}
	rlt := make([]float64, size)

	for i := 0; i < 6; i++ {
		res[i] = 1
	}

	for i := 1; i < n; i++ {
		for j := 5 + i*6; j >= i; j-- {
			res[j] = 0
			for k := j - 1; k >= j-6 && k >= 0; k-- {
				res[j] += res[k]
			}
		}

		res[i-1] = 0
	}

	max := 6*n - 1
	for i := 0; i < size; i++ {
		rlt[i] = float64(res[max-i]) / base
	}

	return rlt
}
