package leetcode

/*Given a positive integer N, find and return the longest distance between two consecutive 1's in the binary representation of N.

If there aren't two consecutive 1's, return 0.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/binary-gap
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

func binaryGap(N int) int {
	max := 0
	cnt := 0
	findFirst := false
	for N > 0 {
		if N&0x1 == 1 {
			if !findFirst {
				findFirst = true
			} else {
				if cnt > max {
					max = cnt
				}
				cnt = 0
			}
		}
		N = N >> 1
		if findFirst {
			cnt++
		}
	}
	return max
}
