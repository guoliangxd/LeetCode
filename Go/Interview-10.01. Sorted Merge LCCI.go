package leetcode

/*You are given two sorted arrays, A and B, where A has a large enough buffer at the end to hold B. Write a method to merge B into A in sorted order.

Initially the number of elements in A and B are m and n respectively.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/sorted-merge-lcci
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

//import "sort"
func merge(A []int, m int, B []int, n int)  {
    /*for i := 0; i < n; i++ {
        A[i + m] = B[i] 
    }
    sort.Ints(A)*/
    for m > 0 && n > 0 {
        if A[m -1] >= B[n - 1] {
            A[m + n - 1] = A[m - 1]
            m--
        } else {
            A[m + n - 1] = B[n - 1]
            n-- 
        }
    }

    for n > 0 {
        A[n -1] = B[n - 1]
        n--
    }
}