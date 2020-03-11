package leetcode

/*A majority element is an element that makes up more than half of the items in an array. Given a positive integers array, find the majority element. If there is no majority element, return -1. Do this in O(N) time and O(1) space.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/find-majority-element-lcci
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

//投票算法
func majorityElement(nums []int) int {
	if len(nums) == 0 {
		return -1
	}
	major := 0
	cnt := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != nums[major] {
			cnt--
		} else {
			cnt++
		}

		if cnt == 0 {
			major = i + 1
		}
	}

	if major == len(nums) {
		return -1
	}

	cnt = 0
	for i := 0; i < len(nums); i++ {
		if nums[major] == nums[i] {
			cnt++
		}
	}
	if cnt > len(nums)/2 {
		return nums[major]
	} else {
		return -1
	}
}
