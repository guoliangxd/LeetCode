package leetcode

/*A sentence S is given, composed of words separated by spaces. Each word consists of lowercase and uppercase letters only.

We would like to convert the sentence to "Goat Latin" (a made-up language similar to Pig Latin.)

The rules of Goat Latin are as follows:

If a word begins with a vowel (a, e, i, o, or u), append "ma" to the end of the word.
For example, the word 'apple' becomes 'applema'.

If a word begins with a consonant (i.e. not a vowel), remove the first letter and append it to the end, then add "ma".
For example, the word "goat" becomes "oatgma".

Add one letter 'a' to the end of each word per its word index in the sentence, starting with 1.
For example, the first word gets "a" added to the end, the second word gets "aa" added to the end and so on.
Return the final sentence representing the conversion from S to Goat Latin.



来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/goat-latin
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

import "strings"

func toGoatLatin(S string) string {
	words := strings.Split(S, " ")
	res := make([]string, len(words))
	for i := 0; i < len(words); i++ {
		switch words[i][0] {
		case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
			res[i] = words[i] + "ma"
		default:
			res[i] = words[i][1:] + words[i][0:1] + "ma"
		}
		for j := i; j >= 0; j-- {
			res[i] += "a"
		}
	}
	return strings.Join(res, " ")
}
