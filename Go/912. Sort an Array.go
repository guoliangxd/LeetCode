package leetcode

/*Given an array of integers nums, sort the array in ascending order.*/

import "sort"

func sortArray(nums []int) []int {
	sort.Ints(nums)
	return nums
}
