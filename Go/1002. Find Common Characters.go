package leetcode

/*Given an array A of strings made only from lowercase letters, return a list of all characters that show up in all strings within the list (including duplicates).  For example, if a character occurs 3 times in all strings but not 4 times, you need to include that character three times in the final answer.

You may return the answer in any order.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/find-common-characters
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

func commonChars(A []string) []string {
	//词频统计结果
	cnt := [26]int{}
	//暂时以第一单词的词频作为结果
	for i := 0; i < len(A[0]); i++ {
		cnt[int(A[0][i]-'a')]++
	}
	//统计其他单词词频并与现有结果取交集
	for i := 1; i < len(A); i++ {
		temp := [26]int{}
		for j := 0; j < len(A[i]); j++ {
			temp[int(A[i][j]-'a')]++
		}
		for k := 0; k < 26; k++ {
			if temp[k] < cnt[k] {
				cnt[k] = temp[k]
			}
		}
	}
	//根据词频返回构造返回结果
	rlt := make([]string, 0)
	for i := 0; i < 26; i++ {
		for cnt[i] != 0 {
			rlt = append(rlt, string('a'+byte(i)))
			cnt[i]--
		}
	}
	return rlt

}
