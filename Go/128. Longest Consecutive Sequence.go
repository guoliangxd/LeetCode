package leetcode

/*Given an unsorted array of integers, find the length of the longest consecutive elements sequence.

Your algorithm should run in O(n) complexity.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/longest-consecutive-sequence
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

func longestConsecutive(nums []int) int {
	size := len(nums)
	set := make(map[int]int, size)
	for i := 0; i < size; i++ {
		set[nums[i]] = 1
	}

	res := 0

	for i := 0; i < size; i++ {
		_, ok := set[nums[i]-1]
		if !ok {
			//只从序列第一个元素开始查找
			temp := 1
			cur := nums[i]
			for {
				cur++
				_, find := set[cur]
				if find {
					temp++
				} else {
					break
				}
			}
			if temp > res {
				res = temp
			}
		}
	}
	return res
}
