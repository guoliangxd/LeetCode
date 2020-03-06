package leetcode

/*Given a column title as appear in an Excel sheet, return its corresponding column number.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/excel-sheet-column-number
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

func titleToNumber(s string) int {
    rlt := 0
    for i := 0; i < len(s); i++ {
        rlt = rlt * 26 + int(s[i] - 'A') + 1
    }
    return rlt
}