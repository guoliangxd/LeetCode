package leetcode

/*Given an array nums of integers, return how many of them contain an even number of digits.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/find-numbers-with-even-number-of-digits
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

import "strconv"
func findNumbers(nums []int) int {
	//两种方法效率相当，但转成字符串更耗内存
    rlt := 0
    length := len(nums)
    for i := 0; i < length; i++ {
        tmpLen := len(strconv.Itoa(nums[i]))
        /*tmpLen := 0
        for nums[i] != 0 {
            nums[i] /= 10
            tmpLen++
        }*/
        if tmpLen % 2 == 0 {
            rlt++
        }
    }
    return rlt
}