package leetcode

/*Given a 2D integer matrix M representing the gray scale of an image, you need to design a smoother to make the gray scale of each cell becomes the average gray scale (rounding down) of all the 8 surrounding cells and itself. If a cell has less than 8 surrounding cells, then use as many as you can.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/image-smoother
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

func imageSmoother(M [][]int) [][]int {
	m, n := len(M), len(M[0])
	res := make([][]int, m)
	for i := 0; i < m; i++ {
		res[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			cnt := 1
			sum := M[i][j]
			if j-1 >= 0 {
				cnt++
				sum += M[i][j-1]
			}
			if j+1 < n {
				cnt++
				sum += M[i][j+1]
			}
			if i-1 >= 0 {
				cnt++
				sum += M[i-1][j]
				if j-1 >= 0 {
					cnt++
					sum += M[i-1][j-1]
				}
				if j+1 < n {
					cnt++
					sum += M[i-1][j+1]
				}
			}
			if i+1 < m {
				cnt++
				sum += M[i+1][j]
				if j-1 >= 0 {
					cnt++
					sum += M[i+1][j-1]
				}
				if j+1 < n {
					cnt++
					sum += M[i+1][j+1]
				}
			}
			res[i][j] = sum / cnt
		}
	}

	return res
}
