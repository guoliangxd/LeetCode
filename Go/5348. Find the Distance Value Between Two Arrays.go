package leetcode

/*Given two integer arrays arr1 and arr2, and the integer d, return the distance value between the two arrays.

The distance value is defined as the number of elements arr1[i] such that there is not any element arr2[j] where |arr1[i]-arr2[j]| <= d.*/

import "sort"

func findTheDistanceValue(arr1 []int, arr2 []int, d int) int {
	sort.Ints(arr2)
	cnt := 0
	for i := 0; i < len(arr1); i++ {
		if !binarySearch(arr1[i], d, arr2) {
			cnt++
		}
	}
	return cnt
}

func binarySearch(target int, offset int, arr []int) bool {
	begin := 0
	end := len(arr) - 1
	for begin <= end {
		mid := begin + (end-begin)/2
		if arr[mid] >= target-offset && arr[mid] <= target+offset {
			return true
		} else if arr[mid] < target-offset {
			begin = mid + 1
		} else {
			end = mid - 1
		}
	}
	return false
}
