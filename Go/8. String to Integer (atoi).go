package leetcode

/*Implement atoi which converts a string to an integer.

The function first discards as many whitespace characters as necessary until the first non-whitespace character is found. Then, starting from this character, takes an optional initial plus or minus sign followed by as many numerical digits as possible, and interprets them as a numerical value.

The string can contain additional characters after those that form the integral number, which are ignored and have no effect on the behavior of this function.

If the first sequence of non-whitespace characters in str is not a valid integral number, or if no such sequence exists because either str is empty or it contains only whitespace characters, no conversion is performed.

If no valid conversion could be performed, a zero value is returned.

Note:

Only the space character ' ' is considered as whitespace character.
Assume we are dealing with an environment which could only store integers within the 32-bit signed integer range: [−231,  231 − 1]. If the numerical value is out of the range of representable values, INT_MAX (231 − 1) or INT_MIN (−231) is returned.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/string-to-integer-atoi
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

const (
	INT_MIN int64 = -2147483648
	INT_MAX int64 = 2147483647
)

func myAtoi(str string) int {
	var res int64 = 0
	str += " "
	size := len(str)
	num := ""
	for i := 0; i < size; i++ {
		if str[i] != ' ' {
			for j := i + 1; j < size; j++ {
				if str[j] >= '0' && str[j] <= '9' {
				} else {
					num = str[i:j]
					break
				}
			}
			break
		}
	}

	if len(num) != 0 {
		positive := true
		index := 0
		if num[0] == '-' {
			positive = false
			index++
		} else if num[0] == '+' {
			index++
		}
		if index < len(num) && num[index] >= '0' && num[index] <= '9' {
			for ; index < len(num); index++ {
				res *= 10
				res += int64(num[index] - '0')
				if positive {
					if res > INT_MAX {
						res = INT_MAX
						break
					}
				} else {
					if -res < INT_MIN {
						res = -INT_MIN
						break
					}
				}
			}
			if !positive {
				res = -res
			}
		}
	}

	return int(res)
}
