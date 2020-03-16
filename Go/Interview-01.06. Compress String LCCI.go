package leetcode

/*Implement a method to perform basic string compression using the counts of repeated characters. For example, the string aabcccccaaa would become a2blc5a3. If the "compressed" string would not become smaller than the original string, your method should return the original string. You can assume the string has only uppercase and lowercase letters (a - z).

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/compress-string-lcci
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

import "strconv"

func compressString(S string) string {
	size := len(S)
	res := ""
	if size == 0 {
		return res
	}

	last := S[0]
	cnt := 1
	for i := 1; i < size; i++ {
		if S[i] != last {
			res += string(last) + strconv.Itoa(cnt)
			cnt = 1
			last = S[i]
		} else {
			cnt++
		}
	}
	res += string(last) + strconv.Itoa(cnt)
	if len(res) >= size {
		res = S
	}
	return res
}
