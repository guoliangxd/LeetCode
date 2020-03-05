package leetcode

/*You are given a map in form of a two-dimensional integer grid where 1 represents land and 0 represents water.

Grid cells are connected horizontally/vertically (not diagonally). The grid is completely surrounded by water, and there is exactly one island (i.e., one or more connected land cells).

The island doesn't have "lakes" (water inside that isn't connected to the water around the island). One cell is a square with side length 1. The grid is rectangular, width and height don't exceed 100. Determine the perimeter of the island.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/island-perimeter
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

func islandPerimeter(grid [][]int) int {
    rlt := 0
    m, n := len(grid), len(grid[0])
    
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if grid[i][j] == 1 {
                if i - 1 < 0 || grid[i - 1][j] == 0 {
                    rlt++
                }
                if i + 1 >= m || grid[i + 1][j] == 0 {
                    rlt++
                }
                if j - 1 < 0 || grid[i][j - 1] == 0 {
                    rlt++
                }
                if j + 1 >= n || grid[i][j + 1] == 0 {
                    rlt++
                }
            }
        }
    }
    return rlt
}