package leetcode

/*Given an array arr, replace every element in that array with the greatest element among the elements to its right,
and replace the last element with -1.
After doing so, return the array.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/replace-elements-with-greatest-element-on-right-side
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

func replaceElements(arr []int) []int {
	length := len(arr)
	lastMax := arr[length-1]
	arr[length-1] = -1
	for i := length - 2; i >= 0; i-- {
		if arr[i] > lastMax {
			lastMax, arr[i] = arr[i], lastMax
		} else {
			arr[i] = lastMax
		}
	}
	return arr
}
