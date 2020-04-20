package leetcode

/*Given a 2d grid map of '1's (land) and '0's (water), count the number of islands. An island is surrounded by water and is formed by connecting adjacent lands horizontally or vertically. You may assume all four edges of the grid are all surrounded by water.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/number-of-islands
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

func numIslands(grid [][]byte) int {
    if len(grid) == 0 {
        return 0
    }

    cnt := 2
    m, n := len(grid), len(grid[0])
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if grid[i][j] == '1' {
                search(i, j, cnt, grid)
                cnt++
            }
        }
    }
    return cnt - 2
}

func search(i int, j int, cnt int, grid [][]byte) {
    if grid[i][j] != '1' {
        return 
    }

    grid[i][j] = '0' + byte(cnt)
    if i - 1 >= 0 && grid[i - 1][j] == '1' {
        search(i - 1, j, cnt, grid)
    }
    if j - 1 >= 0 && grid[i][j - 1] == '1' {
        search(i, j - 1, cnt, grid)
    }
    m, n := len(grid), len(grid[0])
    if i + 1 < m && grid[i + 1][j] == '1' {
        search(i + 1, j, cnt, grid)
    }
    if j + 1 < n && grid[i][j + 1] == '1' {
        search(i, j + 1, cnt, grid)
    }
}