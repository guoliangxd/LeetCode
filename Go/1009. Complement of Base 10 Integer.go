package leetcode

/*Every non-negative integer N has a binary representation.  For example, 5 can be represented as "101" in binary, 11 as "1011" in binary, and so on.  Note that except for N = 0, there are no leading zeroes in any binary representation.

The complement of a binary representation is the number in binary you get when changing every 1 to a 0 and 0 to a 1.  For example, the complement of "101" in binary is "010" in binary.

For a given number N in base-10, return the complement of it's binary representation as a base-10 integer.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/complement-of-base-10-integer
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

func bitwiseComplement(N int) int {
	res := 0
	pow := 1
	for {
		res = (N&0x1^0x1)*pow + res
		pow = pow << 1
		N = N >> 1
		if N == 0 {
			break
		}
	}
	return res
}
