package leetcode

/*Given an array of integers A sorted in non-decreasing order,
return an array of the squares of each number, also in sorted non-decreasing order.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/squares-of-a-sorted-array
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

//import "sort"
func sortedSquares(A []int) []int {
	length := len(A)
	rlt := make([]int, length, length)
	i, j, index := 0, length-1, length-1
	for i <= j {
		if A[i] >= 0 {
			rlt[index] = A[j] * A[j]
			j--
		} else if A[j] <= 0 {
			rlt[index] = A[i] * A[i]
			i++
		} else {
			if -A[i] >= A[j] {
				rlt[index] = A[i] * A[i]
				i++
			} else {
				rlt[index] = A[j] * A[j]
				j--
			}
		}
		index--
	}
	return rlt
	/*
	   for i := 0; i < length; i++ {
	       A[i] = A[i] * A[i]
	   }
	   sort.Ints(A)
	   return A*/
}
