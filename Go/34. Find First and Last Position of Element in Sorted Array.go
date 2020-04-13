package leetcode

/*Given an array of integers nums sorted in ascending order, find the starting and ending position of a given target value.

Your algorithm's runtime complexity must be in the order of O(log n).

If the target is not found in the array, return [-1, -1].

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/find-first-and-last-position-of-element-in-sorted-array
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

func searchRange(nums []int, target int) []int {
	lo, hi := 0, len(nums)-1
	for lo <= hi {
		mid := lo + (hi-lo)/2
		if nums[mid] == target {
			//新的查找策略
			//在lo和mid之间查找小于target的最大值
			//在mid和hi之间查找大于target的最小值
			left, right := mid, mid
			for lo < left {
				lmid := lo + (left-lo)/2
				if nums[lmid] == target {
					left = lmid
				} else {
					lo = lmid + 1
				}
			}

			for right < hi {
				rmid := right + (hi-right+1)/2
				if nums[rmid] == target {
					right = rmid
				} else {
					hi = rmid - 1
				}
			}
			return []int{left, right}
		} else if nums[mid] > target {
			hi = mid - 1
		} else {
			lo = mid + 1
		}
	}
	return []int{-1, -1}
}
