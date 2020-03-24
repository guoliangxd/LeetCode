package leetcode

/*Given a non-empty integer array of size n, find the minimum number of moves required to make all array elements equal, where a move is incrementing n - 1 elements by 1.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/minimum-moves-to-equal-array-elements
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

import "sort"

func minMoves(nums []int) int {
	sort.Ints(nums)
	res := 0
	size := len(nums)
	for i := 0; i < size; i++ {
		res += nums[i] - nums[0]
	}
	return res
}
