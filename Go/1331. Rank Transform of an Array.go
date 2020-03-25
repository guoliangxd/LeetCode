package leetcode

/*Given an array of integers arr, replace each element with its rank.

The rank represents how large the element is. The rank has the following rules:

Rank is an integer starting from 1.
The larger the element, the larger the rank. If two elements are equal, their rank must be the same.
Rank should be as small as possible.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/rank-transform-of-an-array
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

import "sort"

func arrayRankTransform(arr []int) []int {
	size := len(arr)
	res := make([]int, size)
	if size == 0 {
		return res
	}

	curRank := 1
	sorted := make([]int, size)
	//target source
	copy(sorted, arr)
	sort.Ints(sorted)
	rankMap := make(map[int]int, 0)
	for i := 0; i < size; i++ {
		_, ok := rankMap[sorted[i]]
		if !ok {
			rankMap[sorted[i]] = curRank
			curRank++
		}
	}
	for i := 0; i < size; i++ {
		res[i] = rankMap[arr[i]]
	}
	return res
}
