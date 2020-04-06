package leetcode

/*Given a positive integer n, generate a square matrix filled with elements from 1 to n2 in spiral order.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/spiral-matrix-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

func generateMatrix(n int) [][]int {
	res := make([][]int, n)
	for i := 0; i < n; i++ {
		res[i] = make([]int, n)
	}
	l, r, t, b := 0, n-1, 0, n-1
	num, tar := 1, n*n

	for num <= tar {
		for i := l; i <= r; i++ {
			res[t][i] = num
			num++
		}
		t++
		for i := t; i <= b; i++ {
			res[i][r] = num
			num++
		}
		r--
		for i := r; i >= l; i-- {
			res[b][i] = num
			num++
		}
		b--
		for i := b; i >= t; i-- {
			res[i][l] = num
			num++
		}
		l++
	}
	return res
}
