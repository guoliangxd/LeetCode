package leetcode

/*Given a non-negative integer numRows, generate the first numRows of Pascal's triangle.
杨辉三角
来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/pascals-triangle
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

func generate(numRows int) [][]int {
    rlt := make([][]int, 0, 0)
    for i := 1; i <= numRows; i++ {
        temp := make([]int, i, i)
        for j := 0; j < i; j++ {
            if j == 0 || j == i - 1 {
                temp[j] = 1
            } else {
                temp[j] = rlt[i - 2][j - 1] + rlt[i - 2][j]
            }
        }
        rlt = append(rlt, temp)
    }
    return rlt
}