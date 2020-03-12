package leetcode

/*Let's define a function f(s) over a non-empty string s, which calculates the frequency of the smallest character in s. For example, if s = "dcce" then f(s) = 2 because the smallest character is "c" and its frequency is 2.

Now, given string arrays queries and words, return an integer array answer, where each answer[i] is the number of words such that f(queries[i]) < f(W), where W is a word in words.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/compare-strings-by-frequency-of-the-smallest-character
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

func numSmallerByFrequency(queries []string, words []string) []int {
	res := make([]int, len(queries))
	cntMap := [11]int{}
	sumMap := [11]int{}
	for i := 0; i < len(words); i++ {
		cntMap[f(words[i])]++
	}
	sumMap[0] = len(words)
	for i := 1; i < len(cntMap); i++ {
		sumMap[i] = sumMap[i-1] - cntMap[i]
	}
	for i := 0; i < len(res); i++ {
		res[i] = sumMap[f(queries[i])]
	}
	return res
}

func f(s string) int {
	cnt := [26]int{}
	for i := 0; i < len(s); i++ {
		cnt[s[i]-'a']++
	}
	for i := 0; i < len(cnt); i++ {
		if cnt[i] > 0 {
			return cnt[i]
		}
	}
	return -1
}
