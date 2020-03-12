package leetcode

/*Given an integer array sorted in non-decreasing order, there is exactly one integer in the array that occurs more than 25% of the time.

Return that integer.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/element-appearing-more-than-25-in-sorted-array
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

func findSpecialInteger(arr []int) int {
	size := len(arr)
	quart := float64(size) / float64(4)
	ele := -1
	num := 0.0
	for i := 0; i < size; i++ {
		if ele == arr[i] {
			num++
		} else {
			ele = arr[i]
			num = 1.0
		}
		if num > quart {
			return ele
		}
	}
	return -1
}
