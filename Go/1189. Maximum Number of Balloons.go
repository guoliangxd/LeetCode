package leetcode

/*Given a string text, you want to use the characters of text to form as many instances of the word "balloon" as possible.

You can use each character in text at most once. Return the maximum number of instances that can be formed.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/maximum-number-of-balloons
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

func maxNumberOfBalloons(text string) int {
	b, a, l, o, n := 0, 0, 0, 0, 0
	for i := 0; i < len(text); i++ {
		switch text[i] {
		case 'b':
			b++
		case 'a':
			a++
		case 'l':
			l++
		case 'o':
			o++
		case 'n':
			n++
		default:
		}
	}
	rlt := b
	if a < rlt {
		rlt = a
	}
	if l/2 < rlt {
		rlt = l / 2
	}
	if o/2 < rlt {
		rlt = o / 2
	}
	if n < rlt {
		rlt = n
	}
	return rlt
}
