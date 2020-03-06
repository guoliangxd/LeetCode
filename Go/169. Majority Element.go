package leetcode

/*Given an array of size n, find the majority element. The majority element is the element that appears more than ⌊ n/2 ⌋ times.

You may assume that the array is non-empty and the majority element always exist in the array.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/majority-element
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

func majorityElement(nums []int) int {
	mapCnt := make(map[int]int, 0)
	length := len(nums)

	for i := 0; i < length; i++ {
		v, ok := mapCnt[nums[i]]
		if ok {
			mapCnt[nums[i]] = v + 1
		} else {
			mapCnt[nums[i]] = 1
		}

		if mapCnt[nums[i]] == (length+1)/2 {
			return nums[i]
		}
	}
	return -1
}

/*通过排序取中间值
import "sort"
func majorityElement(nums []int) int {
    sort.Ints(nums)
    return nums[len(nums) / 2]
}
*/

/*Boyer-Moore 投票算法
func majorityElement(nums []int) int {
    major := nums[0]
    cnt := 1
    for i := 1; i < len(nums); i++ {
        if nums[i] != major {
            cnt--
        } else {
            cnt++
        }
        if cnt == 0 {
            i++
            major = nums[i]
            cnt = 1
        }
    }
    return major
}*/
