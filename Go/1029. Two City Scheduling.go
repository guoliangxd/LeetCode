package leetcode

/*There are 2N people a company is planning to interview. The cost of flying the i-th person to city A is costs[i][0], and the cost of flying the i-th person to city B is costs[i][1].

Return the minimum cost to fly every person to a city such that exactly N people arrive in each city.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/two-city-scheduling
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

import "sort"

func twoCitySchedCost(costs [][]int) int {
	BtoA := make([]int, len(costs))
	res := 0
	for i := 0; i < len(costs); i++ {
		BtoA[i] = costs[i][0] - costs[i][1]
		res += costs[i][1]
	}
	sort.Ints(BtoA)
	for i := 0; i < len(costs)/2; i++ {
		res += BtoA[i]
	}
	return res
}
