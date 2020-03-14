package leetcode

/*Given an unsorted array of integers, find the length of longest increasing subsequence.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/longest-increasing-subsequence
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	dp := make([]int, len(nums))
	res := dp[0]
	for i := 0; i < len(nums); i++ {
		last := 0
		for j := 0; j < i; j++ {
			if last < dp[j] && nums[j] < nums[i] {
				last = dp[j]
			}
		}
		dp[i] = last + 1
		if dp[i] > res {
			res = dp[i]
		}
	}
	return res
}
