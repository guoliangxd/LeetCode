package leetcode

/*Given a string s, find the longest palindromic substring in s. You may assume that the maximum length of s is 1000.*/

func longestPalindrome(s string) string {
	res := 0
	index := 0
	for i := 0; i < len(s); i++ {
		length := maxLength(s, i, i)
		if length > res {
			res = length
			index = i - length/2
		}
	}

	for i := 0; i < len(s)-1; i++ {
		length := maxLength(s, i, i+1)
		if length > res {
			res = length
			index = i - length/2 + 1
		}
	}
	return s[index : index+res]
}

func maxLength(s string, left, right int) int {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	return right - left - 1
}
