package leetcode

/*Given an integer n, return any array containing n unique integers such that they add up to 0.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/find-n-unique-integers-sum-up-to-zero
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

func sumZero(n int) []int {
    rlt := make([]int, n, n)
    for i := 0; i < n / 2; i++ {
        rlt[i] = i + 1
        rlt[n - i - 1] = 0 - rlt[i]
    }
    return rlt
}