package leetcode

/*数组中有一个数字出现的次数超过数组长度的一半，请找出这个数字。
你可以假设数组是非空的，并且给定的数组总是存在多数元素。*/

func majorityElement(nums []int) int {
	mapCnt := make(map[int]int, 0)
	length := len(nums)

	for i := 0; i < length; i++ {
		v, ok := mapCnt[nums[i]]
		if ok {
			mapCnt[nums[i]] = v + 1
		} else {
			mapCnt[nums[i]] = 1
		}

		if mapCnt[nums[i]] == (length+1)/2 {
			return nums[i]
		}
	}
	return -1
}
