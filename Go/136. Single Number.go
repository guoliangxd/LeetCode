package leetcode

/*Given a non-empty array of integers, every element appears twice except for one. 
Find that single one.

Note:
Your algorithm should have a linear runtime complexity. 
Could you implement it without using extra memory?

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/single-number
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

/*
利用了异或的性质
a ^ 0 = a
a ^ a = 0
a ^ a ^ b = 0 ^ b = b
*/
func singleNumber(nums []int) int {
    rlt := 0
    length := len(nums)
    for i := 0; i < length; i++ {
        rlt ^= nums[i]
    }
    return rlt
}