package leetcode

/*Find the minimum length word from a given dictionary words, which has all the letters from the string licensePlate. Such a word is said to complete the given string licensePlate

Here, for letters we ignore case. For example, "P" on the licensePlate still matches "p" on the word.

It is guaranteed an answer exists. If there are multiple answers, return the one that occurs first in the array.

The license plate might have the same letter occurring multiple times. For example, given a licensePlate of "PP", the word "pair" does not complete the licensePlate, but the word "supper" does.*/

func shortestCompletingWord(licensePlate string, words []string) string {
    res := "11111111111111111111"
    cnt := [27]int{}
    for i := 0; i < len(licensePlate); i++ {
        cnt[toLower(licensePlate[i])]++
    }
    for i := 0; i < len(words); i++ {
        temp := [26]int{}
        match := true
        for j := 0; j < len(words[i]); j++ {
            temp[int(words[i][j] - 'a')]++
        }
        for k := 0; k < 26; k++ {
            if temp[k] < cnt[k] {
                match = false
                break 
            }
        }
        if match && len(res) > len(words[i]){
            res = words[i]
        }
    }
    return res
}

func toLower(char byte) int {
    if char >= 'a' && char <= 'z' {
        return int(char - 'a')
    } else if char >= 'A' && char <= 'Z' {
        return int(char - 'A')
    } else {
        return 26
    }
}

