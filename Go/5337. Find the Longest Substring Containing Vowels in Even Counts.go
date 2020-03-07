package leetcode

/*Given the string s, return the size of the longest substring containing each vowel an even number of times. That is, 'a', 'e', 'i', 'o', and 'u' must appear an even number of times.*/

func findTheLongestSubstring(s string) int {
	max := len(s)
	for {
		for i := 0; i <= len(s)-max; i++ {
			if isOk(s[i : i+max]) {
				return max
			}
		}
		max--
	}
	return -1
}

func isOk(s string) bool {
	a, e, i, o, u := 0, 0, 0, 0, 0
	for _, v := range s {
		switch v {
		case 'a':
			a++
		case 'e':
			e++
		case 'i':
			i++
		case 'o':
			o++
		case 'u':
			u++
		default:
		}
	}
	return a%2 == 0 && e%2 == 0 && i%2 == 0 && o%2 == 0 && u%2 == 0
}
