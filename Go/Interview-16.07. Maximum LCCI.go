package leetcode

/*Write a method that finds the maximum of two numbers.
You should not use if-else or any other comparison operator.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/maximum-lcci
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

func maximum(a int, b int) int {
	c := a - b

	sa := sign(a)
	sb := sign(b)
	sc := sign(c)

	//a, b异号，k = sign(a)
	use_sign_of_a := sa ^ sb
	//a, b同号, k = sign(c)
	use_sign_0f_c := flip(sa ^ sb)
	k := use_sign_of_a*sa + use_sign_0f_c*sc
	q := flip(k)
	return a*k + b*q
}

//若是正数返回1，负数返回0
func sign(x int) int {
	return flip((x >> 31) & 0x1)
}

//取反
func flip(x int) int {
	return 0x1 ^ x
}
