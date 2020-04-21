package leetcode

/*Given an array of integers nums and an integer k. A subarray is called nice if there are k odd numbers on it.

Return the number of nice sub-arrays.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/count-number-of-nice-subarrays
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

func numberOfSubarrays(nums []int, k int) int {
	odd, res := 0, 0
	size := len(nums)

	for i := 0; i < size; i++ {
		if nums[i]%2 == 0 {
			nums[i] = 0
		} else {
			nums[i] = 1
			odd++
		}
	}

	if odd < k {
		return res
	}

	left := 0
	for nums[left] == 0 {
		left++
	}

	cnt := 1
	right := left
	for cnt != k {
		right++
		if nums[right] == 1 {
			cnt++
		}
	}

	//已经找到第一个满足条件的子数组
	for cnt == k {
		//尝试左右扩展
		l, r := left, right
		for l-1 >= 0 && nums[l-1] == 0 {
			l--
		}
		for r+1 < size && nums[r+1] == 0 {
			r++
		}
		res += (left - l + 1) * (r - right + 1)

		left++
		for left < size && nums[left] == 0 {
			left++
		}
		cnt--
		for cnt != k && right < (size-1) {
			right++
			if nums[right] == 1 {
				cnt++
			}
		}

	}
	return res
}
