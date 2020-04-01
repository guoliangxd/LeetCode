package leetcode

/*You are given an array with all the numbers from 1 to N appearing exactly once, except for two number that is missing. How can you find the missing number in O(N) time and 0(1) space?

You can return the missing numbers in any order.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/missing-two-lcci
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

import "math"

func missingTwo(nums []int) []int {
	size := len(nums)
	sum := (size + 2) * (size + 2 + 1) / 2
	for i := 0; i < size; i++ {
		sum -= nums[i]
	}

	squareSum := 0
	for i := 1; i <= size+2; i++ {
		squareSum += i * i
	}

	for i := 0; i < size; i++ {
		squareSum -= nums[i] * nums[i]
	}

	a := 2
	b := -2 * sum
	c := sum*sum - squareSum
	res := make([]int, 2)
	res[0] = (-b + int(math.Sqrt(float64(b*b-4*a*c)))) / (2 * a)
	res[1] = sum - res[0]
	return res
}
