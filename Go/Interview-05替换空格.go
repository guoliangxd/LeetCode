package leetcode

/*请实现一个函数，把字符串 s 中的每个空格替换成"%20"。*/

import "strings"
func replaceSpace(s string) string {
    subStr := strings.Split(s, " ")
    return strings.Join(subStr, "%20")
}