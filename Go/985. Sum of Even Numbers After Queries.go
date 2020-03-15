package leetcode

/*We have an array A of integers, and an array queries of queries.

For the i-th query val = queries[i][0], index = queries[i][1], we add val to A[index].  Then, the answer to the i-th query is the sum of the even values of A.

(Here, the given index = queries[i][1] is a 0-based index, and each query permanently modifies the array A.)

Return the answer to all queries.  Your answer array should have answer[i] as the answer to the i-th query.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/sum-of-even-numbers-after-queries
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

func sumEvenAfterQueries(A []int, queries [][]int) []int {
	evenSum := 0
	res := make([]int, len(queries))
	for i := 0; i < len(A); i++ {
		if A[i]%2 == 0 {
			evenSum += A[i]
		}
	}

	for i := 0; i < len(queries); i++ {
		index := queries[i][1]
		if A[index]%2 == 0 {
			evenSum -= A[index]
		}
		A[index] += queries[i][0]
		if A[index]%2 == 0 {
			evenSum += A[index]
		}
		res[i] = evenSum
	}
	return res
}
