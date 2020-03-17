package leetcode

/*You are given an array of integers (both positive and negative). Find the contiguous sequence with the largest sum. Return the sum.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/contiguous-sequence-lcci
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

func maxSubArray(nums []int) int {
	size := len(nums)
	if size == 0 {
		return 0
	}
	dp := make([]int, size)
	dp[0] = nums[0]
	res := dp[0]
	for i := 1; i < size; i++ {
		if dp[i-1] > 0 {
			dp[i] = dp[i-1]
		}
		dp[i] += nums[i]
		if dp[i] > res {
			res = dp[i]
		}
	}
	return res
}
