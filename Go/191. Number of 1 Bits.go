package leetcode

/*Write a function that takes an unsigned integer 
and return the number of '1' bits it has (also known as the Hamming weight).

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/number-of-1-bits
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

func hammingWeight(n uint32) int {
    rlt := 0
    for n != 0 {
        if n & 1 == 1 {
            rlt++
        }
        n = n >> 1
    }
    return rlt
}