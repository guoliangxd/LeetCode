package leetcode

/*In an alien language, surprisingly they also use english lowercase letters, but possibly in a different order. The order of the alphabet is some permutation of lowercase letters.

Given a sequence of words written in the alien language, and the order of the alphabet, return true if and only if the given words are sorted lexicographicaly in this alien language.*/

var weight [26]int

func isAlienSorted(words []string, order string) bool {
	weight = [26]int{}
	for i := 0; i < 26; i++ {
		weight[int(order[i])-'a'] = i
	}

	for i := 0; i < len(words)-1; i++ {
		if less(words[i+1], words[i]) {
			return false
		}
	}
	return true
}

func less(w1, w2 string) bool {
	size := len(w1)
	if size > len(w2) {
		size = len(w2)
	}
	for i := 0; i < size; i++ {
		if weight[int(w1[i]-'a')] < weight[int(w2[i]-'a')] {
			return true
		} else if weight[int(w1[i]-'a')] > weight[int(w2[i]-'a')] {
			return false
		}
	}

	if size < len(w1) || w1 == w2 {
		return false
	}
	return true
}
