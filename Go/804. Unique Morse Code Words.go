package leetcode

/*International Morse Code defines a standard encoding 
where each letter is mapped to a series of dots and dashes, 
as follows: "a" maps to ".-", "b" maps to "-...", "c" maps to "-.-.", and so on.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/unique-morse-code-words
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

func uniqueMorseRepresentations(words []string) int {
    morse := [...]string{".-","-...","-.-.","-..",".","..-.","--.","....","..",".---","-.-",".-..","--","-.","---",".--.","--.-",".-.","...","-","..-","...-",".--","-..-","-.--","--.."}

    rlt := make(map[string]int)
    length := len(words)
    for i := 0; i < length; i++ {
        letters := []rune(words[i])
        ll := len(letters)
        temp := ""
        for j := 0; j < ll; j++ {
            temp += morse[letters[j] - 'a']
        }
        v, ok := rlt[temp]
        if !ok {
            rlt[temp] = 1
        } else {
            rlt[temp] = v + 1
        }
    }
    return len(rlt)
}