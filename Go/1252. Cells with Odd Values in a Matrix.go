package leetcode

/*Given n and m which are the dimensions of a matrix,
which initialized by zeros
and given an array indices where indices[i] = [ri, ci].
For each pair of [ri, ci] you have to increment all cells in row ri and column ci by 1.
Return the number of cells with odd values in the matrix after applying the increment to all indices.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/cells-with-odd-values-in-a-matrix
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

func oddCells(n int, m int, indices [][]int) int {
	arr := make([][]bool, n, n)
	for i := 0; i < n; i++ {
		arr[i] = make([]bool, m, m)
	}
	length := len(indices)
	for i := 0; i < length; i++ {
		for j := 0; j < m; j++ {
			arr[indices[i][0]][j] = !arr[indices[i][0]][j]
		}

		for j := 0; j < n; j++ {
			arr[j][indices[i][1]] = !arr[j][indices[i][1]]
		}
	}

	rlt := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if arr[i][j] {
				rlt++
			}
		}
	}

	return rlt
}
