package leetcode

/*
Given an array nums, write a function to move all 0's to the end of it while maintaining the relative order of the non-zero elements.*/

func moveZeroes(nums []int) {
	index := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			if i != index {
				nums[i], nums[index] = nums[index], nums[i]
			}
			index++
		}
	}
}
