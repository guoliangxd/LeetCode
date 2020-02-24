package leetcode

/*Given an integer number n, return the difference between the product of its digits and the sum of its digits.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/subtract-the-product-and-sum-of-digits-of-an-integer
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

import "strconv"
func subtractProductAndSum(n int) int {
    str := strconv.Itoa(n)
    length := len(str)
    products := 1
    sum := 0
    for i := 0; i < length; i++ {
        temp := int(str[i] - '0')
        products *= temp
        sum += temp
    }
    return products - sum
}

