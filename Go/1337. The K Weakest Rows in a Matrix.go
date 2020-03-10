package leetcode

/*Given a m * n matrix mat of ones (representing soldiers) and zeros (representing civilians), return the indexes of the k weakest rows in the matrix ordered from the weakest to the strongest.

A row i is weaker than row j, if the number of soldiers in row i is less than the number of soldiers in row j, or they have the same number of soldiers but i is less than j. Soldiers are always stand in the frontier of a row, that is, always ones may appear first and then zeros.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/the-k-weakest-rows-in-a-matrix
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

import "sort"

func kWeakestRows(mat [][]int, k int) []int {
	m, n := len(mat), len(mat[0])
	res := make([]int, m)
	ans := make([]int, m)
	for i := 0; i < m; i++ {
		ans[i] = i
		for j := 0; j < n; j++ {
			if mat[i][j] == 1 {
				res[i]++
			} else {
				break
			}
		}
	}
	sort.Slice(ans, func(i, j int) bool {
		return res[ans[i]] < res[ans[j]] || (res[ans[i]] == res[ans[j]] && ans[i] < ans[j])
	})
	return ans[:k]
}
