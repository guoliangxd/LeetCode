package leetcode

/*Given a non-empty 2D array grid of 0's and 1's, an island is a group of 1's (representing land) connected 4-directionally (horizontal or vertical.) You may assume all four edges of the grid are surrounded by water.

Find the maximum area of an island in the given 2D array. (If there is no island, the maximum area is 0.)

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/max-area-of-island
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

func maxAreaOfIsland(grid [][]int) int {
    res := 0
    for i := 0; i < len(grid); i++ {
        for j := 0; j < len(grid[0]); j++ {
            temp := calcArea(i, j, grid)
            if temp > res {
                res = temp
            }
        }
    }
    return res
}

func calcArea(i, j int, grid [][]int) int {
    m := len(grid)
    n := len(grid[0])
    area := 0
    if grid[i][j] == 1 {
        //标记已查找过的点，避免重复计数
        grid[i][j] = 2
        area = 1
        if i - 1 >= 0 {
            area += calcArea(i - 1, j, grid)
        }
        if i + 1 < m {
            area += calcArea(i + 1, j, grid)
        }
        if j - 1 >= 0 {
            area += calcArea(i, j - 1, grid)
        }
        if j + 1 < n {
            area += calcArea(i, j + 1, grid)
        }
    }
    return area 
}

