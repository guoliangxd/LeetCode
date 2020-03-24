package leetcode

/*A popular masseuse receives a sequence of back-to-back appointment requests and is debating which ones to accept. She needs a break between appointments and therefore she cannot accept any adjacent requests. Given a sequence of back-to-back appoint­ ment requests, find the optimal (highest total booked minutes) set the masseuse can honor. Return the number of minutes.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/the-masseuse-lcci
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

func massage(nums []int) int {
	size := len(nums)
	if size == 0 {
		return 0
	} else if size == 1 {
		return nums[0]
	}
	dp := make([]int, size)
	dp[0] = nums[0]
	for i := 1; i < size; i++ {
		dp[i] = nums[i]
		if i-2 >= 0 {
			if i-3 >= 0 && dp[i-3] > dp[i-2] {
				dp[i] += dp[i-3]
			} else {
				dp[i] += dp[i-2]
			}
		}
	}

	if dp[size-1] >= dp[size-2] {
		return dp[size-1]
	} else {
		return dp[size-2]
	}
}
