package leetcode

/*输入一个递增排序的数组和一个数字s，在数组中查找两个数，使得它们的和正好是s。
如果有多对数字的和等于s，则输出任意一对即可。*/

func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		if binarySearch(nums, i, len(nums)-1, target-nums[i]) {
			return []int{nums[i], target - nums[i]}
		}
	}
	return []int{0, 0}
}

func binarySearch(nums []int, begin int, end int, target int) bool {
	rlt := false
	if target < nums[begin] || target > nums[end] {
		return rlt
	}
	for begin <= end {
		mid := (begin + end) / 2
		if nums[mid] == target {
			rlt = true
			break
		} else if nums[mid] < target {
			begin = mid + 1
		} else {
			end = mid - 1
		}
	}
	return rlt
}
