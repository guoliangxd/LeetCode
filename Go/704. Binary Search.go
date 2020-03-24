package leetcode

/*Given a sorted (in ascending order) integer array nums of n elements and a target value, write a function to search target in nums. If target exists, then return its index, otherwise return -1.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/binary-search
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

func search(nums []int, target int) int {
	begin := 0
	end := len(nums) - 1
	for begin <= end {
		mid := begin + (end-begin)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			end = mid - 1
		} else {
			begin = mid + 1
		}
	}
	return -1
}
