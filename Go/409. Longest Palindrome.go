package leetcode

/*Given a string which consists of lowercase or uppercase letters, find the length of the longest palindromes that can be built with those letters.

This is case sensitive, for example "Aa" is not considered a palindrome here.

Note:
Assume the length of given string will not exceed 1,010.*/

func longestPalindrome(s string) int {
	up := [26]int{}
	low := [26]int{}

	for i := 0; i < len(s); i++ {
		if s[i] >= 'a' && s[i] <= 'z' {
			low[int(s[i]-'a')]++
		} else if s[i] >= 'A' && s[i] <= 'Z' {
			up[int(s[i]-'A')]++
		}
	}
	odd := 0
	for i := 0; i < 26; i++ {
		if up[i]%2 == 1 {
			odd++
		}
		if low[i]%2 == 1 {
			odd++
		}
	}

	if odd > 1 {
		return len(s) - odd + 1
	} else {
		return len(s)
	}
}
