package leetcode

/*Given a non-negative integer num, repeatedly add all its digits until the result has only one digit.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/add-digits
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

import "strconv"

func addDigits(num int) int {
	for num > 9 {
		str := strconv.Itoa(num)
		temp := 0
		for _, v := range str {
			temp += int(v - '0')
		}
		num = temp
	}
	return num
}
