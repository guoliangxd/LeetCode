package leetcode

/*Given n pairs of parentheses, write a function to generate all combinations of well-formed parentheses.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/generate-parentheses
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/
//TODO:回溯算法
func generateParenthesis(n int) []string {
	res := make([]string, 0)
	backtrack(&res, []byte{}, 0, 0, n)
	return res
}

func backtrack(res *[]string, cur []byte, open int, close int, max int) {
	if len(cur) == max*2 {
		*res = append(*res, string(cur))
		return
	}

	if open < max {
		cur = append(cur, '(')
		backtrack(res, cur, open+1, close, max)
		cur = cur[:len(cur)-1]
	}
	if close < open {
		cur = append(cur, ')')
		backtrack(res, cur, open, close+1, max)
		cur = cur[:len(cur)-1]
	}
}
