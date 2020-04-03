package leetcode

/*Given an array nums with n integers, your task is to check if it could become non-decreasing by modifying at most 1 element.

We define an array is non-decreasing if nums[i] <= nums[i + 1] holds for every i (0-based) such that (0 <= i <= n - 2).

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/non-decreasing-array
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

func checkPossibility(nums []int) bool {
	change := false
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			if change {
				return false
			} else {
				change = true
				if i+2 < len(nums) && nums[i] <= nums[i+2] {
					nums[i+1] = nums[i]
				} else if i+2 == len(nums) {
					nums[i+1] = nums[i]
				} else {
					nums[i] = nums[i+1]
					if i-1 >= 0 && nums[i-1] > nums[i] {
						return false
					}
				}
			}
		}
	}
	return true
}
