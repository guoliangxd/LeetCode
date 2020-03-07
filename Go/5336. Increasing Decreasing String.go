package leetcode

/*Given a string s. You should re-order the string using the following algorithm:

Pick the smallest character from s and append it to the result.
Pick the smallest character from s which is greater than the last appended character to the result and append it.
Repeat step 2 until you cannot pick more characters.
Pick the largest character from s and append it to the result.
Pick the largest character from s which is smaller than the last appended character to the result and append it.
Repeat step 5 until you cannot pick more characters.
Repeat the steps from 1 to 6 until you pick all characters from s.
In each step, If the smallest or the largest character appears more than once you can choose any occurrence and append it to the result.

Return the result string after sorting s with this algorithm.*/

func sortString(s string) string {
	cnt := [26]int{}
	length := len(s)
	for i := 0; i < length; i++ {
		cnt[int(s[i]-'a')]++
	}
	rlt := make([]byte, length, length)
	index := 0
	findMin := true
	for index < length {
		if findMin {
			for i := 0; i < 26; i++ {
				if cnt[i] > 0 {
					rlt[index] = byte(i) + 'a'
					index++
					cnt[i]--
				}
			}
			findMin = false
		} else {
			for i := 25; i >= 0; i-- {
				if cnt[i] > 0 {
					rlt[index] = byte(i) + 'a'
					index++
					cnt[i]--
				}
			}
			findMin = true
		}
	}
	return string(rlt)
}
