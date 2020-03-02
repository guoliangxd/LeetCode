package leetcode

/*Given a string, you need to reverse the order of characters in each word within a sentence while still preserving whitespace and initial word order.

Example 1:
Input: "Let's take LeetCode contest"
Output: "s'teL ekat edoCteeL tsetnoc"
Note: In the string, each word is separated by single space and there will not be any extra space in the string.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/reverse-words-in-a-string-iii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

import "strings"

func reverseWords(s string) string {
	subStr := strings.Split(s, " ")
	for i := 0; i < len(subStr); i++ {
		subStr[i] = reverseWord(subStr[i])
	}
	return strings.Join(subStr, " ")
}

func reverseWord(s string) string {
	temp := []rune(s)
	length := len(temp)
	for i := 0; i < length/2; i++ {
		temp[i], temp[length-i-1] = temp[length-i-1], temp[i]
	}
	return string(temp)
}
