package leetcode

/*There are n soldiers standing in a line. Each soldier is assigned a unique rating value.

You have to form a team of 3 soldiers amongst them under the following rules:

Choose 3 soldiers with index (i, j, k) with rating (rating[i], rating[j], rating[k]).
A team is valid if:  (rating[i] < rating[j] < rating[k]) or (rating[i] > rating[j] > rating[k]) where (0 <= i < j < k < n).
Return the number of teams you can form given the conditions. (soldiers can be part of multiple teams).

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/count-number-of-teams
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

func numTeams(rating []int) int {
	size := len(rating)
	lBig := make([]int, size)
	lSmall := make([]int, size)
	rBig := make([]int, size)
	rSmall := make([]int, size)
	res := 0
	for i := 0; i < size; i++ {
		for j := 0; j < i; j++ {
			if rating[j] < rating[i] {
				lSmall[i]++
			} else if rating[j] > rating[i] {
				lBig[i]++
			}
		}
		for k := i + 1; k < size; k++ {
			if rating[k] < rating[i] {
				rSmall[i]++
			} else if rating[k] > rating[i] {
				rBig[i]++
			}
		}
	}

	for i := 0; i < size; i++ {
		res += lSmall[i] * rBig[i]
		res += lBig[i] * rSmall[i]
	}
	return res
}
