package leetcode

/*Given a positive integer, check whether it has alternating bits: namely, if two adjacent bits will always have different values.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/binary-number-with-alternating-bits
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

func hasAlternatingBits(n int) bool {
	bit := n & 0x1
	n = n >> 1
	for n > 0 {
		if n&0x1 != (bit ^ 0x1) {
			return false
		}
		bit ^= 0x1
		n = n >> 1
	}
	return true
}
