package leetcode

/*Given a list of words, we may encode it by writing a reference string S and a list of indexes A.

For example, if the list of words is ["time", "me", "bell"], we can write it as S = "time#bell#" and indexes = [0, 2, 5].

Then for each index, we will recover the word by reading from the reference string from that index until we reach a "#" character.

What is the length of the shortest reference string S possible that encodes the given words?

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/short-encoding-of-words
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

func minimumLengthEncoding(words []string) int {
	size := len(words)

	res := make([]string, 0)

	for i := 0; i < size; i++ {
		find := false
		for j := 0; j < len(res); j++ {
			a, b := len(words[i]), len(res[j])
			if a >= b {
				if res[j] == words[i][a-b:] {
					res[j] = words[i]
					find = true
					break
				}
			} else {
				if words[i] == res[j][b-a:] {
					find = true
					break
				}
			}
		}
		if !find {
			res = append(res, words[i])
		}
	}

	length := 0
	for i := 0; i < len(res); i++ {
		length = length + len(res[i]) + 1
	}
	return length
}
