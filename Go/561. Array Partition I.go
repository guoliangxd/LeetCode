package leetcode

/*Given an array of 2n integers, your task is to group these integers into n pairs of integer, say (a1, b1), (a2, b2), ..., (an, bn) which makes sum of min(ai, bi) for all i from 1 to n as large as possible.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/array-partition-i
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

import "sort"

func arrayPairSum(nums []int) int {
	sort.Ints(nums)
	length := len(nums)
	rlt := 0
	for i := 0; i < length/2; i++ {
		rlt += nums[2*i]
	}
	return rlt
}
