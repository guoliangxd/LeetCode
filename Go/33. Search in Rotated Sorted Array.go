package leetcode

/*Suppose an array sorted in ascending order is rotated at some pivot unknown to you beforehand.

(i.e., [0,1,2,4,5,6,7] might become [4,5,6,7,0,1,2]).

You are given a target value to search. If found in the array return its index, otherwise return -1.

You may assume no duplicate exists in the array.

Your algorithm's runtime complexity must be in the order of O(log n).

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/search-in-rotated-sorted-array
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

func search(nums []int, target int) int {
	size := len(nums)
	if size == 0 {
		return -1
	}

	begin, end := 0, size-1
	for begin <= end {
		mid := begin + (end-begin)/2
		if nums[mid] == target {
			return mid
		}
		// 前半部分有序
		if nums[begin] <= nums[mid] {
			if target >= nums[begin] && target < nums[mid] {
				end = mid - 1
			} else {
				begin = mid + 1
			}
		} else {
			if target > nums[mid] && target <= nums[end] {
				begin = mid + 1
			} else {
				end = mid - 1
			}
		}
	}

	return -1
}
