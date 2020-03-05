package leetcode

/*You are given an array of strings words and a string chars.
A string is good if it can be formed by characters from chars (each character can only be used once).
Return the sum of lengths of all good strings in words.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/find-words-that-can-be-formed-by-characters
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

func countCharacters(words []string, chars string) int {
	cnt := [26]int{}
	for i := 0; i < len(chars); i++ {
		cnt[int(chars[i]-'a')]++
	}

	rlt := 0
	for i := 0; i < len(words); i++ {
		temp := [26]int{}
		for j := 0; j < len(words[i]); j++ {
			temp[int(words[i][j]-'a')]++
		}
		contain := true
		for k := 0; k < 26; k++ {
			if cnt[k] < temp[k] {
				contain = false
				break
			}
		}
		if contain {
			rlt += len(words[i])
		}
	}
	return rlt
}
