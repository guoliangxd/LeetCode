package leetcode

/*A magic index in an array A[0...n-1] is defined to be an index such that A[i] = i. Given a sorted array of distinct integers, write a method to find a magic index, if one exists, in array A. If not, return -1. If there are more than one magic index, return the smallest one.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/magic-index-lcci
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

func findMagicIndex(nums []int) int {
	for i := 0; i < len(nums); i++ {
		if i == nums[i] {
			return i
		}
	}
	return -1
}
