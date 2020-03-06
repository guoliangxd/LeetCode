package leetcode

/*Given an array of distinct integers arr, 
find all pairs of elements with the minimum absolute difference of any two elements. 
Return a list of pairs in ascending order(with respect to pairs), each pair [a, b] follows
a, b are from arr
a < b
b - a equals to the minimum absolute difference of any two elements in arr

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/minimum-absolute-difference
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

import "sort"
func minimumAbsDifference(arr []int) [][]int {
    sort.Ints(arr)
    rlt := make([][]int, 0, 1)
    minDiff := arr[1] - arr[0]
    length := len(arr)
    for i := 0; i < length - 1; i++ {
        if arr[i + 1] - arr[i] < minDiff {
            minDiff = arr[i + 1] - arr[i]
        }
    }
    for i := 0; i < length - 1; i++ {
        if arr[i + 1] - arr[i] == minDiff {
            rlt = append(rlt, []int{arr[i], arr[i + 1]})
        }
    }
    return rlt
}