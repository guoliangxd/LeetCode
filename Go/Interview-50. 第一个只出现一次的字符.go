package leetcode

/*在字符串 s 中找出第一个只出现一次的字符。如果没有，返回一个单空格。*/

func firstUniqChar(s string) byte {
	charMap := [26]int{}
	for i := 0; i < len(s); i++ {
		char := int(s[i] - 'a')
		if charMap[char] == 0 {
			charMap[char] = i + 1
		} else if charMap[char] > 0 {
			charMap[char] = -1
		}
	}
	res := 60000
	for i := 0; i < 26; i++ {
		if charMap[i] > 0 {
			if charMap[i] > 0 && charMap[i] < res {
				res = charMap[i]
			}
		}
	}

	if res <= 50001 {
		return s[res-1]
	} else {
		return ' '
	}
}
