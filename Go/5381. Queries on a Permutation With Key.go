package leetcode

/*Given the array queries of positive integers between 1 and m, you have to process all queries[i] (from i=0 to i=queries.length-1) according to the following rules:

In the beginning, you have the permutation P=[1,2,3,...,m].
For the current i, find the position of queries[i] in the permutation P (indexing from 0) and then move this at the beginning of the permutation P. Notice that the position of queries[i] in P is the result for queries[i].
Return an array containing the result for the given queries.*/

func processQueries(queries []int, m int) []int {
	index := make([]int, m+1)
	for i := 0; i < m+1; i++ {
		index[i] = i - 1
	}

	size := len(queries)
	res := make([]int, size)
	for i := 0; i < size; i++ {
		pos := index[queries[i]]
		for j := 1; j < m+1; j++ {
			if index[j] < pos {
				index[j]++
			}
		}
		index[queries[i]] = 0
		res[i] = pos
	}
	return res
}