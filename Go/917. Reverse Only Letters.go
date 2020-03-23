package leetcode

/*Given a string S, return the "reversed" string where all characters that are not a letter stay in the same place, and all letters reverse their positions.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/reverse-only-letters
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

func reverseOnlyLetters(S string) string {
	str := []byte(S)
	left, right := 0, len(S)-1
	for left < right {
		for left < right && !isAlpha(str[left]) {
			left++
		}
		for left < right && !isAlpha(str[right]) {
			right--
		}

		if left < right {
			str[left], str[right] = str[right], str[left]
			left++
			right--
		}
	}
	return string(str)
}

func isAlpha(char byte) bool {
	return char >= 'a' && char <= 'z' || char >= 'A' && char <= 'Z'
}
