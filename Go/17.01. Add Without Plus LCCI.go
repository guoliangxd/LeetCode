package leetcode

/*Write a function that adds two numbers. You should not use + or any arithmetic operators.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/add-without-plus-lcci
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

func add(a int, b int) int {
	for b != 0 {
		a, b = a^b, (a&b)<<1
	}
	return a
}
