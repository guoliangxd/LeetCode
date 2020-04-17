package leetcode

/*Given an array of non-negative integers, you are initially positioned at the first index of the array.

Each element in the array represents your maximum jump length at that position.

Determine if you are able to reach the last index.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/jump-game
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

func canJump(nums []int) bool {
	size := len(nums)
	max := 0

	for i := 0; i < size && i <= max; i++ {
		temp := nums[i] + i
		if temp > max {
			max = temp
		}
		if max >= size-1 {
			return true
		}
	}
	return false
}
