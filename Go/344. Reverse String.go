package leetcode

/*Write a function that reverses a string.
The input string is given as an array of characters char[].
Do not allocate extra space for another array,
you must do this by modifying the input array in-place with O(1) extra memory.
You may assume all the characters consist of printable ascii characters.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/reverse-string
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

func reverseString(s []byte) {
	length := len(s)

	for i := 0; i < length/2; i++ {
		s[i], s[length-i-1] = s[length-i-1], s[i]
	}
}
