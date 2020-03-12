package leetcode

/*有一个同学在学习分式。他需要将一个连分数化成最简分数，你能帮助他吗？
连分数是形如上图的分式。在本题中，所有系数都是大于等于0的整数。
输入的cont代表连分数的系数（cont[0]代表上图的a0，以此类推）。返回一个长度为2的数组[n, m]，使得连分数的值等于n / m，且n, m最大公约数为1。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/deep-dark-fraction
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

func fraction(cont []int) []int {
	res := make([]int, 2)
	res[0] = cont[len(cont)-1]
	res[1] = 1
	for i := len(cont) - 2; i >= 0; i-- {
		res[0], res[1] = res[1], res[0]
		res[0] = res[1]*cont[i] + res[0]
	}
	return res
}
