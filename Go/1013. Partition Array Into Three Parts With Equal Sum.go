package leetcode

/*Given an array A of integers, return true if and only if we can partition the array into three non-empty parts with equal sums.

Formally, we can partition the array if we can find indexes i+1 < j with (A[0] + A[1] + ... + A[i] == A[i+1] + A[i+2] + ... + A[j-1] == A[j] + A[j-1] + ... + A[A.length - 1])

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/partition-array-into-three-parts-with-equal-sum
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

func canThreePartsEqualSum(A []int) bool {
	sum := 0
	for i := 0; i < len(A); i++ {
		sum += A[i]
	}

	if sum%3 != 0 {
		return false
	}

	three := sum / 3
	l, r := 0, len(A)-1
	lSum, rSum := A[l], A[r]
	l++
	r--

	for l < r && (lSum != three || rSum != three) {
		if lSum != three {
			lSum += A[l]
			l++
		}
		if rSum != three {
			rSum += A[r]
			r--
		}
	}
	return l <= r && (lSum == three && rSum == three)
}
