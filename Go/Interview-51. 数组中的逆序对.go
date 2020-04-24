package leetcode

/*
在数组中的两个数字，如果前面一个数字大于后面的数字，则这两个数字组成一个逆序对。输入一个数组，求出这个数组中的逆序对的总数。
*/

func reversePairs(nums []int) int {
	return mergeSort(nums)
}

func mergeSort(arr []int) int {
	res := 0
	if len(arr) < 2 {
		return res
	}

	mid := len(arr) / 2
	res += mergeSort(arr[:mid])
	res += mergeSort(arr[mid:])
	temp := make([]int, len(arr))
	copy(temp, arr)
	i, j, index := 0, mid, 0
	for i < mid && j < len(arr) {
		if temp[i] <= temp[j] {
			arr[index] = temp[i]
			i++
		} else {
			arr[index] = temp[j]
			j++
			res += (mid - i)
		}
		index++
	}
	for i < mid {
		arr[index] = temp[i]
		i++
		index++
	}
	for j < len(arr) {
		arr[index] = temp[j]
		j++
		index++
	}
	return res
}
