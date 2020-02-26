package leetcode

/*Given a positive integer num consisting only of digits 6 and 9.
Return the maximum number you can get by changing at most one digit (6 becomes 9, and 9 becomes 6).

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/maximum-69-number
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

import "strconv"

func maximum69Number(num int) int {
	str := []rune(strconv.Itoa(num))
	for i := 0; i < len(str); i++ {
		if str[i] == '6' {
			str[i] = '9'
			break
		}
	}

	rlt, _ := strconv.Atoi(string(str))
	return rlt
}
