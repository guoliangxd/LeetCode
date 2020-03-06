package leetcode

/*In a array A of size 2N, there are N+1 unique elements, and exactly one of these elements is repeated N times.

Return the element repeated N times.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/n-repeated-element-in-size-2n-array
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

func repeatedNTimes(A []int) int {
	mapInt := make(map[int]int, 0)
	length := len(A)
	for i := 0; i < length; i++ {
		_, ok := mapInt[A[i]]
		if ok {
			return A[i]
		} else {
			mapInt[A[i]] = 1
		}
	}
	return -1
}
