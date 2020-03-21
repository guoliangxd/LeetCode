package leetcode

/*Given scores of N athletes, find their relative ranks and the people with the top three highest scores, who will be awarded medals: "Gold Medal", "Silver Medal" and "Bronze Medal".

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/relative-ranks
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

import (
	"sort"
	"strconv"
)

func findRelativeRanks(nums []int) []string {
	sorts := make([]int, len(nums))
	copy(sorts, nums)
	sort.Ints(sorts)
	rank := make(map[int]int, len(sorts))
	for i := 0; i < len(sorts); i++ {
		rank[sorts[i]] = i + 1
	}
	res := make([]string, len(nums))
	for i := 0; i < len(nums); i++ {
		ran := rank[nums[i]]
		ran = len(nums) - ran + 1
		switch ran {
		case 1:
			res[i] = "Gold Medal"
		case 2:
			res[i] = "Silver Medal"
		case 3:
			res[i] = "Bronze Medal"
		default:
			res[i] = strconv.Itoa(ran)
		}
	}
	return res
}
