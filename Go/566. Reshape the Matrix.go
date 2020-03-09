package leetcode

/*In MATLAB, there is a very useful function called 'reshape', which can reshape a matrix into a new one with different size but keep its original data.

You're given a matrix represented by a two-dimensional array, and two positive integers r and c representing the row number and column number of the wanted reshaped matrix, respectively.

The reshaped matrix need to be filled with all the elements of the original matrix in the same row-traversing order as they were.

If the 'reshape' operation with given parameters is possible and legal, output the new reshaped matrix; Otherwise, output the original matrix.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/reshape-the-matrix
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

func matrixReshape(nums [][]int, r int, c int) [][]int {
	m, n := len(nums), len(nums[0])
	if m*n != r*c || (m == r && n == c) {
		return nums
	}

	rlt := make([][]int, r)
	index := 0
	for i := 0; i < r; i++ {
		rlt[i] = make([]int, c)
		for j := 0; j < c; j++ {
			rlt[i][j] = nums[index/n][index%n]
			index++
		}
	}
	return rlt
}
