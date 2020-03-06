package leetcode

/*Given an array A of non-negative integers, return an array consisting of all the even elements of A, followed by all the odd elements of A.

You may return any answer array that satisfies this condition.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/sort-array-by-parity
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

func sortArrayByParity(A []int) []int {
	left, right := 0, len(A)-1
	for left < right {
		for left < right && A[left]%2 == 0 {
			left++
		}
		for left < right && A[right]%2 == 1 {
			right--
		}
		if left < right {
			A[left], A[right] = A[right], A[left]
			left++
			right--
		}
	}
	return A
}
