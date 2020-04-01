package leetcode

/*Given two arrays of integers nums and index. Your task is to create target array under the following rules:

Initially target array is empty.
From left to right read nums[i] and index[i], insert at index index[i] the value nums[i] in target array.
Repeat the previous step until there are no elements to read in nums and index.
Return the target array.

It is guaranteed that the insertion operations will be valid.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/create-target-array-in-the-given-order
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

func createTargetArray(nums []int, index []int) []int {
	size := len(nums)
	res := make([]int, size)

	for i := 0; i < size; i++ {
		if index[i] == i {
			res[i] = nums[i]
		} else {
			move(res, index[i], i)
			res[index[i]] = nums[i]
		}
	}

	return res
}

func move(arr []int, index int, size int) {
	for i := size; i > index; i-- {
		arr[i] = arr[i-1]
	}
}
