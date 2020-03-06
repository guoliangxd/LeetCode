package leetcode

/*Given an array of integers arr, write a function that returns true if and only if the number of occurrences of each value in the array is unique.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/unique-number-of-occurrences
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

func uniqueOccurrences(arr []int) bool {
	mapIntCnt := make(map[int]int)
	mapCnt := make(map[int]int)

	for i := 0; i < len(arr); i++ {
		val, ok := mapIntCnt[arr[i]]
		if !ok {
			mapIntCnt[arr[i]] = 1
		} else {
			mapIntCnt[arr[i]] = val + 1
		}
	}

	for _, v := range mapIntCnt {
		_, ok := mapCnt[v]
		if ok {
			return false
		} else {
			mapCnt[v] = 1
		}
	}
	return true
}
