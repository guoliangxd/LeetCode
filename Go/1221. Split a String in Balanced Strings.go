package leetcode

/*Balanced strings are those who have equal quantity of 'L' and 'R' characters.
Given a balanced string s split it in the maximum amount of balanced strings.
Return the maximum amount of splitted balanced strings.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/split-a-string-in-balanced-strings
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

func balancedStringSplit(s string) int {
    rlt := 0
    cnt := 0
    length := len(s)
    for i := 0; i < length; i++ {
        if s[i] == 'R' {
            cnt++
        } else {
            cnt--
        }
        if cnt == 0 {
            rlt++
            cnt = 0
        }
    }
    return rlt 
}