package leetcode

/*Given a fixed length array arr of integers, duplicate each occurrence of zero, shifting the remaining elements to the right.

Note that elements beyond the length of the original array are not written.

Do the above modifications to the input array in place, do not return anything from your function.

*/

func duplicateZeros(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] == 0 {
			move(arr, i+1)
			i++
		}
	}
}

func move(arr []int, index int) {
	for i := len(arr) - 1; i >= index; i-- {
		arr[i] = arr[i-1]
	}
}
