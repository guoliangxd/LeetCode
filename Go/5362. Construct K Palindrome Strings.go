package leetcode

/*Given a string s and an integer k. You should construct k non-empty palindrome strings using all the characters in s.

Return True if you can use all the characters in s to construct k palindrome strings or False otherwise.*/

func canConstruct(s string, k int) bool {

	size := len(s)
	cnt := [26]int{}

	if size < k {
		return false
	} else if size == k {
		return true
	}

	for i := 0; i < size; i++ {
		cnt[int(s[i]-'a')]++
	}

	odd := 0
	for i := 0; i < 26; i++ {
		if cnt[i]%2 == 1 {
			odd++
			cnt[i]--
		}
	}

	if odd > k {
		return false
	}

	return true
}
