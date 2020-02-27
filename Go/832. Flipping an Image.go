package leetcode

/*Given a binary matrix A, we want to flip the image horizontally,
then invert it, and return the resulting image.

To flip an image horizontally means that each row of the image is reversed.
For example, flipping [1, 1, 0] horizontally results in [0, 1, 1].

To invert an image means that each 0 is replaced by 1, and each 1 is replaced by 0.
For example, inverting [0, 1, 1] results in [1, 0, 0].

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/flipping-an-image
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

func flipAndInvertImage(A [][]int) [][]int {
	row := len(A)
	col := len(A[0])
	for i := 0; i < row; i++ {
		for j := 0; j < col/2; j++ {
			A[i][j], A[i][col-j-1] = 1-A[i][col-j-1], 1-A[i][j]
		}
		if col%2 == 1 {
			A[i][col/2] = 1 - A[i][col/2]
		}
	}
	return A
}
