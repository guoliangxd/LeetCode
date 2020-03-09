package leetcode

/*An array contains all the integers from 0 to n, except for one number which is missing.  Write code to find the missing integer. Can you do it in O(n) time?

Note: This problem is slightly different from the original one the book.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/missing-number-lcci
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

func missingNumber(nums []int) int {
	mapInt := make([]bool, len(nums)+1)

	for i := 0; i < len(nums); i++ {
		mapInt[nums[i]] = true
	}

	for i := 0; i < len(nums)+1; i++ {
		if !mapInt[i] {
			return i
		}
	}
	return -1
}


/*利用异或的性质
b^b=0
0^a=a
a^b^b=a
func missingNumber(nums []int) int {
    rlt := 0 
    for i := 0; i < len(nums); i++ {
        rlt ^= i
        rlt ^= nums[i]
    }
    return rlt^len(nums)
}*/