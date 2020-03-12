package leetcode

/*Given words first and second, consider occurrences in some text of the form "first second third",
where second comes immediately after first, and third comes immediately after second.

For each such occurrence, add "third" to the answer, and return the answer.



来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/occurrences-after-bigram
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

import "strings"

func findOcurrences(text string, first string, second string) []string {
	word := strings.Split(text, " ")
	res := make([]string, 0)
	for i := 0; i < len(word)-2; i++ {
		if word[i] == first && word[i+1] == second {
			res = append(res, word[i+2])
		}
	}
	return res
}
