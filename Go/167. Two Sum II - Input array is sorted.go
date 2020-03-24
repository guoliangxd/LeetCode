package leetcode

/*Given an array of integers that is already sorted in ascending order, find two numbers such that they add up to a specific target number.

The function twoSum should return indices of the two numbers such that they add up to the target, where index1 must be less than index2.

Note:

Your returned answers (both index1 and index2) are not zero-based.
You may assume that each input would have exactly one solution and you may not use the same element twice.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/two-sum-ii-input-array-is-sorted
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

func twoSum(numbers []int, target int) []int {
	res := make([]int, 2)
	size := len(numbers)
	for i := 0; i < size-1; i++ {
		find := binarySearch(numbers, i+1, size-1, target-numbers[i])
		if find != -1 {
			res[0] = i + 1
			res[1] = find + 1
			break
		}
	}
	return res
}

func binarySearch(arr []int, begin int, end int, target int) int {
	for begin <= end {
		mid := begin + (end-begin)/2
		if arr[mid] == target {
			return mid
		} else if arr[mid] > target {
			end = mid - 1
		} else {
			begin = mid + 1
		}
	}
	return -1
}
