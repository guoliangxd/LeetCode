package leetcode

/*Given a sorted array of strings that is interspersed with empty strings,
write a method to find the location of a given string.*/

func findString(words []string, s string) int {
	length := len(s)
	for index, str := range words {
		if length != len(str) {
			continue
		}
		if s == str {
			return index
		}
	}
	return -1
}
