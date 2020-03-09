package leetcode

/*We are to write the letters of a given string S, from left to right into lines. Each line has maximum width 100 units, and if writing a letter would cause the width of the line to exceed 100 units, it is written on the next line. We are given an array widths, an array where widths[0] is the width of 'a', widths[1] is the width of 'b', ..., and widths[25] is the width of 'z'.

Now answer two questions: how many lines have at least one character from S, and what is the width used by the last such line? Return your answer as an integer list of length 2.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/number-of-lines-to-write-string
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

func numberOfLines(widths []int, S string) []int {
	letters, lines := 0, 1
	for i := 0; i < len(S); i++ {
		temp := letters + widths[S[i]-'a']
		if temp > 100 {
			lines++
			letters = widths[S[i]-'a']
		} else if temp == 100 {
			lines++
			letters = 0
		} else {
			letters = temp
		}
	}
	return []int{lines, letters}
}
