package leetcode

/*In a given grid, each cell can have one of three values:
the value 0 representing an empty cell;
the value 1 representing a fresh orange;
the value 2 representing a rotten orange.
Every minute, any fresh orange that is adjacent (4-directionally) to a rotten orange becomes rotten.
Return the minimum number of minutes that must elapse until no cell has a fresh orange.
If this is impossible, return -1 instead.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/rotting-oranges
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

func orangesRotting(grid [][]int) int {
	fresh, rotten := 0, 0
	m, n := len(grid), len(grid[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				fresh++
			} else if grid[i][j] == 2 {
				rotten++
			}
		}
	}

	if fresh == 0 {
		return 0
	}

	if rotten == 0 {
		return -1
	}

	increase := 0
	k := 3
	for fresh > 0 {
		k++
		increase = 0
		for i := 0; i < m; i++ {
			for j := 0; j < n; j++ {
				if grid[i][j] == k-1 {
					grid[i][j] = 2
				}

				if grid[i][j] == 2 {
					if i-1 >= 0 && grid[i-1][j] == 1 {
						grid[i-1][j] = k
						fresh--
						increase++
					}
					if i+1 < m && grid[i+1][j] == 1 {
						grid[i+1][j] = k
						fresh--
						increase++
					}
					if j-1 >= 0 && grid[i][j-1] == 1 {
						grid[i][j-1] = k
						fresh--
						increase++
					}
					if j+1 < n && grid[i][j+1] == 1 {
						grid[i][j+1] = k
						fresh--
						increase++
					}
				}
			}
		}
		if increase == 0 {
			return -1
		}
	}
	return k - 3
}
