package leetcode

/*Given an array of integers arr, a lucky integer is an integer which has a frequency in the array equal to its value.

Return a lucky integer in the array. If there are multiple lucky integers return the largest of them. If there is no lucky integer return -1.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/find-lucky-integer-in-an-array
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

func findLucky(arr []int) int {
	cnt := [501]int{}
	size := len(arr)
	luckNum := -1
	for i := 0; i < size; i++ {
		cnt[arr[i]]++
	}
	for i := 1; i < 501; i++ {
		if cnt[i] == i {
			luckNum = i
		}
	}
	return luckNum
}
