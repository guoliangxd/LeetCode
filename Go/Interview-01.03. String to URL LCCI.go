package leetcode

/*Write a method to replace all spaces in a string with '%20'. You may assume that the string has sufficient space at the end to hold the additional characters,and that you are given the "true" length of the string. (Note: If implementing in Java,please use a character array so that you can perform this operation in place.)*/

func replaceSpaces(S string, length int) string {
	str := make([]byte, 0, length)
	for i := 0; i < length; i++ {
		if S[i] == ' ' {
			str = append(str, '%', '2', '0')
		} else {
			str = append(str, S[i])
		}
	}
	return string(str)
}
