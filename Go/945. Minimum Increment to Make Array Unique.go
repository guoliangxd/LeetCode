package leetcode

/*Given an array of integers A, a move consists of choosing any A[i], and incrementing it by 1.

Return the least number of moves to make every value in A unique.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/minimum-increment-to-make-array-unique
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

import "sort"

func minIncrementForUnique(A []int) int {
	res := 0
	cnt := make([]int, 80000)
	repeat := make([]int, 0)
	for i := 0; i < len(A); i++ {
		cnt[A[i]]++
		if cnt[A[i]] > 1 {
			repeat = append(repeat, A[i])
		}
	}
	sort.Ints(repeat)
	for i := 0; i < len(cnt) && len(repeat) > 0; i++ {
		if cnt[i] == 0 {
			if repeat[0] < i {
				res += i - repeat[0]
				repeat = repeat[1:]
			}
		}
	}
	return res
}
