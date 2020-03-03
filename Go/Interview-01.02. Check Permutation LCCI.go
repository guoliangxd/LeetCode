package leetcode

/*Given two strings,write a method to decide if one is a permutation of the other.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/check-permutation-lcci
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

func CheckPermutation(s1 string, s2 string) bool {
	cnt1 := make([]int, 26, 26)
	cnt2 := make([]int, 26, 26)
	len1 := len(s1)
	len2 := len(s2)
	if len1 != len2 {
		return false
	}
	for i := 0; i < len1; i++ {
		cnt1[int(s1[i]-'a')]++
		cnt2[int(s2[i]-'a')]++
	}
	for i := 0; i < 26; i++ {
		if cnt1[i] != cnt2[i] {
			return false
		}
	}
	return true
}
