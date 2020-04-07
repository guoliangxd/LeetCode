package leetcode

/*Given an image represented by an N x N matrix, where each pixel in the image is 4 bytes, write a method to rotate the image by 90 degrees. Can you do this in place?

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/rotate-matrix-lcci
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

func rotate(matrix [][]int) {
	n := len(matrix)

	for i := 0; i < n/2; i++ {
		for j := i; j < n-i-1; j++ {
			r, c := j, n-i-1
			temp := matrix[r][c]
			matrix[r][c] = matrix[i][j]
			r, c = n-i-1, n-j-1
			temp, matrix[r][c] = matrix[r][c], temp
			r, c = n-j-1, i
			temp, matrix[r][c] = matrix[r][c], temp
			matrix[i][j] = temp
		}
	}
}
