package leetcode

/*A valid parentheses string is either empty (""), "(" + A + ")", or A + B, 
where A and B are valid parentheses strings, and + represents string concatenation.  
For example, "", "()", "(())()", and "(()(()))" are all valid parentheses strings.
A valid parentheses string S is primitive if it is nonempty, 
and there does not exist a way to split it into S = A+B, with A and B nonempty valid parentheses strings.

Given a valid parentheses string S, consider its primitive decomposition: S = P_1 + P_2 + ... + P_k, 
where P_i are primitive valid parentheses strings.

Return S after removing the outermost parentheses of every primitive string in the primitive decomposition of S.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/remove-outermost-parentheses
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

import "strings"
func removeOuterParentheses(S string) string {
   
    rlt := []rune(S)
    length := len(rlt)
    if length == 0  {
        return S
    }

    cnt := 0
    for i := 0; i < length; i++ {
        switch rlt[i] {
            case '(':
                if cnt == 0 {
                    rlt[i] = ' '
                }
                cnt++
            case ')':
                if cnt == 0 {
                    rlt[i] = ' '
                }
                cnt--
            default:
        }
        
        if cnt == 0 {
            rlt[i] = ' '
        }
    } 
    return strings.Join(strings.Split(string(rlt), " "), "")
}