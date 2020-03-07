package leetcode

/*Given a matrix A, return the transpose of A.

The transpose of a matrix is the matrix flipped over it's main diagonal, switching the row and column indices of the matrix.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/transpose-matrix
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

func transpose(A [][]int) [][]int {
	m, n := len(A), len(A[0])
	rlt := make([][]int, n, n)
	for i := 0; i < n; i++ {
		temp := make([]int, m, m)
		for j := 0; j < m; j++ {
			temp[j] = A[j][i]
		}
		rlt[i] = temp
	}
	return rlt
}
