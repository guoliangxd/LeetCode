package leetcode

/*Given the array nums, obtain a subsequence of the array whose sum of elements is strictly greater than the sum of the non included elements in such subsequence.

If there are multiple solutions, return the subsequence with minimum size and if there still exist multiple solutions, return the subsequence with the maximum total sum of all its elements. A subsequence of an array can be obtained by erasing some (possibly zero) elements from the array.

Note that the solution with the given constraints is guaranteed to be unique. Also return the answer sorted in non-increasing order.*/

import "sort"

func minSubsequence(nums []int) []int {
	sort.Ints(nums)
	sum := 0
	index := 0

	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}

	subSum := 0
	for i := len(nums) - 1; i >= 0; i-- {
		subSum += nums[i]
		if subSum > sum-subSum {
			index = i
			break
		}
	}

	res := make([]int, len(nums)-index)
	for i, j := len(nums)-1, 0; i >= index; i-- {
		res[j] = nums[i]
		j++
	}

	return res
}
