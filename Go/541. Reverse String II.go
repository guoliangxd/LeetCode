package leetcode

/*Given a string and an integer k, you need to reverse the first k characters for every 2k characters counting from the start of the string. If there are less than k characters left, reverse all of them. If there are less than 2k but greater than or equal to k characters, then reverse the first k characters and left the other as original.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/reverse-string-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

func reverseStr(s string, k int) string {
	str := []byte(s)
	round := len(str) / (2 * k)
	left := len(str) % (2 * k)
	for i := 0; i < round; i++ {
		reverse(str, i*2*k, i*2*k+k-1)
	}
	if left >= k {
		reverse(str, round*2*k, round*2*k+k-1)
	} else {
		reverse(str, round*2*k, len(str)-1)
	}
	return string(str)
}

func reverse(s []byte, begin, end int) {
	for begin < end {
		s[begin], s[end] = s[end], s[begin]
		begin++
		end--
	}
}
