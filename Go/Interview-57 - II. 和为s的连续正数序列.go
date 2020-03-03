package leetcode

/*输入一个正整数 target ，输出所有和为 target 的连续正整数序列（至少含有两个数）。

序列内的数字由小到大排列，不同序列按照首个数字从小到大排列。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/he-wei-sde-lian-xu-zheng-shu-xu-lie-lcof
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

//单向双指针
func findContinuousSequence(target int) [][]int {
	l, r := 1, 2
	rlt := make([][]int, 0, 0)
	for l < r {
		temp := sum(l, r)
		if temp > target {
			l++
		} else if temp < target {
			r++
		} else {
			data := make([]int, r-l+1)
			for i := 0; i < r-l+1; i++ {
				data[i] = l + i
			}
			rlt = append(rlt, data)
			l++
			r = l + 1
		}
	}
	return rlt
}

func sum(l, r int) int {
	return (l + r) * (r - l + 1) / 2
}
