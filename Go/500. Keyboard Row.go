package leetcode

/*Given a List of words, return the words that can be typed using letters of alphabet on only one row's of American keyboard like the image below.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/keyboard-row
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

func findWords(words []string) []string {
    rlt := make([]string, 0)
    for i := 0; i < len(words); i++ {
        line := getLine(words[i][0])
        sameLine := true
        for j := 1; j < len(words[i]); j++ {
            if line != getLine(words[i][j]) {
                sameLine = false
                break
            }
        }
        if sameLine {
            rlt = append(rlt, words[i])
        }
    }
    return rlt
}

func getLine(char byte) int {
    rlt := 0
    switch char {
        case 'Q', 'W', 'E', 'R', 'T', 'Y', 'U', 'I', 'O', 'P', 'q', 'w', 'e', 'r', 't', 'y', 'u', 'i', 'o', 'p' :
            rlt = 1
        case 'A', 'S', 'D', 'F', 'G', 'H', 'J', 'K', 'L', 'a', 's', 'd', 'f', 'g', 'h', 'j','k', 'l' : 
            rlt = 2
        case 'Z', 'X', 'C', 'V', 'B', 'N', 'M', 'z', 'x', 'c', 'v', 'b', 'n', 'm':
            rlt = 3
    }
    return rlt
}