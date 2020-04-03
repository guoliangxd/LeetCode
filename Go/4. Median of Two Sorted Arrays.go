package leetcode

/*There are two sorted arrays nums1 and nums2 of size m and n respectively.

Find the median of the two sorted arrays. The overall run time complexity should be O(log (m+n)).

You may assume nums1 and nums2 cannot be both empty.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/median-of-two-sorted-arrays
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	size1 := len(nums1)
	size2 := len(nums2)
	size := size1 + size2

	newArray := make([]int, size)
	index := 0
	i, j := 0, 0
	for ; i < size1 && j < size2; index++ {
		if nums1[i] <= nums2[j] {
			newArray[index] = nums1[i]
			i++
		} else {
			newArray[index] = nums2[j]
			j++
		}
	}

	if i < size1 {
		copy(newArray[index:], nums1[i:])
	}

	if j < size2 {
		copy(newArray[index:], nums2[j:])
	}

	if size%2 == 1 {
		return float64(newArray[size/2])
	} else {
		return (float64(newArray[size/2-1]) + float64(newArray[size/2])) / 2.0
	}
}
