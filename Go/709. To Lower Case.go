package leetcode

/*Implement function ToLowerCase() that has a string parameter str, and returns the same string in lowercase.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/to-lower-case
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

func toLowerCase(str string) string {
	strArr := []rune(str)
	diff := 'a' - 'A'
	length := len(strArr)

	for i := 0; i < length; i++ {
		if strArr[i] <= 'Z' && strArr[i] >= 'A' {
			strArr[i] += diff
		}
	}

	return string(strArr)
}
