package leetcode

/*Alice and Bob have candy bars of different sizes: A[i] is the size of the i-th bar of candy that Alice has, and B[j] is the size of the j-th bar of candy that Bob has.

Since they are friends, they would like to exchange one candy bar each so that after the exchange, they both have the same total amount of candy.  (The total amount of candy a person has is the sum of the sizes of candy bars they have.)

Return an integer array ans where ans[0] is the size of the candy bar that Alice must exchange, and ans[1] is the size of the candy bar that Bob must exchange.

If there are multiple answers, you may return any one of them.  It is guaranteed an answer exists.



来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/fair-candy-swap
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

import "sort"

func fairCandySwap(A []int, B []int) []int {
	res := make([]int, 2)
	sort.Ints(A)
	sort.Ints(B)
	sumA, sumB := 0, 0
	sizeA, sizeB := len(A), len(B)
	for i := 0; i < sizeA; i++ {
		sumA += A[i]
	}
	for i := 0; i < sizeB; i++ {
		sumB += B[i]
	}

	diff := (sumA - sumB) / 2
	for i := 0; i < sizeA; i++ {
		find := binarySearch(B, A[i]-diff)
		if find != -1 {
			res[0] = A[i]
			res[1] = B[find]
			break
		}
	}
	return res
}

func binarySearch(arr []int, target int) int {
	begin := 0
	end := len(arr) - 1
	for begin <= end {
		mid := (begin + end) / 2
		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			begin = mid + 1
		} else {
			end = mid - 1
		}
	}
	return -1
}
