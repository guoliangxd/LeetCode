package leetcode

/*Given an input string, reverse the string word by word.*/

import "strings"

func reverseWords(s string) string {
	subStr := make([]string, 0)
	size := len(s)
	end := -1
	for i := size - 1; i >= 0; i-- {
		if s[i] == ' ' {
			if end != -1 {
				subStr = append(subStr, s[i+1:end+1])
				end = -1
			}
		} else {
			if end == -1 {
				end = i
			}
		}
	}
	if end != -1 {
		subStr = append(subStr, s[0:end+1])
	}

	return strings.Join(subStr, " ")
}
