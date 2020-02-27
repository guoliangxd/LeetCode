package leetcode

/*Given a string s formed by digits ('0' - '9') and '#' .
We want to map s to English lowercase characters as follows:

Characters ('a' to 'i') are represented by ('1' to '9') respectively.
Characters ('j' to 'z') are represented by ('10#' to '26#') respectively.
Return the string formed after mapping.

It's guaranteed that a unique mapping will always exist.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/decrypt-string-from-alphabet-to-integer-mapping
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

import (
	"strconv"
)

func freqAlphabets(s string) string {
	length := len(s)
	rlt := []byte{}
	for i := 0; i < length; i++ {
		if i+2 < length {
			if s[i+2] == '#' {
				n, _ := strconv.Atoi(s[i : i+2])
				rlt = append(rlt, 'j'+byte(n-10))
				i += 2
			} else {
				rlt = append(rlt, 'a'+s[i]-'1')
			}
		} else {
			rlt = append(rlt, 'a'+s[i]-'1')
		}
	}
	return string(rlt)
}
