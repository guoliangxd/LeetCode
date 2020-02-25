package leetcode

/*Given a m * n matrix grid which is sorted in non-increasing order both row-wise and column-wise. 
Return the number of negative numbers in grid.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/count-negative-numbers-in-a-sorted-matrix
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

func countNegatives(grid [][]int) int {
    rlt := 0
    row := len(grid)
    col := len(grid[0])
    for i := 0; i < row; i++ {
        for j := 0; j < col; j++ {
            if grid[i][j] < 0 {
                rlt += col - j
                break
            }
        }
    }
    return rlt
}