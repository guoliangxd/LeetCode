package leetcode

/*Given two arrays arr1 and arr2, the elements of arr2 are distinct,
and all elements in arr2 are also in arr1.
Sort the elements of arr1 such that the relative ordering of items in arr1 are the same as in arr2.
Elements that don't appear in arr2 should be placed at the end of arr1 in ascending order.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/relative-sort-array
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

func relativeSortArray(arr1 []int, arr2 []int) []int {
	larr1 := len(arr1)
	larr2 := len(arr2)
	cnt := [1001]int{}
	rlt := make([]int, 0, larr1)

	for i := 0; i < larr1; i++ {
		cnt[arr1[i]]++
	}
	for i := 0; i < larr2; i++ {
		for cnt[arr2[i]] != 0 {
			rlt = append(rlt, arr2[i])
			cnt[arr2[i]]--
		}
	}
	for i := 0; i < 1001; i++ {
		for cnt[i] != 0 {
			rlt = append(rlt, i)
			cnt[i]--
		}
	}
	return rlt
}
