package leetcode

/*Given a string S that only contains "I" (increase) or "D" (decrease), let N = S.length.

Return any permutation A of [0, 1, ..., N] such that for all i = 0, ..., N-1:

If S[i] == "I", then A[i] < A[i+1]
If S[i] == "D", then A[i] > A[i+1]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/di-string-match
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

func diStringMatch(S string) []int {
	length := len(S)
	min, max := 0, length
	rlt := make([]int, length+1, length+1)
	for i := 0; i < length; i++ {
		if S[i] == 'I' {
			rlt[i] = min
			min++
		} else {
			rlt[i] = max
			max--
		}
	}
	rlt[length] = max
	return rlt
}
