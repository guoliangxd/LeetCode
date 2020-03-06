package leetcode

/*Given a positive integer, output its complement number.
The complement strategy is to flip the bits of its binary representation.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/number-complement
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

func findComplement(num int) int {
	rlt := 0
	pow := 1
	for num != 0 {
		rlt = (num&0x1^0x1)*pow + rlt
		num = num >> 1
		pow = pow << 1
	}
	return rlt
}
