package leetcode

/*X is a good number if after rotating each digit individually by 180 degrees, we get a valid number that is different from X.  Each digit must be rotated - we cannot choose to leave it alone.

A number is valid if each digit remains a digit after rotation. 0, 1, and 8 rotate to themselves; 2 and 5 rotate to each other; 6 and 9 rotate to each other, and the rest of the numbers do not rotate to any other number and become invalid.

Now given a positive number N, how many numbers X from 1 to N are good?

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/rotated-digits
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

func rotatedDigits(N int) int {
	res := 0
	for i := 1; i <= N; i++ {
		if isGood(i) {
			res++
		}
	}
	return res
}

func isGood(N int) bool {
	differ := false
	for N > 0 {
		switch N % 10 {
		case 0:
		case 1:
		case 8:
		case 2, 5, 6, 9:
			differ = true
		default:
			return false
		}
		N = N / 10
	}
	return differ
}
