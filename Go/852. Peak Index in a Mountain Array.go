package leetcode

/*Let's call an array A a mountain if the following properties hold:

A.length >= 3
There exists some 0 < i < A.length - 1 such that A[0] < A[1] < ... A[i-1] < A[i] > A[i+1] > ... > A[A.length - 1]
Given an array that is definitely a mountain, return any i such that A[0] < A[1] < ... A[i-1] < A[i] > A[i+1] > ... > A[A.length - 1].

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/peak-index-in-a-mountain-array
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

//暴力解法，用二分查找能有更高效率
func peakIndexInMountainArray(A []int) int {
    index := 0
    length := len(A)

    for i := 0; i < length - 1; i++ {
        if A[i] > A[i + 1] {
            index = i
            break
        }
    }
    return index
}