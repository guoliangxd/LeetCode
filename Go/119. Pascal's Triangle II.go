package leetcode

/*Given a non-negative index k where k ≤ 33, return the kth index row of the Pascal's triangle.
Note that the row index starts from 0.
来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/pascals-triangle-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

func getRow(rowIndex int) []int {
	if rowIndex == 0 {
		return []int{1}
	} else {
		res := make([]int, rowIndex+1)
		last := getRow(rowIndex - 1)
		for i := 0; i < rowIndex+1; i++ {
			if i == 0 || i == rowIndex {
				res[i] = 1
			} else {
				res[i] = last[i-1] + last[i]
			}
		}
		return res
	}
}
