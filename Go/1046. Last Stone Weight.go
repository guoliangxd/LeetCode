package leetcode

/*We have a collection of rocks, each rock has a positive integer weight.

Each turn, we choose the two heaviest rocks and smash them together.  Suppose the stones have weights x and y with x <= y.  The result of this smash is:

If x == y, both stones are totally destroyed;
If x != y, the stone of weight x is totally destroyed, and the stone of weight y has new weight y-x.
At the end, there is at most 1 stone left.  Return the weight of this stone (or 0 if there are no stones left.)

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/last-stone-weight
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

import "sort"

func lastStoneWeight(stones []int) int {
	size := len(stones)
	if size == 1 {
		return stones[0]
	}

	sort.Ints(stones)
	for stones[size-2] != 0 {
		stones[size-1] = stones[size-1] - stones[size-2]
		stones[size-2] = 0
		sort.Ints(stones)
	}
	return stones[size-1]
}
