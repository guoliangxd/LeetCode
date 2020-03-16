package leetcode

/*Given two strings s and t , write a function to determine if t is an anagram of s.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/valid-anagram
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

func isAnagram(s string, t string) bool {
	sMap := [26]int{}
	tMap := [26]int{}

	size := len(s)
	if len(t) != size {
		return false
	}

	for i := 0; i < size; i++ {
		sMap[int(s[i]-'a')]++
		tMap[int(t[i]-'a')]++
	}
	for i := 0; i < 26; i++ {
		if sMap[i] != tMap[i] {
			return false
		}
	}
	return true
}
