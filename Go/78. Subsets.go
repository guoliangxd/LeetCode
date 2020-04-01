package leetcode

/*Given a set of distinct integers, nums, return all possible subsets (the power set).

Note: The solution set must not contain duplicate subsets.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/subsets
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

func subsets(nums []int) [][]int {
	res := make([][]int, 0)
	size := len(nums)

	res = append(res, []int{})
	for i := 0; i < size; i++ {
		temp := make([][]int, len(res))
		for j := 0; j < len(res); j++ {
			data := make([]int, len(res[j])+1)
			copy(data, res[j])
			data[len(res[j])] = nums[i]
			temp[j] = data
		}
		res = append(res, temp...)
	}

	return res
}
