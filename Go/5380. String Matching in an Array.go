package leetcode

/*Given an array of string words. Return all strings in words which is substring of another word in any order.

String words[i] is substring of words[j], if can be obtained removing some characters to left and/or right side of words[j].*/

import "strings"

func stringMatching(words []string) []string {
	res := make([]string, 0)
	size := len(words)
	for i := 0; i < size; i++ {
		subSize := len(words[i])
		for j := 0; j < size; j++ {
			if subSize >= len(words[j]) {
				continue
			}
			if strings.Contains(words[j], words[i]) {
				res = append(res, words[i])
				break
			}
		}
	}
	return res
}
