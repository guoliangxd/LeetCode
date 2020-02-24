package leetcode

/*Given a non-negative integer num, return the number of steps to reduce it to zero.
If the current number is even, you have to divide it by 2, otherwise, you have to subtract 1 from it.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/number-of-steps-to-reduce-a-number-to-zero
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

func numberOfSteps(num int) int {
	if num == 0 {
		return 0
	} else if num%2 == 0 {
		return 1 + numberOfSteps(num/2)
	} else {
		return 1 + numberOfSteps(num-1)
	}
}
