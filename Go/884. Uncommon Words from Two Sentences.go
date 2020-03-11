package leetcode

/*We are given two sentences A and B.  (A sentence is a string of space separated words.  Each word consists only of lowercase letters.)

A word is uncommon if it appears exactly once in one of the sentences, and does not appear in the other sentence.

Return a list of all uncommon words.

You may return the list in any order.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/uncommon-words-from-two-sentences
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

import "strings"

func uncommonFromSentences(A string, B string) []string {
	strMap := make(map[string]int)
	res := make([]string, 0)
	subA := strings.Split(A, " ")
	subB := strings.Split(B, " ")
	for _, str := range subA {
		v, ok := strMap[str]
		if !ok {
			strMap[str] = 1
		} else {
			strMap[str] = v + 1
		}
	}
	for _, str := range subB {
		v, ok := strMap[str]
		if !ok {
			strMap[str] = 1
		} else {
			strMap[str] = v + 1
		}
	}

	for k, v := range strMap {
		if v == 1 {
			res = append(res, k)
		}
	}
	return res
}
