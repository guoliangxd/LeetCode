package leetcode

/*输入数字 n，按顺序打印出从 1 到最大的 n 位十进制数。比如输入 3，则打印出 1、2、3 一直到最大的 3 位数 999。*/

import "strconv"
func printNumbers(n int) []int {
    strlen := ""
    for i := 0; i < n; i++ {
        strlen += "9"
    }
    length, _ := strconv.Atoi(strlen)

    rlt := make([]int, length, length)
    for i := 0; i < length; i++ {
        rlt[i] = i + 1
    }
    return rlt
}