package leetcode

/*Given an array A of integers, for each integer A[i] we may choose any x with -K <= x <= K, and add x to A[i].

After this process, we have some array B.

Return the smallest possible difference between the maximum value of B and the minimum value of B.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/smallest-range-i
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

func smallestRangeI(A []int, K int) int {
	max, min := A[0], A[0]
	for _, v := range A {
		if v > max {
			max = v
		}
		if v < min {
			min = v
		}
	}

	ans := (max - K) - (min + K)
	if ans < 0 {
		ans = 0
	}
	return ans
}
