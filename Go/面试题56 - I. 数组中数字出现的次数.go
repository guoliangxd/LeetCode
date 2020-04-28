package leetcode

/*一个整型数组 nums 里除两个数字之外，其他数字都出现了两次。请写程序找出这两个只出现一次的数字。要求时间复杂度是O(n)，空间复杂度是O(1)。*/

func singleNumbers(nums []int) []int {
	size := len(nums)
	xor := 0

	for i := 0; i < size; i++ {
		xor ^= nums[i]
	}

	bit := 0
	for i := 0; ; i++ {
		if xor&(0x1<<i) != 0 {
			bit = i
			break
		}
	}

	a, b := 0, 0
	for i := 0; i < size; i++ {
		if nums[i]&(0x1<<bit) != 0 {
			a ^= nums[i]
		} else {
			b ^= nums[i]
		}
	}

	return []int{a, b}
}
