package leetcode

/*给定一个数组 A[0,1,…,n-1]，请构建一个数组 B[0,1,…,n-1]，其中 B 中的元素 B[i]=A[0]×A[1]×…×A[i-1]×A[i+1]×…×A[n-1]。不能使用除法。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/gou-jian-cheng-ji-shu-zu-lcof
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

func constructArr(a []int) []int {
	//对称遍历法
	left, right := 1, 1
	res := make([]int, len(a))
	for i := 0; i < len(a); i++ {
		res[i] = left
		left *= a[i]
	}
	for i := len(a) - 1; i >= 0; i-- {
		res[i] *= right
		right *= a[i]
	}
	return res
}
