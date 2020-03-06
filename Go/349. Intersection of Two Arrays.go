package leetcode

/*Given two arrays, write a function to compute their intersection.*/

func intersection(nums1 []int, nums2 []int) []int {
	rlt := make([]int, 0)
	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}

	mapNum1 := make(map[int]int, len(nums1))
	for i := 0; i < len(nums1); i++ {
		_, ok := mapNum1[nums1[i]]
		if !ok {
			mapNum1[nums1[i]] = 1
		}
	}

	for i := 0; i < len(nums2); i++ {
		_, ok := mapNum1[nums2[i]]
		if ok {
			rlt = append(rlt, nums2[i])
			delete(mapNum1, nums2[i])
		}
	}
	return rlt
}
