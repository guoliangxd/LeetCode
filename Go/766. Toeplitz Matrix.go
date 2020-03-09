package leetcode

/*A matrix is Toeplitz if every diagonal from top-left to bottom-right has the same element.

Now given an M x N matrix, return True if and only if the matrix is Toeplitz.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/toeplitz-matrix
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

func isToeplitzMatrix(matrix [][]int) bool {
	m, n := len(matrix), len(matrix[0])
	for j := 0; j < n; j++ {
		val := matrix[0][j]
		for a, b := 0, j; a < m && b < n; a, b = a+1, b+1 {
			if matrix[a][b] != val {
				return false
			}
		}
	}

	for i := 1; i < m; i++ {
		val := matrix[i][0]
		for a, b := i, 0; a < m && b < n; a, b = a+1, b+1 {
			if matrix[a][b] != val {
				return false
			}
		}
	}

	return true
}
