package leetcode

/*Given a string S and a character C, return an array of integers representing the shortest distance from the character C in the string.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/shortest-distance-to-a-character
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

func shortestToChar(S string, C byte) []int {
	length := len(S)
	rlt := make([]int, length, length)

	prv := -length
	for i := 0; i < length; i++ {
		if S[i] == C {
			prv = i
		}
		rlt[i] = i - prv
	}

	prv = 2 * length
	for i := length - 1; i >= 0; i-- {
		if S[i] == C {
			prv = i
		}
		temp := prv - i
		if temp < rlt[i] {
			rlt[i] = temp
		}
	}
	return rlt
}
