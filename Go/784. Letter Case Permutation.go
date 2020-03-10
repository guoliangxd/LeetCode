package leetcode

/*Given a string S, we can transform every letter individually to be lowercase or uppercase to create another string.  Return a list of all possible strings we could create.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/letter-case-permutation
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

var res []string

func letterCasePermutation(S string) []string {
	res = make([]string, 0)
	for i := 0; i < len(S); i++ {
		Permutation(S[i:])
	}
	return res
}

func Permutation(S string) {
	if len(S) <= 0 {
		return
	}

	char := byte('a')
	if S[0] >= 'a' && S[0] <= 'z' {
		char = S[0]
	} else if S[0] >= 'A' && S[0] <= 'Z' {
		char = S[0] + byte(32)
	} else {
		char = S[0]
		if len(res) == 0 {
			res = append(res, string(char))
			return
		}
		for i := 0; i < len(res); i++ {
			res[i] = res[i] + string(char)
		}
		return
	}

	Char := char - byte(32)
	if len(res) == 0 {
		res = append(res, string(char), string(Char))
		return
	}
	temp := make([]string, len(res))
	for i := 0; i < len(res); i++ {
		temp[i] = res[i] + string(char)
		res[i] = res[i] + string(Char)
	}
	res = append(res, temp...)
}
