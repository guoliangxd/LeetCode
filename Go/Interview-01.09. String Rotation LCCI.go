package leetcode

/*Given two strings, s1 and s2, write code to check if s2 is a rotation of s1 (e.g.,"waterbottle" is a rotation of"erbottlewat"). Can you use only one call to the method that checks if one word is a substring of another?

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/string-rotation-lcci
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

func isFlipedString(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	size := len(s1)

	if size == 0 {
		return true
	}

	for i := 0; i < size; i++ {
		if s1[i] == s2[0] {
			if s1[i:size]+s1[:i] == s2 {
				return true
			}
		}
	}
	return false
}
