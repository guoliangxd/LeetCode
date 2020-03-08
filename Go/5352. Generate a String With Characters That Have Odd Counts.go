package leetcode

/*Given an integer n, return a string with n characters such that each character in such string occurs an odd number of times.

The returned string must contain only lowercase English letters. If there are multiples valid strings, return any of them.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/generate-a-string-with-characters-that-have-odd-counts
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

func generateTheString(n int) string {
	rlt := make([]byte, n, n)
	for i := 0; i < n-1; i++ {
		rlt[i] = 'a'
	}
	if n%2 == 0 {
		rlt[n-1] = 'b'
	} else {
		rlt[n-1] = 'a'
	}
	return string(rlt)
}
